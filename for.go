package main

import "fmt"

func main() {
	for i:=7;i <=9;i++ {
		fmt.Println(i)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
