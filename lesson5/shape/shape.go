package shape

type Perimeterable interface {
	Perimeter() int
}

type Areaable interface {
	Area() int
}

type Shaper interface {
	Perimeterable
	Areaable
}

type Circle struct {
	radii int
}

type Square struct {
	side int
}

type Triangle struct {
	side1   int
	side2   int
	side3   int
	height3 int
}

func (c *Circle) Perimeter() int {
	return 2 * 3 * c.radii
}

func (s *Square) Perimeter() int {
	return 4 * s.side
}

func (t *Triangle) Perimeter() int {
	return t.side1 + t.side2 + t.side3
}

func (c *Circle) Area() int {
	return 3 * c.radii * c.radii
}

func (s *Square) Area() int {
	return s.side * s.side
}

func (t *Triangle) Area() int {
	return t.side3 * t.height3
}
