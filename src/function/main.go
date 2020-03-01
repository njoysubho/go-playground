package main

import (
	"fmt"
)

func inc(i *int) int {
	return *i + 1
}

func main() {
	i := 0
	j := inc(&i)
	fmt.Println("i= ", i)
	fmt.Println("j= ", j)
}
