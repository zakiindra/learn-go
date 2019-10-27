package slice

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1,2}, []int{0,9})
	expected := []int{3,9}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got actual %v", expected, actual)
	}
}
