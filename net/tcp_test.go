package net

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

func TestTCP(t *testing.T) {

	fmt.Println(int(^uint(0) >> 1))
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	listener, err := net.Listen("tcp", "127.0.0.1:8880")
	if err != nil {
		panic(err)
	}

	//go func() {
	//	//for i := 0; i < 200000; i++ {
	//	//
	//	//	if i%1000 == 0 {
	//	//		time.Sleep(time.Second * 10)
	//	//	}
	//	conn, err := net.Dial("tcp", "127.0.0.1:8880")
	//	if err != nil {
	//		println("dial err:", err.Error())
	//
	//	}
	//	for i := 0; i < 200000; i++ {
	//		_, err := conn.Write([]byte("这是一个完整的包"))
	//		if err != nil {
	//			println("dial err:", err.Error())
	//		}
	//		//fmt.Println("write length:", n)
	//		//}
	//	}
	//}()

	var count int
	for {
		time.Sleep(time.Second * 10)
		conn, err := listener.Accept()
		count++
		fmt.Println("this is count:", count)
		if err != nil {
			println("err accept:" + err.Error())
		}
		go func() {
			result := bytes.NewBuffer(nil)
			var buf [200]byte
			for {
				n, err := conn.Read(buf[0:])
				result.Write(buf[0:n])
				if err != nil {
					if err == io.EOF {
						continue
					} else {
						fmt.Println("read err:", err)
						break
					}
				} else {
					//if result.Len() != 24 {
					//	fmt.Println("xxxxxxxxxxxxxxxxxx")
					//}
					fmt.Println("recv:", result.String())
				}
				result.Reset()
			}
		}()
	}

}
