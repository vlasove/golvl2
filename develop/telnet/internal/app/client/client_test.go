package client

import (
	"bytes"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:")
	assert.NoError(t, err)
	defer func() { assert.NoError(t, l.Close()) }()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		in := &bytes.Buffer{}
		out := &bytes.Buffer{}

		timeout, err := time.ParseDuration("10s")
		assert.NoError(t, err)

		client := NewClient(l.Addr().String(), timeout, ioutil.NopCloser(in), out)
		assert.NoError(t, client.BuildConnection())
		defer func() { assert.NoError(t, client.Close()) }()

		in.WriteString("hello\n")
		err = client.Send()
		assert.NoError(t, err)

		err = client.Get()
		assert.NoError(t, err)
		assert.Equal(t, "world\n", out.String())
	}()

	go func() {
		defer wg.Done()

		conn, err := l.Accept()
		assert.NoError(t, err)
		assert.NotNil(t, conn)
		defer func() { assert.NoError(t, conn.Close()) }()

		request := make([]byte, 1024)
		n, err := conn.Read(request)
		assert.NoError(t, err)
		assert.Equal(t, "hello\n", string(request)[:n])

		n, err = conn.Write([]byte("world\n"))
		assert.NoError(t, err)
		assert.NotEqual(t, 0, n)
	}()

	wg.Wait()
}
