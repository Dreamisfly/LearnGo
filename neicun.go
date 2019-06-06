package main

import "fmt"

func main() {
	var a int = 10
	var b = 123
	type data struct {
		d int
		c float32
	}
	var q = data{1234,1.23213}
	var p  *data
	p = &q

	fmt.Printf("%p, %v\n", p, p.c) // 直接用指针访问目标对象成员，无须转换。

	//声明指针c
	var c *int
	c = &b

	//使用指针访问b的值
	fmt.Print(*c,"\n")

	//获取变量b的值
	fmt.Print(b,"\n")

	//获取变量的地址  %x
	fmt.Printf("变量的地址:%x\n",&a)
}
