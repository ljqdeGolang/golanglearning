package main

import "fmt"

func test4() {
	a :=[10]int{0,5,7,9,1}
	A := a[:5]
	for _,n:=range A {
		// switch {
		// 	case n==1:
		// 		fmt.Println("break")
		// 	case n!=1:
		// 		fmt.Println("continue")
		// 		continue
		// }
		// fmt.Println("here")
		if n>1 {
			fmt.Println("break")
			break
		}
		if n<1 {
			fmt.Println("continue")
			continue
		}
		fmt.Println("this")
	}
}