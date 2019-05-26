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

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("Got: %.2f; Want: %.2f", got, want)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Got: %.2f; Want: %.2f", got, want)
		}
	})

}