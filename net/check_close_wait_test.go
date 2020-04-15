package main

import (
	"fmt"
	"golang.org/x/net/netutil"
	"net"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.Body = ", r.Body)
		fmt.Fprintf(w, "HelloWorld!")
	})
	server := &http.Server{Addr: "", Handler: mux}
	listener, err := net.Listen("tcp4", "0.0.0.0:8880")
	if err != nil {
		fmt.Println("服务器错误")
	}
	// 这个地方要把限制变成1
	err = server.Serve(netutil.LimitListener(listener, 1))
	if err != nil {
		fmt.Println("服务器错误2")
	}
}
