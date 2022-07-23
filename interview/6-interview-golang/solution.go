package main

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show) // new 无法初始化Show内部的属性
	s.Param["RMB"] = 10000
}
