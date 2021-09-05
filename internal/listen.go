package internal

import (
	"fmt"
	"io"
	"log"
	"net"
)

func NewListenPlex(l net.Listener) *ListenPlex {
	lp := &ListenPlex{
		Listener: l,
		m:        make(map[byte]*listener),
	}
	go lp.loop()
	return lp
}

type ListenPlex struct {
	net.Listener

	m map[byte]*listener
}

func (l *ListenPlex) loop() {
	n := 0
	for {
		c, err := l.Listener.Accept()
		if err != nil {
			for _, v := range l.m {
				go func(x *listener) {
					x.ch <- accept{err: err}
				}(v)
			}
			// Should maybe stop now!  First error typically kills the whole thing?
			log.Printf("ListenPlex: Accept() returned error (%v), stopping loop...", err)
			break
		}

		log.Printf("Received incoming connection (#%d)...", n)
		go func(x net.Conn) {
			var b [1]byte
			if _, err := io.ReadFull(x, b[:]); err != nil {
				log.Printf("Could not read byte from incoming connection (#%d): %v", n, err)
				if err := c.Close(); err != nil {
					log.Printf("Could not close incoming connection (#%d): %v", n, err)
				}
				return
			}
			log.Printf("Incoming connection identifier (#%d): %c", n, b)
			l.m[b[0]].ch <- accept{c: x}
			log.Printf("Accept done handling connection (#%d)", n)
		}(c)
	}
}

type accept struct {
	c   net.Conn
	err error
}

// NewListener creates a listener that will accept connections with the
// specified byte prefix.
func (l *ListenPlex) NewListener(prefix byte) net.Listener {
	ll := &listener{
		prefix: prefix,
		lp:     l,
		ch:     make(chan accept),
	}
	l.m[prefix] = ll
	return ll
}

type listener struct {
	prefix byte
	lp     *ListenPlex
	ch     chan accept
}

func (l *listener) Accept() (net.Conn, error) {
	log.Printf("[%c] Accept() waiting...", l.prefix)
	x := <-l.ch
	log.Printf("[%c] Accept() got it...", l.prefix)
	return x.c, x.err
}

func (l *listener) Close() error { return l.lp.Listener.Close() }
func (l *listener) Addr() net.Addr {
	log.Printf("listener.Addr")
	return l.lp.Listener.Addr()
}

func DialPlex(network, address string, b byte) (net.Conn, error) {
	log.Printf("DialPlex: dialing...")
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	log.Printf("DialPlex: writing identifier byte...")
	var x [1]byte
	x[0] = b
	if n, err := c.Write(x[:]); err != nil || n != 1 {
		return nil, fmt.Errorf("could not write plex identifier: %w", err)
	}
	log.Printf("DialPlex: done")
	return c, err
}
