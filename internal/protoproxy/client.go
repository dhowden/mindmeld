package protoproxy

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld/pb"
)

var _ net.Conn = (*Conn)(nil)

// Dial creates a connection to the proxy service.
func Dial(cc *grpc.ClientConn) (*Conn, error) {
	ps := pb.NewProxyServiceClient(cc)
	x, err := ps.ProxyConnection(context.Background())
	if err != nil {
		return nil, fmt.Errorf("could not init proxy connection: %w", err)
	}

	c := newConn(x)
	return c, nil
}
