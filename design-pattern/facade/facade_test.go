package facade_test

import (
	"fmt"
	"github.com/go-study/src/design-pattern/facade"
	"testing"
)

func TestFacade(t *testing.T) {
	fcd := new(facade.Facade)
	fcd.Facade(new(facade.SubSystemLight), new(facade.SubSystemTelevision), new(facade.SubSystemAirCondition))
	fcd.Sleep()
	fmt.Println("-----------------------------------------")
	fcd.WakeUp()
}
