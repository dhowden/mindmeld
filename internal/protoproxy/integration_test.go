package protoproxy_test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld/internal/protoproxy"

	"github.com/dhowden/mindmeld/pb"
)

func TestProxyClientWrite(t *testing.T) {
	const msg = "hello world!"

	done := make(chan struct{})

	ps := protoproxy.NewServer()
	go func() {
		c, err := ps.Accept()
		if err != nil {
			t.Errorf("Accept(): %v", err)
			return
		}
		t.Logf("conection received")

		got := &bytes.Buffer{}
		if n, err := io.Copy(got, c); err != nil {
			t.Errorf("copy failed after reading %d bytes: %v", n, err)
		}

		if got.String() != msg {
			t.Errorf("got %q, want %q", got.String(), msg)
		}

		close(done)
	}()

	s := grpc.NewServer()
	pb.RegisterProxyServiceServer(s, ps)
	ts := NewTestServer(s)

	cc := ts.ClientConn(t)
	c, err := protoproxy.Dial(cc)
	if err != nil {
		t.Errorf("could not dial server: %v", err)
	}

	if n, err := io.WriteString(c, msg); err != nil {
		t.Errorf("write failed (after %d bytes): %v", n, err)
	}
	c.Close()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("timedout waiting for done")
	}
}

func TestProxyClientRead(t *testing.T) {
	const msg = "hello world!"

	ps := protoproxy.NewServer()
	go func() {
		c, err := ps.Accept()
		if err != nil {
			t.Errorf("Accept(): %v", err)
			return
		}
		t.Logf("conection received")

		if n, err := io.WriteString(c, msg); err != nil {
			t.Errorf("write failed (after %d bytes): %v", n, err)
		}
		c.Close()
	}()

	s := grpc.NewServer()
	pb.RegisterProxyServiceServer(s, ps)
	ts := NewTestServer(s)

	cc := ts.ClientConn(t)
	c, err := protoproxy.Dial(cc)
	if err != nil {
		t.Errorf("could not dial server: %v", err)
	}

	got := &bytes.Buffer{}
	if n, err := io.Copy(got, c); err != nil {
		t.Errorf("copy failed after reading %d bytes: %v", n, err)
	}

	if got.String() != msg {
		t.Errorf("got %q, want %q", got.String(), msg)
	}
}
