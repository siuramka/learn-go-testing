package struct_methods_interfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{height: 10.0, width: 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testcases := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{height: 10.0, width: 10.0}, 100},
		{Square{height: 10.0, width: 12.0}, 120},
		{Triangle{width: 10.0}, 31.4},
	}
	for _, tc := range testcases {
		t.Run("shape area returns 100", func(t *testing.T) {
			got := tc.shape.Area()
			want := tc.want

			if !withinTolerance(got, want) {
				t.Errorf("got %g want %g", got, want)
			}
		})
	}
}

func withinTolerance(a, b float64) bool {
	e := 1e-12
	if a == b {
		return true
	}
	d := math.Abs(a - b)
	if b == 0 {
		return d < e
	}
	return (d / math.Abs(b)) < e
}
