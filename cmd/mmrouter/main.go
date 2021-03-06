package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld"
	"github.com/dhowden/mindmeld/internal/protoproxy"
)

var (
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

	pps := protoproxy.NewServer()

	s := mindmeld.NewServer(*proxyDial)
	go func() {
		if err := s.ProxyListen(pps); err != nil {
			log.Printf("Listen(): %v", err)
		}
	}()

	gs := grpc.NewServer()
	mindmeld.RegisterServer(gs, s)
	protoproxy.RegisterServer(gs, pps)

	log.Printf("Listening for gRPC control messages on %q...", *proxyBind)
	if err := gs.Serve(l); err != nil {
		log.Printf("gRPC.Serve(): %v", err)
	}
}
