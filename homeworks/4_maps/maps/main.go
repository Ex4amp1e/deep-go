package main

import (
	"fmt"
	"unsafe"
)

type hmap struct {
	count      int
	flags      uint8
	B          uint8
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      unsafe.Pointer
}

func main() {

	m := make(map[int]int, 8)

	m[0] = 0
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 5
	m[6] = 6
	m[7] = 7
	m[8] = 8
	m[9] = 9
	m[10] = 10
	m[11] = 11
	m[12] = 12
	// m[13] = 13
	// m[14] = 14
	// m[15] = 15
	// m[16] = 16
	// m[17] = 17
	// m[18] = 18
	// m[19] = 19

	clear(m)
	// delete(m, 0)

	h := *(*hmap)(*(*unsafe.Pointer)(unsafe.Pointer(&m)))

	fmt.Printf("count=%d\n", h.count)
	fmt.Printf("B=%d\n", h.B)
	fmt.Printf("flags=%08b\n", h.flags)
	fmt.Printf("hash0=%d\n", h.hash0)
}
