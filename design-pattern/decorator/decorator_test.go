package decorator_test

import (
	"fmt"
	"github.com/go-study/src/design-pattern/decorator"
	"testing"
)

func TestDecorator(t *testing.T) {
	var shape decorator.Shape
	shape = new(decorator.Circle)
	shape.Draw()
	fmt.Println("without color")

	redShapeDecorator := new(decorator.RedShapeDecorator)
	redShapeDecorator.RedShapeDecorator(new(decorator.Circle))
	shape = redShapeDecorator
	shape.Draw()

	redShapeDecorator.RedShapeDecorator(new(decorator.Rectangle))
	shape = redShapeDecorator
	shape.Draw()

	fmt.Println("-----------------------------")
	greenShapeDecorator := new(decorator.GreenShapeDecorator)
	greenShapeDecorator.GreenShapeDecorator(shape)
	greenShapeDecorator.Draw()
}
