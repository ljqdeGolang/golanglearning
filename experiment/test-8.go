package main

import (
	"fmt"	
	"time"
)


func test8() {
	fmt.Println("1")
	time.Sleep(5 * time.Second)
	fmt.Println("2")
	time.Sleep(time.Duration(8) * time.Second)
	fmt.Println("3") 	
}