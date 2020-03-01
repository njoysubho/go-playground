package main

import "fmt"

func main() {
	i := 42
	x := &i
	type mystruct struct{ x, y int }
	fmt.Println(x)
	fmt.Println(*x)
	y := mystruct{1, 3}
	z := &y
	fmt.Println(z.x)
}
