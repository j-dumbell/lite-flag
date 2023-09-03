package array

import (
	"reflect"
	"testing"
)

// ToDo - where should this live?
func AssertDeepEqual(t *testing.T, actual any, expected any) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestArrMap1(t *testing.T) {
	square := func(i int) int {
		return i * i
	}

	type someStruct struct {
		str string
	}

	stringToStruct := func(s string) someStruct {
		return someStruct{str: s}
	}

	t.Run("int -> int", func(t *testing.T) {
		ints := []int{1, 2, 3}
		actual := ArrMap(ints, square)
		expected := []int{1, 4, 9}
		AssertDeepEqual(t, expected, actual)
	})

	t.Run("empty slice", func(t *testing.T) {
		arr := []int{}
		actual := ArrMap(arr, square)
		expected := []int{}
		AssertDeepEqual(t, expected, actual)
	})

	t.Run("string -> struct", func(t *testing.T) {
		arr := []string{"a", "b", "c"}
		actual := ArrMap(arr, stringToStruct)
		expected := []someStruct{{"a"}, {"b"}, {"c"}}
		AssertDeepEqual(t, expected, actual)
	})
}
