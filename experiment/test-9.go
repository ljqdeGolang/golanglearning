package main

import (
	"fmt"
	"os"
)


func test9() {
	fp, _ := os.OpenFile("test9.ini",os.O_RDWR,0666)
	byt := []byte("I am")
	_,err := fp.WriteAt(byt,13)
	if err != nil {
		fmt.Println("wrong")
	}
}