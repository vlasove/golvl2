package client

import (
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	normalExitMsg    = "go-telnet: program trying to exit...."
	endOfFileExitMsg = "go-telnet: recieved EOF"
)

// Runner ...
type Runner struct {
	host    string
	port    string
	timeout time.Duration
}

// New ...
func New(host, port string, timeout time.Duration) *Runner {
	return &Runner{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

// Start ...
func (r *Runner) Start() error {
	client := NewClient(r.host+":"+r.port, r.timeout, os.Stdin, os.Stdout)
	if err := client.BuildConnection(); err != nil {
		return err
	}
	defer client.Close()

	signalCh := make(chan os.Signal, 1)
	errorCh := make(chan error, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go get(client, errorCh)
	go send(client, errorCh)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-signalCh:
				log.Println(normalExitMsg)
				return
			case err := <-errorCh:
				if err != nil {
					if err == io.EOF {
						log.Println(endOfFileExitMsg)
					}
					return
				}
			default:
				continue
			}
		}
	}()

	wg.Wait()
	return nil
}

func send(c *Client, errorCh chan error) {
	for {
		if err := c.Send(); err != nil {
			errorCh <- err
			return
		}
	}
}

func get(c *Client, errorCh chan error) {
	for {
		if err := c.Get(); err != nil {
			errorCh <- err
			return
		}
	}
}
