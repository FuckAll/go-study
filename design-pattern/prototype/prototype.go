package prototype

import "fmt"

/*
	原型模式：创建型设计模式，主要用作一些对象的创建过程复杂，部分创建过程可以复用的场景。（克隆）
*/

type Role interface {
	Clone() Role
	SetHeadColor(headColor string)
	SetEyesColor(eyesColor string)
	SetTall(tall int)
	Show()
}

type RoleChinese struct {
	headColor string
	eyesColor string
	tall      int
}

// 这个是原型模式的关键
func (r *RoleChinese) Clone() Role {
	return &RoleChinese{headColor: r.headColor, eyesColor: r.eyesColor, tall: r.tall}
}

func (r *RoleChinese) SetHeadColor(headColor string) {
	r.headColor = headColor
}

func (r *RoleChinese) SetEyesColor(eyesColor string) {
	r.eyesColor = eyesColor
}

func (r *RoleChinese) SetTall(tall int) {
	r.tall = tall
}

func (r *RoleChinese) Show() {
	fmt.Println("Role's HeadColor is:", r.headColor, " EyesColor is:", r.eyesColor, " tall is:", r.tall)
}
