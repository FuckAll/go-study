package sso

type IHandleInterceptor interface {
	PreHandle(request string, response string, handler interface{}) bool
}
