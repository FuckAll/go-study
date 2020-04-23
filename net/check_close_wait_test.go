package net

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
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
	err = server.Serve(LimitListener(listener, 1))
	//err = server.Serve(listener)
	if err != nil {
		fmt.Println("服务器错误2")
	}
}

var httpclient *http.Client

func TestClient(t *testing.T) {
	//for i := 0; i < 1000; i++ {
	//	req, err := http.NewRequest("POST", "http://127.0.0.1:8880/hello", strings.NewReader("{}"))
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	// 长连接
	//	req.Close = false
	//	httpclient = &http.Client{}
	//	resp, err := httpclient.Do(req)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		fmt.Println(string(body))
	//		fmt.Println(err)
	//	}
	//
	//}

	//time.Sleep(time.Second * 1000)
	for i := 0; i < 1000; i++ {
		//time.Sleep(1 * time.Second)
		err := Dial("127.0.0.1:8880", time.Millisecond*500)
		if err != nil {
			fmt.Println("this is i:", i)
		}
		//time.Sleep(100 * time.Second)
	}
	fmt.Println("out")

	time.Sleep(time.Second * 1000)
}

// Dial dial 指定的端口，失败返回错误
func Dial(addr string, timeout time.Duration) error {
	//conn, err := net.DialTimeout("tcp", addr, timeout)
	conn, err := net.Dial("tcp", addr)
	fmt.Println("dial err------>", err)
	if err != nil {
		return err
	}
	fmt.Println("conn:", &conn)
	//defer conn.Close()
	return err
}

func TestChan(t *testing.T) {
	l := make(chan struct{})
	go func() {
		select {
		case <-l:
			fmt.Println("in_in_in")
		}
	}()
	time.Sleep(time.Second * 3)

	close(l)
	time.Sleep(time.Second * 100)

}
