package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dhowden/mindmeld"
	"github.com/dhowden/mindmeld/pb"
)

var (
	node     = flag.String("node", "", "mindmeld node")
	insecure = flag.Bool("insecure", false, "connection to gRPC is insecure")

	mode = flag.String("mode", "listen", "mode to operate: listen|dial|list")

	serviceName    = flag.String("service-name", "", "service name to use")
	serviceForward = flag.String("service-forward", "", "will forward incoming service traffic to `host:port`")

	forwardFrom = flag.String("forward-from", "localhost:9999", "bind address for listener")
)

func main() {
	flag.Parse()

	grpc.EnableTracing = true

	cc, err := dialGRPC(*node, *insecure)
	if err != nil {
		log.Fatalf("Could not dial: %v", err)
	}

	if *mode == "list" {
		resp, err := pb.NewControlServiceClient(cc).ListServices(context.Background(), &pb.ListServicesRequest{})
		if err != nil {
			log.Fatalf("Could not list services: %v", err)
		}

		tw := newTabWriter()
		tw.Writef("CREATED\tNAME\n")
		for _, svc := range resp.GetServices() {
			tw.Writef("%v\t%v\n", svc.GetCreateTime().AsTime().Local().Format(time.Stamp), svc.GetName())
		}
		tw.Flush()
		return
	}

	if *serviceName == "" {
		log.Fatalf("-service must not be empty")
	}

	if *mode == "listen" {
		sc := mindmeld.NewServiceClient(cc, *serviceName, *serviceForward)
		if err := sc.Register(context.Background()); err != nil {
			log.Fatalf("Could not register service: %v", err)
		}
		sc.Close()
		return
	}

	if *mode == "dial" {
		log.Printf("Creating forward from %q to service %q", *forwardFrom, *serviceName)
		fc := mindmeld.NewForwardClient(cc, *serviceName, *forwardFrom)
		if err := fc.Forward(); err != nil {
			log.Fatalf("Could not forward: %v", err)
		}
		fc.Close()
		return
	}
}

func dialGRPC(addr string, insecure bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithBlock()}
	if *node != "" {
		opts = append(opts, grpc.WithAuthority(*node))
	}
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := SystemRoots()
		if err != nil {
			return nil, fmt.Errorf("could not get system certs: %w", err)
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	cc, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not dial: %w", err)
	}
	return cc, nil
}

type tabWriter struct {
	*tabwriter.Writer
}

func (t *tabWriter) reset() {
	t.Init(os.Stdout, 10, 4, 3, ' ', 0)
}

func newTabWriter() *tabWriter {
	return &tabWriter{tabwriter.NewWriter(os.Stdout, 10, 4, 3, ' ', 0)}
}

func (tw *tabWriter) Writef(pattern string, vars ...interface{}) {
	io.WriteString(tw.Writer, fmt.Sprintf(pattern, vars...))
}
