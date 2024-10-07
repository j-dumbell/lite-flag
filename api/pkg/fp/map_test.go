package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		actual := Map(ints, square)
		expected := []int{1, 4, 9}
		assert.Equal(t, expected, actual)
	})

	t.Run("empty slice", func(t *testing.T) {
		arr := []int{}
		actual := Map(arr, square)
		expected := []int{}
		assert.Equal(t, expected, actual)
	})

	t.Run("string -> struct", func(t *testing.T) {
		arr := []string{"a", "b", "c"}
		actual := Map(arr, stringToStruct)
		expected := []someStruct{{"a"}, {"b"}, {"c"}}
		assert.Equal(t, expected, actual)
	})
}
