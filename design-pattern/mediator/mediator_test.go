package mediator_test

import (
	//"github.com/go-study/src/design-pattern/mediator"
	"github.com/go-study/src/design-pattern/mediator"
	"testing"
)

func TestMediator(t *testing.T) {
	partnerA := mediator.NewPartnerA(10)
	partnerB := mediator.NewPartnerB(10)

	mediator := mediator.NewMediator(partnerA, partnerB)

	// A 赢取 5元
	mediator.AWin(5)
	if partnerA.GetMoney() != 15 && partnerB.GetMoney() != 5 {
		t.Errorf("A money should be 15, b 5, but a %d, b %d", partnerA.GetMoney(), partnerB.GetMoney())
	}

	// A 赢取 5元
	mediator.AWin(5)
	if partnerA.GetMoney() != 20 && partnerB.GetMoney() != 0 {
		t.Errorf("A money should be 20, b 0, but a %d, b %d", partnerA.GetMoney(), partnerB.GetMoney())
	}

	// B 赢取 10元
	mediator.BWin(10)
	if partnerA.GetMoney() != 10 && partnerB.GetMoney() != 10 {
		t.Errorf("A money should be 10, b 10, but a %d, b %d", partnerA.GetMoney(), partnerB.GetMoney())
	}
}
