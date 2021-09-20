// Конкретные реализации посетителей
package figures

import (
	"log"
	"math"
)

type PerimeterCalculatorVisitor struct {
	Perimeter float64
}

func (pcv *PerimeterCalculatorVisitor) VisitForSquare(s *Square) {
	pcv.Perimeter = float64(s.Side * 4)
	log.Println("Perimeter of square is :", pcv.Perimeter)
}

func (pcv *PerimeterCalculatorVisitor) VisitForCircle(c *Circle) {
	pcv.Perimeter = 2 * math.Pi * c.Radius
	log.Println("Perimeter of circle is :", pcv.Perimeter)
}

func (pcv *PerimeterCalculatorVisitor) VisitForRectangle(r *Rectangle) {
	pcv.Perimeter = float64(r.A+r.B) * 2
	log.Println("Perimeter of rectangle is :", pcv.Perimeter)
}

type AreaCalculatorVisitor struct {
	Area float64
}

func (acv *AreaCalculatorVisitor) VisitForSquare(s *Square) {
	acv.Area = float64(s.Side * s.Side)
	log.Println("Area of square is :", acv.Area)
}

func (acv *AreaCalculatorVisitor) VisitForCircle(c *Circle) {
	acv.Area = math.Pi * c.Radius * c.Radius
	log.Println("Area of circle is :", acv.Area)
}

func (acv *AreaCalculatorVisitor) VisitForRectangle(r *Rectangle) {
	acv.Area = float64(r.A * r.B)
	log.Println("Area of rectangle is :", acv.Area)
}
