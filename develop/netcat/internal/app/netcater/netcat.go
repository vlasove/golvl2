package netcater

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
)

const (
	// Buffer ...
	Buffer = 2<<16 - 1
	// Disconnect ...
	Disconnect = "~!"
)

var (
	errUnsuportedProtocol       = errors.New("netcat: provided protocol usuported")
	errCanNotBuildTCPConnection = errors.New("netcat: can not build TCP connection")
	errCanNotResolveUDP         = errors.New("netcat: can not resolve UDP address")
	errCanNotBuildUDPConnection = errors.New("netcat: can not build UDP connection")
)

// NetCat ...
type NetCat struct {
	host     string
	port     string
	protocol string
}

// New ...
func New(host, port, protocol string) *NetCat {
	return &NetCat{
		host:     host,
		port:     port,
		protocol: protocol,
	}
}

// Start ...
func (nc *NetCat) Start() error {
	switch nc.protocol {
	case "tcp":
		if err := nc.TCPConnector(); err != nil {
			return err
		}
	case "udp":
		if err := nc.UDPConnector(); err != nil {
			return err
		}
	default:
		return errUnsuportedProtocol
	}
	return nil
}

// TCPConnector ...
func (nc *NetCat) TCPConnector() error {
	creds := nc.host + ":" + nc.port
	conn, err := net.Dial("tcp", creds)
	if err != nil {
		return errCanNotBuildTCPConnection
	}
	log.Println("netcat: connected to", creds)
	nc.transferTCP(conn)
	return nil
}

func (nc *NetCat) transferTCP(conn net.Conn) {
	c := make(chan tcpProgress)

	copy := func(r io.ReadCloser, w io.WriteCloser) {
		defer func() {
			r.Close()
			w.Close()
		}()
		n, err := io.Copy(w, r)
		if err != nil {
			log.Printf("[%s]: ERROR: %s\n", conn.RemoteAddr(), err)
		}
		c <- tcpProgress{bytes: uint64(n)}
	}

	go copy(conn, os.Stdout)
	go copy(os.Stdin, conn)

	p := <-c
	log.Printf(
		"[%s]: Connection has been closed by remote peer, %d bytes has been received\n",
		conn.RemoteAddr(),
		p.bytes,
	)
	p = <-c
	log.Printf(
		"[%s]: Local peer has been stopped, %d bytes has been sent\n",
		conn.RemoteAddr(),
		p.bytes,
	)
}

// UDPConnector ...
func (nc *NetCat) UDPConnector() error {
	creds := nc.host + ":" + nc.port
	addr, err := net.ResolveUDPAddr("udp", creds)
	if err != nil {
		return errCanNotResolveUDP
	}
	conn, err := net.DialUDP(nc.protocol, nil, addr)
	if err != nil {
		return errCanNotBuildUDPConnection
	}
	log.Println("netcat: connected to", creds)
	nc.transferUDP(conn)
	return nil
}

func (nc *NetCat) transferUDP(conn net.Conn) {
	c := make(chan udpProgress)

	copy := func(r io.ReadCloser, w io.WriteCloser, ra net.Addr) {
		defer func() {
			r.Close()
			w.Close()
		}()

		buf := make([]byte, Buffer)
		bytes := uint64(0)
		var n int
		var err error
		var addr net.Addr

		for {
			if connUDP, ok := r.(*net.UDPConn); ok {
				n, addr, err = connUDP.ReadFrom(buf)
				if connUDP.RemoteAddr() == nil && ra == nil {
					ra = addr
					c <- udpProgress{remoteAddr: ra}
				}
			} else {
				n, err = r.Read(buf)
			}
			if err != nil {
				if err != io.EOF {
					log.Printf("[%s]: ERROR: %s\n", ra, err)
				}
				break
			}
			if string(buf[0:n-1]) == Disconnect {
				break
			}

			if con, ok := w.(*net.UDPConn); ok && con.RemoteAddr() == nil {
				n, err = con.WriteTo(buf[0:n], ra)
			} else {
				n, err = w.Write(buf[0:n])
			}
			if err != nil {
				log.Printf("[%s]: ERROR: %s\n", ra, err)
				break
			}
			bytes += uint64(n)
		}
		c <- udpProgress{bytes: bytes}
	}

	ra := conn.RemoteAddr()
	go copy(conn, os.Stdout, ra)
	if ra == nil {
		p := <-c
		ra = p.remoteAddr
		log.Printf("[%s]: Datagram has been received\n", ra)
	}
	go copy(os.Stdin, conn, ra)

	p := <-c
	log.Printf(
		"[%s]: Connection has been closed, %d bytes has been received\n",
		ra,
		p.bytes,
	)
	p = <-c
	log.Printf(
		"[%s]: Local peer has been stopped, %d bytes has been sent\n",
		ra,
		p.bytes,
	)
}

type tcpProgress struct {
	bytes uint64
}

type udpProgress struct {
	remoteAddr net.Addr
	bytes      uint64
}
