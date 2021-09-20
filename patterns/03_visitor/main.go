package main

import "github.com/vlasove/materials/tasks_2/patterns/visitor/figures"

func main() {
	circle := figures.NewCircle(10.5)
	rectangle := figures.NewRectangle(2, 4)
	square := figures.NewSquare(10)

	areaCalculator := &figures.AreaCalculatorVisitor{}
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)
	square.Accept(areaCalculator)

	perimeterCalculator := &figures.PerimeterCalculatorVisitor{}
	circle.Accept(perimeterCalculator)
	rectangle.Accept(perimeterCalculator)
	square.Accept(perimeterCalculator)
}
