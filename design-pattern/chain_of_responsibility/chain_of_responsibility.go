package chain_of_responsibility

import "fmt"

/*
	职责链模式 (一种是遇到能处理的就退出，另外一种是无论如何都要全部过一遍)
	还有一种是通过将所有的Handler串联起来实现，不如这种清晰
*/

type Handler interface {
	handel() bool
}

type HandlerOne struct {
}

func (one *HandlerOne) handel() bool {
	fmt.Println("HandlerOne handel")
	return true
}

type HandlerTwo struct {
}

func (two *HandlerTwo) handel() bool {
	fmt.Println("HandlerTwo handel")
	return false
}

type HandlerChain struct {
	chains []Handler
}

func (hc *HandlerChain) SetChain(handler Handler) {
	hc.chains = append(hc.chains, handler)
}

//HandlerChainWhileHandel 当遇到能处理的就退出
func (hc *HandlerChain) HandelWhileHandled() {
	for _, c := range hc.chains {
		if c.handel() {
			break
		}
	}
	fmt.Println("handelWhileHandled complete")
}

func (hc *HandlerChain) HandelAll() {
	for _, c := range hc.chains {
		c.handel()
	}
	fmt.Println("HandlerAll handel complete")
}
