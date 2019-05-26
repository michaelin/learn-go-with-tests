package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got: %.2f; Want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkarea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got: %.2f; Want: %.2f", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkarea(t, rectangle, 72.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkarea(t, circle, 314.1592653589793)
	})

}
