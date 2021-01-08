package main

import "fmt"

type Animal interface {
	eat()
}

type Cat struct {
	name string
}

func (cat Cat) eat() {
	fmt.Println(cat.name + "猫吃东西")
}

type Dog struct{}

func (dog Dog) eat() {
	fmt.Println("狗吃东西")
}
func main() {
	var animal1 Animal = Cat{"maomao"}
	var animal2 Animal = Dog{}
	animal1.eat()
	animal2.eat()
}
