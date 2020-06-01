package builder

/*
	建造者模式：用来构建复杂类。将一个复杂的类的构建与类本身分离。将变与不变分离开。
*/
type Builder interface {
	SetName(name string) Builder
	SetArms(arms string) Builder
	Build() *Character
}

type Director struct {
	builder Builder
}

func (d *Director) Create(name, arms string) *Character {
	return d.builder.SetArms(arms).SetName(name).Build()
}

type CharacterBuilder struct {
	character *Character
}

func (c *CharacterBuilder) SetName(name string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.setName(name)
	return c
}

func (c *CharacterBuilder) SetArms(arms string) Builder {
	if c.character == nil {
		c.character = &Character{}
	}
	c.character.setArms(arms)
	return c
}

func (c *CharacterBuilder) Build() *Character {
	if c.character == nil {
		c.character = &Character{}
	}
	return c.character
}

type Character struct {
	Name string
	Arms string
}

func (c *Character) setName(name string) {
	c.Name = name
}

func (c *Character) setArms(arms string) {
	c.Arms = arms
}

func (c *Character) getName() string {
	return c.Name
}

func (c *Character) getArms() string {
	return c.Arms
}
