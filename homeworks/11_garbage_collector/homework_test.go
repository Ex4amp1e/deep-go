package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func Trace(stacks [][]uintptr) []uintptr {
	seen := map[uintptr]bool{}
	result := []uintptr{}

	var dfs func(ptr uintptr)
	dfs = func(ptr uintptr) {
		if ptr == 0 || seen[ptr] {
			return
		}
		seen[ptr] = true
		result = append(result, ptr)

		//nolint:govet
		uptr := *(*uintptr)(unsafe.Pointer(ptr))
		dfs(uptr)
	}

	for _, row := range stacks {
		for _, ptr := range row {
			dfs(ptr)
		}
	}
	return result

}

// For test purpose we want to allocate this globally
var heapObjects = []int{
	0x00, 0x00, 0x00, 0x00, 0x00,
}

func TestTrace(t *testing.T) {
	heapPointer1 := &heapObjects[1]
	heapPointer2 := &heapObjects[2]
	heapPointer3 := (*int)(nil)
	heapPointer4 := &heapPointer3

	var stacks = [][]uintptr{
		{
			uintptr(unsafe.Pointer(&heapPointer1)), 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[0])),
			0x00, 0x00, 0x00, 0x00,
		},
		{
			uintptr(unsafe.Pointer(&heapPointer2)), 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[1])),
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[2])),
			uintptr(unsafe.Pointer(&heapPointer4)), 0x00, 0x00, 0x00,
		},
		{
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[3])),
		},
	}

	pointers := Trace(stacks)
	expectedPointers := []uintptr{
		uintptr(unsafe.Pointer(&heapPointer1)),
		uintptr(unsafe.Pointer(&heapObjects[0])),
		uintptr(unsafe.Pointer(&heapPointer2)),
		uintptr(unsafe.Pointer(&heapObjects[1])),
		uintptr(unsafe.Pointer(&heapObjects[2])),
		uintptr(unsafe.Pointer(&heapPointer4)),
		uintptr(unsafe.Pointer(&heapPointer3)),
		uintptr(unsafe.Pointer(&heapObjects[3])),
	}
	assert.ElementsMatch(t, expectedPointers, pointers)
}
