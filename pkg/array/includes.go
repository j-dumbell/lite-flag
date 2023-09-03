package array

import "reflect"

func Includes[T comparable](arr []T, value T) bool {
	for _, t := range arr {
		if reflect.DeepEqual(t, value) {
			return true
		}
	}
	return false
}
