package simplerpc

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

/*
	将RPC封装成HTTP请求，
	测试用例：curl -vvv  http://localhost:1234/jsonrpc -X POST --data '{"method":"path/to/pkg.HelloService.Hello","params":["hello"], "id":0}'
*/
func TestHttpRPCService(t *testing.T) {
	err := rpc.RegisterName(HelloServiceName, &HelloServiceImpl{})
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}
		err2 := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err2 != nil {
			t.Error(err2)
		}
	})
	http.ListenAndServe(":1234", nil)
}
