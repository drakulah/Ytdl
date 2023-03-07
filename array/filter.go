package array

func Filter[T any](arr []T, less func(e T) bool) []T {
	res := []T{}
	for _, e := range arr {
		if less(e) {
			res = append(res, e)
		}
	}
	return res
}
