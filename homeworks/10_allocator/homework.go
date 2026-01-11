package homework10

import "unsafe"

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	if len(memory) == 0 || len(pointers) == 0 {
		return
	}

	base := uintptr(unsafe.Pointer(&memory[0]))

	ptrmap := make(map[int]int)
	for i, v := range pointers {
		ptrmap[int(uintptr(v)-base)] = i
	}

	write := 0
	for read := range memory {
		if i, ok := ptrmap[read]; ok {
			if write != read {
				memory[write] = memory[read]
			}
			pointers[i] = unsafe.Pointer(&memory[write])
			write++
		}
	}

	// Clear tail data
	for i := write; i < len(memory); i++ {
		memory[i] = 0
	}
}
