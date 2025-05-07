package lesson7

type Rectangle struct {
}

// old name: NewRectangle
// new name: BySides
func (rectangle *Rectangle) BySides(firstSide, secondSide, thirdSide int) *Rectangle {
	return &Rectangle{}
}

// rectangle := Rectungle.BySides(...)

type Circle struct {
}

// old names: NewCircle1, NewCircle2, NewCircle3
// new names: ByRadius, ByDiametr, ByLength
func (circle *Circle) ByRadius(radius float64) Circle {
	return Circle{}
}
func (circle *Circle) ByDiametr(diametr float64) Circle {
	return Circle{}
}
func (circle *Circle) ByLength(length float64) Circle {
	return Circle{}
}

// circle := Circle.ByRadius(4)
