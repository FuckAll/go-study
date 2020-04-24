package protobufrpc

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
	"testing"
)

// 简单的一个http/2的服务
func TestGRPCHttp(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})
	h2Handler := h2c.NewHandler(mux, &http2.Server{})
	server := &http.Server{Addr: ":3999", Handler: h2Handler}
	server.ListenAndServe()
}


func TestGRPCHTTPMixture(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	// 这个地方由于没有证书，可以弄个证书试一下
	panic(http.ListenAndServeTLS(":3999","","", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.ProtoMajor !=2 {
			mux.ServeHTTP(w,r)
			return
		}
		log.Println(r.Header.Get("Content-Type"))
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc"){
			log.Println("进来了")
			grpcServer.ServeHTTP(w,r)
			return
		}
		mux.ServeHTTP(w,r)
		return
	})))
}