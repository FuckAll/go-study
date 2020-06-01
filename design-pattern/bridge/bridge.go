package bridge

import "fmt"

/*
	桥接模式： 桥接模式的目的主要是隔离"抽象"与"实现"，抽象并不是具体的接口，实现也并不是具体的对象。方便"抽象"和"实现"独立的变化
*/

// 桥接接口，也就是所谓的"抽象"
type DrawAPI interface {
	drawCircle(radius int, x int, y int)
}

// 桥接接口的实现，也就是所谓的"实现"
type RedCircle struct {
}

func (r *RedCircle) drawCircle(radius int, x int, y int) {
	fmt.Printf("radius:%d, x:%d, y:%d\n", radius, x, y)
}

// 桥接接口的实现，也就是所谓的"实现"
type GreenCircle struct {
}

func (g *GreenCircle) drawCircle(radius int, x int, y int) {
	fmt.Printf("radius:%d, x:%d, y:%d\n", radius, x, y)
}

// 桥
type Shape struct {
	drawAPI DrawAPI
}

// 初始化
func (s *Shape) Shape(drawAPI DrawAPI) {
	s.drawAPI = drawAPI
}

// 创建实现了Shape接口的实体类。
type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func (c *Circle) Circle(x int, y int, radius int, drawApi DrawAPI) {
	c.shape.Shape(drawApi)
	c.x = x
	c.y = y
	c.radius = radius
}

func (c *Circle) Draw() {
	c.shape.drawAPI.drawCircle(c.radius, c.x, c.y)
}
