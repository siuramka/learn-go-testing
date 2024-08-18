package struct_methods_interfaces

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Square struct {
	width  float64
	height float64
}

type Triangle struct {
	width float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (s Square) Area() float64 {
	return s.width * s.height
}

func (t Triangle) Area() float64 {
	return t.width * 3.14
}
