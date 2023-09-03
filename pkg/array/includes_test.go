package array

import "testing"

func TestIncludes(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		actual := Includes([]int{}, 10)
		AssertDeepEqual(t, actual, false)
	})

	t.Run("exists", func(t *testing.T) {
		actual := Includes([]string{"a", "b", "c"}, "c")
		AssertDeepEqual(t, actual, true)
	})

	t.Run("does not exist", func(t *testing.T) {
		actual := Includes([]float32{1.1, 2.2}, 10.3)
		AssertDeepEqual(t, actual, false)
	})
}
