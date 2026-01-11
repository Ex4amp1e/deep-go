package homework3

import (
	"runtime"
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) *COWBuffer {
	ref := 1
	buff := &COWBuffer{
		data: data,
		refs: &ref,
	}
	runtime.SetFinalizer(buff, (*COWBuffer).Close)
	return buff
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return COWBuffer{
		data: b.data,
		refs: b.refs,
	}
}

func (b *COWBuffer) Close() {
	runtime.SetFinalizer(b, nil)
	*b.refs--
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if (index >= int(len(b.data))) || (index < 0) {
		return false
	}

	if *b.refs > 1 {
		newData := make([]byte, len(b.data))
		copy(newData, b.data)
		newBuff := NewCOWBuffer(newData)
		*b.refs--
		b.data = newBuff.data
		b.refs = newBuff.refs
	}

	b.data[index] = value

	return true
}

func (b *COWBuffer) String() string {
	return unsafe.String(&b.data[0], len(b.data))
}
