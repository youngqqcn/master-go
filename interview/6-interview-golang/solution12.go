package main

import "fmt"

type student struct {
	Name string
	Age  int
}

// func pase_student() {
// 	m := make(map[string]*student)
// 	stus := []student{
// 		{Name: "A", Age: 1},
// 		{Name: "B", Age: 2},
// 		{Name: "C", Age: 3},
// 	}
// 	for _, stu := range stus {
// 		m[stu.Name] = &stu
// 	}

// 	for k, v := range m {
// 		fmt.Printf("%v,%v,%v\n", k, v.Name, v.Age)
// 	}
// }

// func pase_student() {
// 	m := make(map[string]student)
// 	stus := []student{
// 		{Name: "A", Age: 1},
// 		{Name: "B", Age: 2},
// 		{Name: "C", Age: 3},
// 	}
// 	for _, stu := range stus {
// 		m[stu.Name] = stu
// 	}

// 	for k, v := range m {
// 		fmt.Printf("%v,%v,%v\n", k, v.Name, v.Age)
// 	}
// }


func pase_student() {
	m := make(map[string]student)
	stus := []student{
		{Name: "A", Age: 1},
		{Name: "B", Age: 2},
		{Name: "C", Age: 3},
	}
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = stus[i]
	}

	for k, v := range m {
		fmt.Printf("%v,%v,%v\n", k, v.Name, v.Age)
	}
}


func main() {
	pase_student()
}
