package main

import (
	"fmt"
	"net/http"
	"strings"
	//"io/ioutil"
	"os"
	"bufio"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.ProtoAtLeast(1, 2))
	fmt.Println(r.UserAgent())
	//new := r.WithContext(context.Background())
	//fmt.Println(new.UserAgent())

	fmt.Println("-----------------")
	//s, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(s))
	//fmt.Println(r.Method)

	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//fmt.Println(r.PostForm)

	//buf := make([]byte, 1024)
	f, _, _ := r.FormFile("cv")
	defer f.Close()

	bufferedReader := bufio.NewReader(f)
	//bufferedReader.ReadString('\n')
	//bufferedReader.WriteTo(f)


	//f.Read(buf)
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	//io.Copy()
	bufferedReader.WriteTo(fout)

	//fout.Write(buf)

	//r.ParseMultipartForm(10 << 20)
	//fmt.Println("++++++++++++++++++")
	//f := r.MultipartForm.File["cv"]
	//r.FormFile()
	//for _, v := range f {
	//	fmt.Println(v.Header)
	//	fmt.Println(v.Filename)
	//
	//	file, _:= v.Open()
	//
	//	buf := make([]byte, 1024)
	//	file.Read(buf)
	//	fmt.Println(string(buf))
	//	fmt.Println(file.)
	//
	//file.Close()
	//}

	//mReader , _:= r.MultipartReader()
	//if mReader == nil {
	//	fmt.Println("mReader is nil")
	//}
	//form, _ := mReader.ReadForm(10<<20)
	//fmt.Println(form)
	//fmt.Println(mReader.NextPart())
	//fmt.Println(mReader.NextPart())

	//for k, v := range r.PostForm{
	//	fmt.Println("kkkk", k )
	//	fmt.Println("vvvv", strings.Join(v ,""))
	//}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	//http.HandleFunc("/", sayhelloName)       //设置访问的路由
	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//p := []byte{'A', 'b'}
	////p[0] = 'A'
	////p[1] = 'b'
	//fmt.Println(p[0] | 0x20)
	//fmt.Println(p[1] | 0x20)


	//var size []string

	//var input string
	//var v1, v2, v3 , v4 string
	//fmt.Scanf("%s %s %s %s", &v1, &v2, &v3, &v4)
	var tmp string
	fmt.Scanln(tmp)
	fmt.Println(tmp)

	//if
	//
	//var size []string

	//fmt.Sprintln(v1, v2, v3, v4)


}
