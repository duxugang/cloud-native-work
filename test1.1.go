package main

import "fmt"

/*
*

给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
func main() {
	strs := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("替换前：%v\n", strs)
	for index, str := range strs {
		if str == "stupid" {
			strs[index] = "smart"
		} else if str == "weak" {
			strs[index] = "strong"
		}
	}
	fmt.Printf("替换后：%v", strs)
}
