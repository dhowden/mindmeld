package protoproxy

import (
	"bytes"
	"errors"
	"io"
	"net"
	"sync"
	"time"

	"github.com/dhowden/mindmeld/pb"
)

// Number of individual write/read requests to buffer at a time.
const bufferSize = 10

// PayloadStream represents the key methods required from gRPC streaming
// to emulate a connection.
type PayloadStream interface {
	// Send a message on the stream, translates to a Write.
	Send(*pb.Payload) error

	// CloseSend marks the end of sends.  Only implemented on client streams,
	// so we need to mock it for server streams.
	CloseSend() error

	// Recv a message from the stream, translates to a Read.
	Recv() (*pb.Payload, error)
}

func newConn(s PayloadStream) *Conn {
	c := &Conn{
		r: newReadBuffer(bufferSize),
		w: newWriteBuffer(bufferSize),
	}
	go c.loop(s)
	return c
}

// Conn wraps a PayloadStream into a net.Conn.
type Conn struct {
	r *readBuffer
	w *writeBuffer
}

func (c *Conn) LocalAddr() net.Addr  { return nil }
func (c *Conn) RemoteAddr() net.Addr { return nil }

func (c *Conn) SetDeadline(_ time.Time) error      { return errors.New("not supported") }
func (c *Conn) SetReadDeadline(_ time.Time) error  { return errors.New("not supported") }
func (c *Conn) SetWriteDeadline(_ time.Time) error { return errors.New("not supported") }

// Write some bytes, implements io.Writer. Translates to a Send on the PayloadStream.
func (c *Conn) Write(b []byte) (n int, err error) { return c.w.Write(b) }

// Read some bytes, implements io.Reader. Translates to a Recv on the PayloadStream.
func (c *Conn) Read(b []byte) (n int, err error) { return c.r.Read(b) }

// Close the connection.  Translates to a CloseSend on the PayloadStream.
func (c *Conn) Close() error {
	c.w.Close()
	c.r.Close()
	return nil
}

func (c *Conn) loop(s PayloadStream) {
	go func() {
		for {
			p, err := s.Recv()
			c.r.append(p.GetData(), err)
			if err != nil {
				return
			}
		}
	}()

	for data := range c.w.writes() {
		if err := s.Send(&pb.Payload{
			Payload: &pb.Payload_Data{
				Data: data,
			},
		}); err != nil {
			c.w.setErr(err)
			break
		}
	}
	s.CloseSend() // ignore the error for now
}

func newReadBuffer(n int) *readBuffer {
	return &readBuffer{
		ch: make(chan readPayload, n),
	}
}

type readPayload struct {
	data []byte
	err  error
}

type readBuffer struct {
	ch chan readPayload

	done bool

	r *bytes.Reader
}

func (rb *readBuffer) append(b []byte, err error) {
	rb.ch <- readPayload{
		data: b,
		err:  err,
	}
}

func (rb *readBuffer) Read(b []byte) (n int, err error) {
	if rb.done {
		return 0, io.EOF
	}

	if rb.r == nil {
		rp := <-rb.ch
		if rp.err != nil {
			rb.done = true
			return 0, err
		}
		rb.r = bytes.NewReader(rp.data)
	}

	n, err = rb.r.Read(b)
	if err == io.EOF {
		rb.r = nil
	}
	return n, nil
}

func (rb *readBuffer) Close() error {
	rb.done = true
	return nil
}

func newWriteBuffer(n int) *writeBuffer {
	return &writeBuffer{
		ch:     make(chan []byte, n),
		errCh:  make(chan error, 1),
		doneCh: make(chan struct{}),
	}
}

type writeBuffer struct {
	ch    chan []byte
	errCh chan error

	done     bool
	doneOnce sync.Once
	doneCh   chan struct{}
}

func (wb *writeBuffer) setErr(err error) {
	wb.errCh <- err
	close(wb.errCh)
}

func (wb *writeBuffer) writes() <-chan []byte {
	return wb.ch
}

func (wb *writeBuffer) Write(b []byte) (n int, err error) {
	if wb.done {
		return 0, io.EOF
	}

	// Must not retain the original value beyond the call
	// (as per io.Writer).
	b2 := make([]byte, len(b))
	copy(b2, b)

	select {
	case wb.ch <- b2:
		return len(b2), nil

	case err := <-wb.errCh:
		wb.close()
		return 0, err

	case <-wb.doneCh:
		wb.close()
		return 0, io.EOF
	}
}

func (wb *writeBuffer) close() {
	wb.done = true
	wb.doneOnce.Do(func() {
		close(wb.doneCh)
		close(wb.ch)
	})
}

func (wb *writeBuffer) Close() error {
	wb.close()
	return nil
}
