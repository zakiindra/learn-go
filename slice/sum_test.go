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

func TestSumAllTails(t *testing.T) {

	assertEquals := func(t *testing.T, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v but got actual %v", expected, actual)
		}
	}

	t.Run("sum of tail of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1,2}, []int{0,9})
		expected := []int{2,9}

		assertEquals(t, actual, expected)
	})

	t.Run("safe sum of empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3,4,5})
		expected := []int{0,9}

		assertEquals(t, actual, expected)
	})
}
