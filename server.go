package mindmeld

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dhowden/mindmeld/internal"
	"github.com/dhowden/mindmeld/pb"
)

// NewTokenSource creates a new TokenSource.
func NewTokenSource() *TokenSource {
	return &TokenSource{}
}

// TokenSource is a simple token generator.
type TokenSource struct {
	mu sync.Mutex

	buffer [32]byte
	n      uint64
}

// Token returns a new token.
func (t *TokenSource) Token() string {
	defer t.mu.Unlock()
	t.mu.Lock()

	rand.Read(t.buffer[:24])
	binary.LittleEndian.PutUint64(t.buffer[24:], t.n)
	t.n++

	return hex.EncodeToString(t.buffer[:])
}

var _ pb.ControlServiceServer = (*Server)(nil)

// service represents a service, and handles incoming
// forwards via in.
type service struct {
	name string

	in chan *forward
}

func newService(name string) *service {
	return &service{
		name: name,
		in:   make(chan *forward),
	}
}

func (s *service) waitForFwd() <-chan *forward {
	return s.in
}

func (s *service) String() string {
	return fmt.Sprintf("svc[name:%q]", s.name)
}

// forward is a handler for incoming forwards.
type forward struct {
	service string

	conn net.Conn
}

func newForward(service string) *forward {
	return &forward{
		service: service,
	}
}

func (f *forward) String() string {
	return fmt.Sprintf("fwd[service:%q]", f.service)
}

// NewServer creates a new Server.
func NewServer(proxyDial string) *Server {
	return &Server{
		ts:            NewTokenSource(),
		proxyDial:     proxyDial,
		services:      make(map[string]*service),
		serviceTokens: make(map[string]chan net.Conn),
		forwardTokens: make(map[string]*forward),
	}
}

func RegisterServer(gs *grpc.Server, s *Server) {
	pb.RegisterControlServiceServer(gs, s)
}

// Server.
type Server struct {
	proxyDial string
	ts        *TokenSource

	mu            sync.RWMutex        // protects services, serviceToken and forwardTokens
	services      map[string]*service // name -> service
	serviceTokens map[string]chan net.Conn
	forwardTokens map[string]*forward

	doneOnce sync.Once
	done     chan bool

	pb.UnimplementedControlServiceServer
}

// serviceFromToken identifies an incoming proxy connection as coming from
// a service.
func (s *Server) serviceFromToken(token string) (chan<- net.Conn, bool) {
	defer s.mu.Unlock()
	s.mu.Lock()

	ch, ok := s.serviceTokens[token]
	if ok {
		delete(s.serviceTokens, token)
	}
	return ch, ok
}

// forwardFromToken identifies an incoming proxy connection as coming from a
// forward request.
func (s *Server) forwardFromToken(token string) (*forward, bool) {
	defer s.mu.Unlock()
	s.mu.Lock()

	fwd, ok := s.forwardTokens[token]
	if ok {
		delete(s.forwardTokens, token)
	}
	return fwd, ok
}

// ProxyListen starts the net.Listener l and handles the incoming connections
// as proxy connections.
func (s *Server) ProxyListen(l net.Listener) error {
	for {
		c, err := l.Accept()
		if err != nil {
			return fmt.Errorf("could not accept incoming connection: %w", err)
		}

		go s.handleProxyConn(c)
	}
}

func (s *Server) handleProxyConn(c net.Conn) {
	h := &pb.Header{}
	if err := internal.ReadHeader(c, h); err != nil {
		log.Printf("Could not read header from incoming proxy connection: %v", err)
		c.Close()
		return
	}
	token := h.GetToken()

	if ch, ok := s.serviceFromToken(token); ok {
		select {
		case ch <- c:
		case <-s.done:
			log.Printf("Could not connect service: server closed")
		}
		return
	}

	if fwd, ok := s.forwardFromToken(token); ok {
		fwd.conn = c

		// Lookup the service for this forward (check that it's still available).
		svc, ok := s.services[fwd.service]
		if !ok {
			log.Printf("No service for forward %v", fwd)
			c.Close()
			return
		}

		select {
		case svc.in <- fwd:
		case <-s.done:
			log.Printf("Could not connect forward: server closed")
		}
		return
	}

	log.Printf("Unknown token: %q", token)
	c.Close()
}

