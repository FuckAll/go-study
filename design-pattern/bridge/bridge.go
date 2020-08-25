package bridge

import "fmt"

/*
	桥接模式： 桥接模式的目的主要是隔离"抽象"与"实现"，抽象并不是具体的接口，实现也并不是具体的对象。方便"抽象"和"实现"独立的变化
	两个不同的维度在"桥"的地方互相组合，避免在两个维度变化直接组合造成的笛卡尔积的复杂度
*/

// 桥接接口，也就是所谓的"抽象"
type DrawAPI interface {
	drawCircle(radius int, x int, y int)
}

// 图形颜色的实现
type RedCircle struct {
}

func (r *RedCircle) drawCircle(radius int, x int, y int) {
	fmt.Printf("red circle radius:%d, x:%d, y:%d\n", radius, x, y)
}

// 桥接接口的实现，也就是所谓的"实现"
type GreenCircle struct {
}

func (g *GreenCircle) drawCircle(radius int, x int, y int) {
	fmt.Printf("green circle radius:%d, x:%d, y:%d\n", radius, x, y)
}

// Shape是抽象，把绘制图形的功能进行抽象
type Shape interface {
	SetDrawAPI(drawApi DrawAPI)
	Draw()
}

// Circle是实现，实现了图形
type Circle struct {
	drawAPI DrawAPI
	x       int
	y       int
	radius  int
}

func (c *Circle) SetDrawAPI(drawAPI DrawAPI) {
	c.drawAPI = drawAPI
}

func NewCircle(x int, y int, radius int) *Circle {
	return &Circle{
		x:      x,
		y:      y,
		radius: radius,
	}
}

func (c *Circle) Draw() {
	c.drawAPI.drawCircle(c.radius, c.x, c.y)
}
