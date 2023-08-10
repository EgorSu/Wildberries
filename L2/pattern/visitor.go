package pattern

import (
	"fmt"
)

type Visitor interface {
	rectangleVisitor(*rectangle) string
	circleVisitor(*circle) string
	triangleVisitor(*triangle) string
}

type rectangle struct {
	sideOne int
	sideTwo int
}

func (r *rectangle) accept(v Visitor) {
	fmt.Printf("%s was calculated for a rectangle with sides: %d, %d", v.rectangleVisitor(r), r.sideOne, r.sideTwo)
}

type circle struct {
	radius int
}

func (c *circle) accept(v Visitor) {
	fmt.Printf("%s was calculated for a circle with radius: %d", v.circleVisitor(c), c.radius)
}

type triangle struct {
	sideOne   int
	sideTwo   int
	sideThree int
}

func (t *triangle) accept(v Visitor) {
	fmt.Printf("%s was calculated for a triangle with sides: %d, %d, %d", v.triangleVisitor(t), t.sideOne, t.sideTwo, t.sideThree)
}

type area struct {
}

func (a *area) circleVisitor(c *circle) string {
	return "The area of a circle"
}
func (a *area) triangleVisitor(t *triangle) string {
	return "The area of a triangle"
}
func (a *area) rectangleVisitor(c *rectangle) string {
	return "The area of a rectangle"
}

type perimeter struct {
}

func (p *perimeter) circleVisitor(c *circle) string {
	return "The perimeter of a circle"
}
func (p *perimeter) triangleVisitor(t *triangle) string {
	return " The perimeter of a triangle"
}
func (p *perimeter) rectangleVisitor(c *rectangle) string {
	return "The perimeter of a rectangle"
}
