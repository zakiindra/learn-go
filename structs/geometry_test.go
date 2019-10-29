package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("expected %.2f but got actual %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		expected float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{shape: Circle{Radius:10}, expected: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		if actual != tt.expected {
			t.Errorf("expected %.2f but got actual %.2f", tt.expected, actual)
		}
	}
}
