package fp

func ToPtr[T any](t T) *T {
	return &t
}

func FromPtr[T any](t *T) T {
	return *t
}
