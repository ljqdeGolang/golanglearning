package main

import (
	"fmt"
	"strings"
)


func test5() {
	a :=`"assad
1	11111"`
	s :=strings.Split(a,"\n")
	fmt.Println(s)
}