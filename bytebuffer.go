package bytebuffer

import (
	"encoding/binary"
	"math"
)

// ByteBuffer is a minimal port of the Java ByteBuffer class.
// https://docs.oracle.com/javase/7/docs/api/java/nio/ByteBuffer.html
type ByteBuffer struct {
	buf []byte
	pos int
}

// Allocate creates a new ByteBuffer with the given capacity.
func Allocate(capacity int) *ByteBuffer {
	if capacity < 0 {
		panic("negative capacity")
	}

	return &ByteBuffer{
		buf: make([]byte, capacity),
		pos: 0,
	}
}

// Remaining returns the number of bytes remaining in the buffer.
func (b *ByteBuffer) Remaining() int {
	return len(b.buf) - b.pos
}

// Wrap creates a new ByteBuffer from the given slice.
func Wrap(buf []byte) *ByteBuffer {
	return &ByteBuffer{
		buf: buf,
	}
}

// Array returns the underlying byte slice.
func (b *ByteBuffer) Array() []byte {
	return b.buf
}

// Put writes the given byte to the buffer.
func (b *ByteBuffer) Put(v byte) *ByteBuffer {
	b.buf[b.pos] = v
	b.pos++
	return b
}

// PutInt writes the given int to the buffer.
func (b *ByteBuffer) PutInt(v int) *ByteBuffer {
	if b.Remaining() < 4 {
		panic("buffer overflow")
	}

	b.buf[b.pos] = byte(v >> 24)
	b.buf[b.pos+1] = byte(v >> 16)
	b.buf[b.pos+2] = byte(v >> 8)
	b.buf[b.pos+3] = byte(v)
	b.pos += 4
	return b
}

// PutInt64 writes the given int64 to the buffer.
func (b *ByteBuffer) PutInt64(v int64) *ByteBuffer {
	if b.Remaining() < 8 {
		panic("buffer overflow")
	}

	b.buf[b.pos] = byte(v >> 56)
	b.buf[b.pos+1] = byte(v >> 48)
	b.buf[b.pos+2] = byte(v >> 40)
	b.buf[b.pos+3] = byte(v >> 32)
	b.buf[b.pos+4] = byte(v >> 24)
	b.buf[b.pos+5] = byte(v >> 16)
	b.buf[b.pos+6] = byte(v >> 8)
	b.buf[b.pos+7] = byte(v)
	b.pos += 8
	return b
}

// PutUint writes the given uint to the buffer.
func (b *ByteBuffer) PutUint(v uint) *ByteBuffer {
	if b.Remaining() < 4 {
		panic("buffer overflow")
	}

	b.buf[b.pos] = byte(v >> 24)
	b.buf[b.pos+1] = byte(v >> 16)
	b.buf[b.pos+2] = byte(v >> 8)
	b.buf[b.pos+3] = byte(v)
	b.pos += 4
	return b
}

// PutUint64 writes the given uint64 to the buffer.
func (b *ByteBuffer) PutUint64(v uint64) *ByteBuffer {
	if b.Remaining() < 8 {
		panic("buffer overflow")
	}

	b.buf[b.pos] = byte(v >> 56)
	b.buf[b.pos+1] = byte(v >> 48)
	b.buf[b.pos+2] = byte(v >> 40)
	b.buf[b.pos+3] = byte(v >> 32)
	b.buf[b.pos+4] = byte(v >> 24)
	b.buf[b.pos+5] = byte(v >> 16)
	b.buf[b.pos+6] = byte(v >> 8)
	b.buf[b.pos+7] = byte(v)
	b.pos += 8
	return b
}

// PutString writes the given string to the buffer.
func (b *ByteBuffer) PutString(v string) *ByteBuffer {
	if len(v) > b.Remaining() {
		panic("buffer overflow")
	}

	copy(b.buf[b.pos:], v)
	b.pos += len(v)
	return b
}

// PutSlice writes the given slice to the buffer.
func (b *ByteBuffer) PutSlice(v []byte) *ByteBuffer {
	if len(v) > b.Remaining() {
		panic("buffer overflow")
	}

	copy(b.buf[b.pos:], v)
	b.pos += len(v)
	return b
}

// PutFloat64 writes the given float to the buffer.
func (b *ByteBuffer) PutFloat64(v float64) *ByteBuffer {
	binary.BigEndian.PutUint64(b.buf[b.pos:], math.Float64bits(v))

	return b
}
