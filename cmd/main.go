package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["1"] = "hi"
	fmt.Println(m["1"])
	fmt.Println(m["2"])
}
