package netcater

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	host  = "127.0.0.1"
	port  = "9991"
	data  = "test data"
	creds = host + ":" + port
)

func TestNecat_TCPTransfer(t *testing.T) {
	w, oldStdin := mockStreams(t)

	go func() {
		// дать паузу, чтобы сервер успел завестись
		time.Sleep(time.Second)
		nc := New(host, port, "tcp")
		conn, err := net.Dial("tcp", creds)
		assert.Nil(t, err)
		_, err = w.Write([]byte(data))
		assert.Nil(t, err)
		nc.transferTCP(conn)
	}()

	listener, err := net.Listen("tcp", ":"+port)
	assert.NoError(t, err)

	conn, err := listener.Accept()
	assert.NoError(t, err)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	assert.NoError(t, err)

	assert.Equal(t, data, string(buf[0:n]))
	// вернуть поток назад
	os.Stdin = oldStdin
}
func TestNetcat_UDPTransfer(t *testing.T) {
	w, oldStdin := mockStreams(t)

	go func() {
		time.Sleep(time.Second)
		conn, err := net.Dial("udp", creds)
		assert.Nil(t, err)
		_, err = w.Write([]byte(data))
		assert.Nil(t, err)
		nc := New(host, port, "udp")
		nc.transferUDP(conn)
	}()

	con, err := net.ListenPacket("udp", ":"+port)
	assert.Nil(t, err)

	buf := make([]byte, 1024)
	n, _, err := con.ReadFrom(buf)
	assert.Nil(t, err)

	assert.Equal(t, data, string(buf[0:n]))

	os.Stdin = oldStdin
}

func mockStreams(t *testing.T) (w *os.File, oldStdin *os.File) {
	t.Helper()
	oldStdin = os.Stdin
	r, w, err := os.Pipe()
	assert.Nil(t, err)
	os.Stdin = r
	return
}
