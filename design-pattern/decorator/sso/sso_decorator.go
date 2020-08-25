package sso

import "fmt"

type SSODecorator struct {
	handlerInterceptor IHandleInterceptor
}

func NewSsoDecorator(handlerInterceptor IHandleInterceptor) *SSODecorator {
	return &SSODecorator{
		handlerInterceptor: handlerInterceptor,
	}
}

func (s *SSODecorator) PreHandle(request string, response string, handler interface{}) bool {
	// 模拟获取cookie
	fmt.Println("处理装饰器内容")
	return s.handlerInterceptor.PreHandle(request, response, handler)
}
