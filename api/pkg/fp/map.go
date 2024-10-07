package fp

func Map[T any, R any](arr []T, fn func(t T) R) []R {
	mapped := make([]R, len(arr))
	for i, t := range arr {
		mapped[i] = fn(t)
	}
	return mapped
}

func MapErr[T any, R any](arr []T, fn func(t T) (R, error)) ([]R, error) {
	mapped := make([]R, len(arr))
	for i, t := range arr {
		r, err := fn(t)
		if err != nil {
			return nil, err
		}
		mapped[i] = r
	}
	return mapped, nil
}
