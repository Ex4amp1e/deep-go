package homework11

import "unsafe"

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

var (
	heapPointer1 = &heapObjects[1]
	heapPointer2 = &heapObjects[2]
	heapPointer3 = (*int)(nil)
	heapPointer4 = &heapPointer3
)
