package fp

func ArrMap[T any, R any](arr []T, fn func(t T) R) []R {
	mapped := make([]R, len(arr))
	for i, t := range arr {
		mapped[i] = fn(t)
	}
	return mapped
}
