package gobyexample

import "math"

type interfacesGeometry interface {
	Area() float64
	Perimeter() float64
}

type interfacesRectangle struct {
	width, height float64
}

type interfacesCircle struct {
	radius float64
}

// rectangle interface implementation
func (r interfacesRectangle) Area() float64 {
	return r.width * r.height
}

func (r interfacesRectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// circle interface implementation
func (c interfacesCircle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c interfacesCircle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measureInterfaceGeometry(g interfacesGeometry) {
	println(g.Area())
	println(g.Perimeter())
}

func isInterfaceRectangle(g interfacesGeometry) bool {
	_, ok := g.(interfacesRectangle)
	return ok
}

func isInterfaceCircle(g interfacesGeometry) bool {
	_, ok := g.(interfacesCircle)
	return ok
}

func ShowInterfacesAndTypeAssertions() {
	r := interfacesRectangle{width: 5, height: 3}
	c := interfacesCircle{radius: 2}

	measureInterfaceGeometry(r)
	measureInterfaceGeometry(c)

	if isInterfaceRectangle(r) {
		println("r is a rectangle")
	}

	if isInterfaceCircle(c) {
		println("c is a circle")
	}
}
