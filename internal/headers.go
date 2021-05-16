package internal

import (
	"encoding/binary"
	"fmt"
	"io"

	"google.golang.org/protobuf/proto"
)

// WriteHeader writes the proto message to the writer.
func WriteHeader(w io.Writer, x proto.Message) error {
	b, err := proto.Marshal(x)
	if err != nil {
		return fmt.Errorf("could not marshal header: %v", err)
	}

	if err := binary.Write(w, binary.LittleEndian, uint32(len(b))); err != nil {
		return fmt.Errorf("could not write header size: %v", err)
	}
	if n, err := w.Write(b); err != nil {
		return fmt.Errorf("could not write header (failed after %d bytes): %v", n, err)
	}
	return nil
}

// ReadHeader read the proto message from the reader.
func ReadHeader(r io.Reader, x proto.Message) error {
	var sizeBuf [4]byte
	n, err := io.ReadFull(r, sizeBuf[:])
	if err != nil {
		return fmt.Errorf("could not read size bytes from header (failed after %d bytes): %w", n, err)
	}
	size := binary.LittleEndian.Uint32(sizeBuf[:])

	buf := make([]byte, size)
	n, err = io.ReadFull(r, buf)
	if err != nil {
		return fmt.Errorf("could not read header message (%d bytes, failed after %d bytes): %w", size, n, err)
	}

	if err := proto.Unmarshal(buf, x); err != nil {
		return fmt.Errorf("could not unmarshal header: %w", err)
	}
	return nil
}
