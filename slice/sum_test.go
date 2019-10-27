package slice

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func (t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("expected %d but got actual %d", expected, actual)
		}
	})
}