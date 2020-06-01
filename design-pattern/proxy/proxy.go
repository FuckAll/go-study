package proxy

import "fmt"

/*
	代理模式：为对象提供一个代理对象，用来控制对象的访问，追踪对象的日志等。（代理对象不影响正常的被代理对象的功能）
	例如：火车票代售点
*/

type Seller interface {
	Sell(name string)
}

// 火车站
type Station struct {
	ticketsLeft int // 库存
}

func (s *Station) Sell(name string) {
	if s.ticketsLeft > 0 {
		s.ticketsLeft--
		fmt.Printf("火车站售票：%s买了一张票，剩余票数:%d\n", name, s.ticketsLeft)
	} else {
		fmt.Println("票已售空")
	}
}

// 火车售票代理点
type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) Sell(name string) {
	if proxy.station.ticketsLeft > 0 {
		proxy.station.ticketsLeft--
		fmt.Printf("火车售票代理点:%s买了一张票，剩余：%d\n", name, proxy.station.ticketsLeft)
	} else {
		fmt.Println("票已售空")
	}
}
