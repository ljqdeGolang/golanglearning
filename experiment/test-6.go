package main

import "fmt"

func test6() {
	type amzing int
	type astonishing string

	var a = amzing(5)
	var b = astonishing("yes")
	fmt.Println(a,b)
}