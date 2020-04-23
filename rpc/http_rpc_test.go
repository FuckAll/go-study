package rpc

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

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
		line, _, _ := bufio.NewReader(r.Body).ReadLine()

		fmt.Println(string(line))
		err2 := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err2 != nil {
			t.Error(err2)
		}
	})
	http.ListenAndServe(":1234", nil)
}
