package main

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"

}

// func main() {
// 	m := map[string]*Student{"people": {"zhoujielun"}}
// 	m["people"].name = "wuyanzu"
// }
