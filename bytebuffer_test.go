package bytebuffer_test

import (
	bytebuffer "github.com/caleflat/byte-buffer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutInt(t *testing.T) {
	b := bytebuffer.Allocate(8)
	b.PutInt(0x01020304)
	assert.Equal(t, []byte{0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00}, b.Array())

	b.PutInt(4)
	assert.Equal(t, []byte{0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x04}, b.Array())
}

func TestPut(t *testing.T) {
	b := bytebuffer.Allocate(1)
	b.Put(0x01)
	assert.Equal(t, []byte{0x01}, b.Array())
}

func TestWrap(t *testing.T) {
	b := bytebuffer.Wrap([]byte{0x01, 0x02, 0x03, 0x04})
	assert.Equal(t, []byte{0x01, 0x02, 0x03, 0x04}, b.Array())
}

func TestRemaining(t *testing.T) {
	b := bytebuffer.Allocate(4)
	assert.Equal(t, 4, b.Remaining())

	b.Put(0x01)
	assert.Equal(t, 3, b.Remaining())
}

func TestArray(t *testing.T) {
	b := bytebuffer.Allocate(4)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, b.Array())
}

func TestPutIntOverflow(t *testing.T) {
	b := bytebuffer.Allocate(3)
	assert.Panics(t, func() {
		b.PutInt(0x01020304)
	})
}

func TestAllocateNegativeCapacity(t *testing.T) {
	assert.Panics(t, func() {
		bytebuffer.Allocate(-1)
	})
}
