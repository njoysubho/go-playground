package main

import "fmt"

type Animal interface {
	Say() string
}

type Man struct {
}

type Dog struct {
}

func (m Man) Say() string {
	return "speak"
}

func (d Dog) Say() string {
	return "bark"
}
func speak(animal Animal) {
	fmt.Println(animal.Say())
}

func main() {
	m := Man{}
	d := Dog{}
	speak(m)
	speak(d)
}
