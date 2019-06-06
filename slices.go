package main

import "fmt"

func main() {
	//切片
	s := make([]string,3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//在切片区域后面添加，不影响切片创建值
	fmt.Println("len:", len(s))
	c := make([]string, len(s))
	copy(c,s)
	s = append(s,"a")
	s = append(s,"c")
	s = append(s,"d")
	fmt.Println(c,s)
	//copy(c, s)//切片copy相同长度的切片，不相同为空
	fmt.Println("len:", len(s))
	l := s[2:5]	//例如，该获取的元素的切片s[2]，s[3]和s[4]。切成（但不包括）s[5]。
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	twoD := make([][]int, 4)   //创建二维数组切片，4个二维值
	for i := 0; i < 4; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
