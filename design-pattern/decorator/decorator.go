package decorator

import "fmt"

/*
	装饰器模式：允许向一个现有的对象添加新的功能，同时又不改变其结构。可以对原始类嵌套使用多个装饰器（在设计的时候要跟原始类实现相同的接口或者是抽象类）
*/

type Shape interface {
	Draw()
}

type RedShapeDecorator struct {
	shape Shape
}

func (s *RedShapeDecorator) RedShapeDecorator(shape Shape) {
	s.shape = shape
}

func (s *RedShapeDecorator) Draw() {
	s.shape.Draw()
	fmt.Println("添加边框颜色为红色")
}

type GreenShapeDecorator struct {
	shape Shape
}

func (g *GreenShapeDecorator) GreenShapeDecorator(shape Shape) {
	g.shape = shape
}

func (g *GreenShapeDecorator) Draw() {
	g.shape.Draw()
	fmt.Println("添加边框颜色为绿色")
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Shape: Circle")
}

type Rectangle struct {
}

func (c *Rectangle) Draw() {
	fmt.Println("Shape: Rectangle")
}
