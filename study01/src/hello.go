package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "85903034-1e04-4633-a903-74a4452a7064"
	st1 := strings.Replace(str, "-", "_", -1)
	fmt.Println(st1)
	//resp, err := http.Get("http://www.baidu.com")
	//if err != nil {
	//fmt.Println("err")
	//}
	//defer resp.Body.Close()
	//fmt.Println(resp.Header.Get("Connection"))
	//resp, err := http.Post("http://www.baidu.com", "image/jpeg", &buf)

	//fmt.Println(resp)
}
