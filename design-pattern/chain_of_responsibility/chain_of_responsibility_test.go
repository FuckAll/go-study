package chain_of_responsibility_test

import (
	"fmt"
	"github.com/go-study/src/design-pattern/chain_of_responsibility"
	"testing"
)

func TestHandlerChain(t *testing.T) {
	chain := new(chain_of_responsibility.HandlerChain)
	chain.SetChain(new(chain_of_responsibility.HandlerOne))
	chain.SetChain(new(chain_of_responsibility.HandlerTwo))
	chain.HandelWhileHandled()
	fmt.Println("---------------------------")
	chain.HandelAll()

}
