package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Number interface {
	~uint16 | ~uint32 | ~uint64
}

func ToLittleEndian[T Number](number T) T {
	switch v := any(number).(type) {
	case uint16:
		return T(
			(v & 0x00FF << 8) | (v & 0xFF00 >> 8),
		)
	case uint32:
		return T(
			(v & 0x000000FF << 24) | (v & 0x0000FF00 << 8) | (v & 0x00FF0000 >> 8) | (v & 0xFF000000 >> 24),
		)
	case uint64:
		return T(
			(v & 0x00000000000000FF << 56) |
				(v & 0x000000000000FF00 << 40) |
				(v & 0x0000000000FF0000 << 24) |
				(v & 0x00000000FF000000 << 8) |
				(v & 0x000000FF00000000 >> 8) |
				(v & 0x0000FF0000000000 >> 24) |
				(v & 0x00FF000000000000 >> 40) |
				(v & 0xFF00000000000000 >> 56),
		)
	default:
		panic("unsupported type")
	}
}

func TestConversion16(t *testing.T) {
	tests := map[string]struct {
		number uint16
		result uint16
	}{
		"test case #1": {
			number: 0x0000,
			result: 0x0000,
		},
		"test case #2": {
			number: 0xFFFF,
			result: 0xFFFF,
		},
		"test case #3": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #4": {
			number: 0x0102,
			result: 0x0201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestСonversion32(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversion64(t *testing.T) {
	tests := map[string]struct {
		number uint64
		result uint64
	}{
		"test case #1": {
			number: 0x0000000000000000,
			result: 0x0000000000000000,
		},
		"test case #2": {
			number: 0xFFFFFFFFFFFFFFFF,
			result: 0xFFFFFFFFFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF00FF00FF,
			result: 0xFF00FF00FF00FF00,
		},
		"test case #4": {
			number: 0x0000000000FFFFFF,
			result: 0xFFFFFF0000000000,
		},
		"test case #5": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
