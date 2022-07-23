package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`  
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}


// 属性或 struct 是私有的，同样，在json 解码或转码的时候也无法上线私有属性的转换。