package main

import "fmt"

func main() {

	glodDog := GoldDog{Name: "glodGog"}
	//这里不同调用 Bark()
	//glodDog.Bark()
	glodDog.Swim()

	//转成Dog类型，就可以调用Dog类型方法
	dog := Dog(glodDog)
	dog.Bark()

    strongDog := StrongDog{Name: "strongDog"}
    strongDog.Swim()
    //这里调用的StrongDog Bark()方法
    strongDog.Bark()

    dog2 := Dog(strongDog)
    //这里调用的Dog类型的Bark()方法
    dog2.Bark()

}

type Dog struct {
	Name string
}

func(d *Dog) Bark() {
	fmt.Printf("Dog barking nmae: %s\n", d.Name)
}

//定义一个新类型，这里完全定义一个新的类型
type GoldDog Dog

func (g *GoldDog) Swim() {
	fmt.Printf("GlodDog swim nmae: %s\n", g.Name)
}

type StrongDog Dog

func (s *StrongDog) Bark() {
	fmt.Printf("Strong Dog barking nmae: %s\n", s.Name)
}

func (s *StrongDog) Swim() {
	fmt.Printf("Strong Dog swim nmae: %s\n", s.Name)
}