func (s *Server) getOrCreateService(name string) (*service, bool) {
	defer s.mu.Unlock()
	s.mu.Lock()

	if _, ok := s.services[name]; ok {
		return nil, false
	}

	svc := newService(name)
	s.services[name] = svc
	return svc, true
}

func (s *Server) getService(name string) (*service, bool) {
	defer s.mu.RUnlock()
	s.mu.RLock()

	svc, ok := s.services[name]
	return svc, ok
}

func (s *Server) deleteService(name string) {
	defer s.mu.Unlock()
	s.mu.Lock()

	delete(s.services, name)
}

func (s *Server) createServiceToken() (token string, ch <-chan net.Conn) {
	token = s.ts.Token()
	cch := make(chan net.Conn)

	s.serviceTokens[token] = cch
	return token, cch
}

func (s *Server) createForwardToken(service string) string {
	t := s.ts.Token()

	defer s.mu.Unlock()
	s.mu.Lock()

	s.forwardTokens[t] = newForward(service)
	return t
}

// Create a service.
func (s *Server) CreateService(r *pb.CreateServiceRequest, css pb.ControlService_CreateServiceServer) error {
	name := r.GetName()
	svc, ok := s.getOrCreateService(name)
	if !ok {
		return status.Errorf(codes.AlreadyExists, "service %q already exists", name)
	}

	defer func() {
		s.deleteService(name)
	}()

	log.Printf("Created service %q", name)

	ctx := css.Context()
	for {
		select {
		case fwd, ok := <-svc.waitForFwd(): // forwarding request to the service
			if !ok {
				log.Printf("service is closed")
				return status.Errorf(codes.Unknown, "service is closed")
			}

			// Setup the service token to wait for the incoming connection
			// from the service.
			serviceToken, ch := s.createServiceToken()

			// Send the token to the service client, telling them to create a
			// connection to serve the forward.
			if err := css.Send(&pb.CreateServiceResponse{
				Token:    serviceToken,
				DialAddr: s.proxyDial,
			}); err != nil {
				return status.Errorf(codes.Unknown, "could not send service response: %v", err)
			}

			go s.handleForward(ctx, fwd, ch)

		case <-ctx.Done():
			return ctx.Err()

		case <-s.done:
			return status.Errorf(codes.Unavailable, "server closed")
		}
	}
}

func (s *Server) handleForward(ctx context.Context, fwd *forward, ch <-chan net.Conn) {
	defer fwd.conn.Close()

	var serviceConn net.Conn
	select {
	case serviceConn = <-ch: // wait for the service connection to arrive
		defer serviceConn.Close()

	case <-s.done:
		log.Printf("Could not handle forward: server closed")
		return

	case <-ctx.Done():
		log.Printf("timeout waiting for outgoing service connection")
		return
	}

	if err := copyUpDown(serviceConn, fwd.conn, s.done); err != nil {
		log.Printf("FWD%v: %v", fwd, err)
	}
}

// Forward to remote service.
func (s *Server) ForwardToService(ctx context.Context, r *pb.ForwardToServiceRequest) (*pb.ForwardToServiceResponse, error) {
	name := r.GetName()
	if _, ok := s.getService(name); !ok {
		return nil, status.Errorf(codes.NotFound, "service %q does not exist", name)
	}

	token := s.createForwardToken(name)
	return &pb.ForwardToServiceResponse{
		Token:    token,
		DialAddr: s.proxyDial,
	}, nil
}

// Close shutsdown any running forwards.
func (s *Server) Close() error {
	s.doneOnce.Do(func() { close(s.done) })
	return nil
}
