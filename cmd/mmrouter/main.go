package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld"
)

var (
	bind      = flag.String("grpc-bind", "", "host:port gRPC server")
	proxyBind = flag.String("proxy-bind", "", "host:port for TCP proxy")
	proxyDial = flag.String("proxy-dial", "", "dial address for clients to reach TCP proxy")
)

func main() {
	flag.Parse()

	log.Printf("Listening for TCP proxy traffic on %q...", *proxyBind)
	l, err := net.Listen("tcp", *proxyBind)
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}

	s := mindmeld.NewServer(*proxyDial)
	go func() {
		if err := s.ProxyListen(l); err != nil {
			log.Printf("Listen(): %v", err)
		}
	}()

	gs := grpc.NewServer()
	mindmeld.RegisterServer(gs, s)

	log.Printf("Listening for gRPC control messages on %q...", *bind)
	grpcListener, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatalf("gRPC Listener: %v", err)
	}
	if err := gs.Serve(grpcListener); err != nil {
		log.Printf("gRPC.Serve(): %v", err)
	}
}
