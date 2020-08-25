package sso

import "testing"

func TestSSODecorator(t *testing.T) {
	var handleInterceptor IHandleInterceptor
	handleInterceptor = new(SSOInterceptor)
	handleInterceptor = NewSsoDecorator(handleInterceptor)
	handleInterceptor = NewSsoDecorator(handleInterceptor)
	handleInterceptor.PreHandle("1successhuahua", "ewcdqwt40liuiu", "t")
}
