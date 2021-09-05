package protoproxy

import (
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld/pb"
)

var _ net.Listener = (*Server)(nil)

type accept struct {
	c   net.Conn
	err error
}

func RegisterServer(s *grpc.Server, x pb.ProxyServiceServer) {
	pb.RegisterProxyServiceServer(s, x)
}

// NewServer creates a new Listener which accepts incoming
// connections from the ProxyService server.
func NewServer() *Server {
	return &Server{
		ch: make(chan accept),
	}
}

type Server struct {
	ch chan accept

	*pb.UnimplementedProxyServiceServer
}

func (s *Server) Accept() (net.Conn, error) {
	x := <-s.ch
	return x.c, x.err
}

func (s *Server) Addr() net.Addr { return nil }
func (s *Server) Close() error   { return nil }

func (s *Server) ProxyConnection(x pb.ProxyService_ProxyConnectionServer) error {
	// The Conn will call CloseSend when it's done writing, but the server stream
	// doesn't have this method (you need to return).  So we wrap the stream and
	// add a method that will emulate this behaviour on the server side.
	ss := newServerStream(x)
	s.ch <- accept{c: newConn(ss)}
	<-ss.closed()
	return nil
}

func newServerStream(x pb.ProxyService_ProxyConnectionServer) *serverStream {
	return &serverStream{
		ProxyService_ProxyConnectionServer: x,
		done:                               make(chan struct{}),
	}
}

type serverStream struct {
	pb.ProxyService_ProxyConnectionServer

	doneOnce sync.Once
	done     chan struct{}
}

func (sc *serverStream) closed() <-chan struct{} {
	return sc.done
}

func (sc *serverStream) CloseSend() error {
	sc.doneOnce.Do(func() {
		close(sc.done)
	})
	return nil
}
