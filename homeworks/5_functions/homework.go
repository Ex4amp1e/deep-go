package homework5

func Map[T any](data []T, action func(T) T) []T {
	if data == nil {
		return nil
	}
	res := make([]T, len(data))
	for i := range data {
		res[i] = action(data[i])
	}
	return res
}

func Filter[T any](data []T, action func(T) bool) []T {
	if data == nil {
		return nil
	}
	res := []T{}
	for i := range data {
		if action(data[i]) {
			res = append(res, data[i])
		}
	}
	return res
}

func Reduce[T any](data []T, initial T, action func(T, T) T) T {
	if data == nil {
		return initial
	}
	for _, v := range data {
		initial = action(v, initial)
	}
	return initial
}
