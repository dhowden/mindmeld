package mindmeld

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld/internal"
	"github.com/dhowden/mindmeld/pb"
)

// ServiceClient registers services which will be forwarded through this
// process.
type ServiceClient struct {
	cc *grpc.ClientConn

	name, target string

	doneOnce sync.Once
	done     chan bool
}

// NewServiceClient creates a new ServiceClient.
func NewServiceClient(cc *grpc.ClientConn, name, target string) *ServiceClient {
	return &ServiceClient{
		cc:     cc,
		name:   name,
		target: target,
		done:   make(chan bool),
	}
}

// Register the service and run the service.  Running connections will
// continue to operate after Register returns (even with non-nil error).
// Call Close to shutdown all running connections.
func (sc *ServiceClient) Register(ctx context.Context) error {
	log.Printf("Creating service %q forwarding to %q...", sc.name, sc.target)
	msc := pb.NewControlServiceClient(sc.cc)

	csc, err := msc.CreateService(ctx, &pb.CreateServiceRequest{
		Name: sc.name,
	})
	if err != nil {
		return fmt.Errorf("could not create service: %w", err)
	}

	for {
		resp, err := csc.Recv()
		if err != nil {
			if err == io.EOF {
				// Ended peacefully!
				return nil
			}
			return fmt.Errorf("could not receive: %v", err)
		}
		go sc.handleConn(resp.GetToken(), resp.GetDialAddr())
	}
}

func (sc *ServiceClient) handleConn(token, dialAddr string) {
	log.Printf("Creating connection to host traffic for forward %q", token)
	c, err := net.Dial("tcp", dialAddr)
	if err != nil {
		log.Printf("Could not dial proxy: %v", err)
		return
	}
	defer c.Close()

	if err := internal.WriteHeader(c, &pb.Header{Token: token}); err != nil {
		fmt.Printf("Could not write header: %v", err)
	}

	fconn, err := net.Dial("tcp", sc.target)
	if err != nil {
		fmt.Printf("Could not dial %q for incoming connection: %v", sc.target, err)
		return
	}
	defer fconn.Close()

	if err := copyUpDown(fconn, c, sc.done); err != nil {
		log.Printf("Forwarding ended: %v", err)
	}
}

// Close shuts down all running connections.
func (sc *ServiceClient) Close() error {
	sc.doneOnce.Do(func() {
		close(sc.done)
	})
	return nil
}

func NewForwardClient(cc *grpc.ClientConn, service, localAddr string) *ForwardClient {
	return &ForwardClient{
		cc:        cc,
		service:   service,
		localAddr: localAddr,
		done:      make(chan bool),
	}
}

type ForwardClient struct {
	cc *grpc.ClientConn

	service   string
	localAddr string

	doneOnce sync.Once
	done     chan bool
}

func (fc *ForwardClient) Forward() error {
	l, err := net.Listen("tcp", fc.localAddr)
	if err != nil {
		return fmt.Errorf("could not listen for incoming connections: %v", err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return fmt.Errorf("could not accept incoming connection: %v", err)
		}

		go fc.handleConn(c)
	}
}

func (fc *ForwardClient) handleConn(c net.Conn) {
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := pb.NewControlServiceClient(fc.cc).ForwardToService(ctx, &pb.ForwardToServiceRequest{
		Name: fc.service,
	})
	if err != nil {
		log.Printf("Could not forward to service: %v", err)
		return
	}

	// Dial the proxy.
	fconn, err := net.Dial("tcp", resp.GetDialAddr())
	if err != nil {
		log.Printf("Could not dial: %v", err)
	}
	defer fconn.Close()

	// Identify this forward.
	if err := internal.WriteHeader(fconn, &pb.Header{Token: resp.GetToken()}); err != nil {
		log.Printf("Could not write header: %v", err)
		return
	}

	if err := copyUpDown(fconn, c, fc.done); err != nil {
		log.Printf("Forwarding ended: %v", err)
	}
}

// Close shuts down all running connections.
func (fc *ForwardClient) Close() error {
	fc.doneOnce.Do(func() {
		close(fc.done)
	})
	return nil
}

func copyUpDown(up, down io.ReadWriter, done <-chan bool) error {
	errc := make(chan error, 1)
	go cp(up, down, errc)
	go cp(down, up, errc)

	select {
	case err := <-errc:
		return err
	case <-done:
	}
	return nil
}

func cp(w io.Writer, r io.Reader, errCh chan error) {
	_, err := io.Copy(w, r)
	errCh <- err
}
