package homework1

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
