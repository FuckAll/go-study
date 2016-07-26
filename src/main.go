package main

import "fmt"
import "encoding/json"
import "os"
import "reflect"

func main() {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println(dat)
		for k, v := range dat {
			fmt.Sprintln(k, v)
		}
	}
	// str, err := json.Marshal(dat)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(str))
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)
	fmt.Println(reflect.TypeOf(dat))

}
func test2() {
	type Tes struct {
		good string
	}
	type MyData struct {
		Name  string  `json:"item"`
		Other float32 `json:"amount"`
		Test  Tes
	}

	var detail MyData
	detail.Name = "1"
	detail.Other = 2
	detail.Test = Tes{good: "wonderful"}
	fmt.Println(detail.Test)
	body, err := json.Marshal(detail)
	fmt.Println(string(body))
	if err != nil {
		panic(err.Error())
	}
	var temp MyData
	// var temp map[string]interface{}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v", temp)
}

func test1() {
	test := struct {
		name string
		age  int
	}{name: "wonder", age: 11}
	fmt.Println(test)
	mashle(test)
}

func mashle(test interface{}) {
	message, err := json.Marshal(test)
	if err != nil {
		fmt.Println("wonder")
	}
	fmt.Printf("%s,%v", string(message), err)
}

func test(test interface{}) {
	switch test.(type) {
	case string:
		fmt.Println("string")
	default:
		fmt.Println("stringsssssss")
	}
}

func funcName(a interface{}) {
	value, ok := a.(string)
	if !ok {
		fmt.Println("not ")
		fmt.Println(value)
		return
	}
	fmt.Println("The value is ", value)
}

func assert(test interface{}) {
	switch t := test.(type) {
	default:
		fmt.Printf("unexpected type %T", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("interger %d\n", t)
	}
}
