package sso

type SSOInterceptor struct {
}

func (S *SSOInterceptor) PreHandle(request string, response string, handler interface{}) bool {
	// 模拟获取cookie
	ticket := request[1:8]
	return "success" == ticket
}
