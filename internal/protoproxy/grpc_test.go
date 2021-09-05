package protoproxy_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// TestServer wraps a running GRPC server and creates clients which
// talk to it using a local buffered connection.
type TestServer struct {
	l        *bufconn.Listener
	shutdown <-chan struct{}

	*grpc.Server
}

// StopTimeout
const StopTimeout = 50 * time.Millisecond

func (s *TestServer) Stop(t *testing.T) {
	s.Server.Stop()
	select {
	case <-s.shutdown:
	case <-time.After(StopTimeout):
		t.Errorf("timeout (%v) waiting for server to shutdown after Stop()", StopTimeout)
	}
}

// ClientConn creates a new connection to the underlying server.  Overrides the dialer
// and sets grpc.WithInsecure.  Calls t.Fatal if the dial fails.
func (s *TestServer) ClientConn(t *testing.T, opts ...grpc.DialOption) *grpc.ClientConn {
	opts = append(opts,
		grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
			return s.l.Dial()
		}),
		grpc.WithInsecure(),
	)
	conn, err := grpc.Dial("bufconn", opts...)
	if err != nil {
		t.Fatalf("unexpected error from grpc.Dial: %v", err)
	}
	return conn
}

// NewTestServer wraps a gRPC server to create a testing server running with
// in-memory networking.
func NewTestServer(s *grpc.Server) *TestServer {
	l := bufconn.Listen(1024 * 1024)
	shutdown := make(chan struct{})
	go func() {
		if err := s.Serve(l); err != nil {
			if err != grpc.ErrServerStopped {
				log.Fatalf("server exited with error: %v", err)
			}
		}
		close(shutdown)
	}()

	return &TestServer{
		l:        l,
		shutdown: shutdown,
		Server:   s,
	}
}
