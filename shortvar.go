package main

import (
	"fmt"
)

var c, python, java = true, false, "no!"  //:=不能声明在函数外

func main() {
	k := 3
	c, python, java := false, true, "yes!"
	fmt.Println(c, python, java, k)
	//注意检查，是定义新的局部变量，还是修改全局变量。该方式容易造成错误！
}