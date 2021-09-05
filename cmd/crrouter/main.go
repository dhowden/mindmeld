package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld"
	"github.com/dhowden/mindmeld/internal/protoproxy"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	dialAddr := os.Getenv("DIAL_ADDR")
	log.Printf("DIAL_ADDR: %q", dialAddr)

	log.Printf("Listening for gRPC traffic on :%v...", port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}

	pps := protoproxy.NewServer()

	s := mindmeld.NewServer(dialAddr)
	go func() {
		if err := s.ProxyListen(pps); err != nil {
			log.Printf("Listen(): %v", err)
		}
	}()

	gs := grpc.NewServer()
	mindmeld.RegisterServer(gs, s)
	protoproxy.RegisterServer(gs, pps)
	if err := gs.Serve(l); err != nil {
		log.Printf("gRPC.Serve(): %v", err)
	}
}
