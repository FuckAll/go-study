package build

import (
	"errors"
	"fmt"
	"time"
)

type PlayerParam string
type MonsterParam string

type Monster struct {
	Name string
}

func NewMonsterParam(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

type Player struct {
	Name string
}

func NewPlayerParam(name PlayerParam) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("player dead")
	}
	return Player{Name: string(name)}, nil
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

type EndingA struct {
	Player  Player
	Monster Monster
}

func NewEndingA(p Player, m Monster) EndingA {
	return EndingA{p, m}
}

func (p EndingA) Appear() {
	fmt.Printf("%s defeats %s, world peace!\n", p.Player.Name, p.Monster.Name)
}

type EndingB struct {
	Player  Player
	Monster Monster
}

func NewEndingB(p Player, m Monster) EndingB {
	return EndingB{p, m}
}

func (p EndingB) Appear() {
	fmt.Printf("%s defeats %s, but become monster, world darker!\n", p.Player.Name, p.Monster.Name)
}
