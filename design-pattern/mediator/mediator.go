package mediator

/*
	中介模式：在现实生活中，有很多中介者模式的身影，例如QQ游戏平台，聊天室、QQ群、短信平台和房产中介。
	不论是QQ游戏还是QQ群，它们都是充当一个中间平台，QQ用户可以登录这个中间平台与其他QQ用户进行交流，如果没有这些中间平台，
	我们如果想与朋友进行聊天的话，可能就需要当面才可以了。电话、短信也同样是一个中间平台，有了这个中间平台，
	每个用户都不要直接依赖与其他用户，只需要依赖这个中间平台就可以了，一切操作都由中间平台去分发。
	例如：此时，两个牌友打牌，其中有个荷官帮忙发牌，并且负责分配输赢的筹码
*/

type ICardPartner interface {
	ChangeMoney(money int, partner ICardPartner)
}

type PartnerA struct {
	money int
}

func NewPartnerA(money int) *PartnerA {
	return &PartnerA{money: money}
}

func (p *PartnerA) ChangeMoney(money int, partner ICardPartner) {
	p.money += money
	if partner == nil {
		return
	}
	partner.ChangeMoney(-money, nil)
}

func (p *PartnerA) GetMoney() int {
	return p.money
}

type PartnerB struct {
	money int
}

func NewPartnerB(money int) *PartnerB {
	return &PartnerB{money: money}
}

func (p *PartnerB) ChangeMoney(money int, partner ICardPartner) {
	p.money += money
	if partner == nil {
		return
	}
	partner.ChangeMoney(-money, nil)
}

func (p *PartnerB) GetMoney() int {
	return p.money
}

type AbstractMediator interface {
	AWin(money int)
	BWin(money int)
}

type Mediator struct {
	a *PartnerA
	b *PartnerB
}

func NewMediator(a *PartnerA, b *PartnerB) *Mediator {
	return &Mediator{a: a, b: b}
}

func (m *Mediator) AWin(money int) {
	m.a.ChangeMoney(money, m.b)
}

func (m *Mediator) BWin(money int) {
	m.b.ChangeMoney(money, m.a)
}
