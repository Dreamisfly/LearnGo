package main

import (
	"fmt"
	"math"
)

//使用const 关键字用来定义常量
const s string = "appropriate"

func main() {
	fmt.Println(s)
	//const 可以出现在任何"var"关键字出现的地方，区别是常量必须有初始值
	const n = 20
	//常量表达式可以执行任意精度数学计算
	const d = 1500 / n
	fmt.Println(d)
	// 数值型常量没有具体类型，除非指定一个类型
	// 比如显式类型转换
	fmt.Println(int64(d))

	// 数值型常量可以在程序的逻辑上下文中获取类型
	// 比如变量赋值或者函数调用。
	// 例如，对于math包中的Sin函数,它需要一个float64类型的变量
	fmt.Println(math.Sin(n))
}

