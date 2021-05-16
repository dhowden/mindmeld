package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	"github.com/dhowden/mindmeld"
)

var (
	node = flag.String("node", "mindmeld.zed.ai:1111", "mindmeld node")
	mode = flag.String("mode", "listen", "mode to operate: listen|dial")

	serviceName    = flag.String("service-name", "", "service name to use")
	serviceForward = flag.String("service-forward", "", "will forward incoming service traffic to `host:port`")

	forwardFrom = flag.String("forward-from", "localhost:9999", "bind address for listener")
)

func main() {
	flag.Parse()

	grpc.EnableTracing = true

	if *serviceName == "" {
		log.Fatalf("-service must not be empty")
	}

	log.Printf("Dialing mmrouter at %q...", *node)
	cc, err := grpc.Dial(*node, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not dial: %v", err)
	}
	log.Printf("Connected to mmrouter...")

	if *mode == "listen" {
		log.Printf("Creating service %q forwarding to %q...", *serviceName, *serviceForward)
		sc := mindmeld.NewServiceClient(cc, *serviceName, *serviceForward)
		if err := sc.Register(context.Background()); err != nil {
			log.Fatalf("Could not register service: %v", err)
		}
		sc.Close()
		return
	}

	fc := mindmeld.NewForwardClient(cc, *serviceName, *forwardFrom)
	if err := fc.Forward(); err != nil {
		log.Fatalf("Could not forward: %v", err)
	}
	fc.Close()
}
