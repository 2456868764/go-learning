package main

import "fmt"

func main() {
	var cat fakeCat = fakeCat{Name: "a little cat"}
	cat.MiaoMiao()
}

type Cat struct {
	Name string
}

func (c *Cat) MiaoMiao() {
	fmt.Printf("cat miaomaio name:%s\n", c.Name)
}

type fakeCat = Cat
