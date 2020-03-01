package main

import "fmt"

func myDefferedCall() {

	fmt.Println("Deferred is called")
	if p := recover(); p != nil {
		fmt.Println(p)
	}

}
func myFunc() {
	fmt.Println("I am in myFunc")

	defer myDefferedCall()
	var a [2]int
	for i := 0; i <= 3; i++ {
		a[i] = 0
	}
	fmt.Println("I am ending myFunc")
}
func main() {
	fmt.Println("main is started")
	myFunc()
	fmt.Println("All good")
}
