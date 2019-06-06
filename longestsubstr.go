package main

import "fmt"

func main () {
	var s string = "dadsca"

	var a = lengthOfLongestSubstring(s)

	fmt.Print(a)
}

func lengthOfLongestSubstring(s string) int {
	dataMap := make([]bool,96)
	length := len(s)
	n,m,max := 0,0,0
	//Map := make(map[byte]int)
	for m<length {
		fmt.Print(s[m]-32)
		if dataMap[s[m]-32] {
			dataMap[s[n]-32] = false
			n++
		}else{
			dataMap[s[m]-32] = true
			m++
			if max < m-n{
				max = m-n
			}
		}

	}
	fmt.Print("长度",length)
	fmt.Printf("这是dataMap",dataMap)
	return max
}
