package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(input string) string {
	// 将字符串转换为字符数组
	str := []rune(input)

	// 反转字符数组
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	// 将字符数组转换回字符串
	reversed := string(str)
	return reversed
}

func main() {
	// 读取用户输入
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入字符串：")
	input, _ := reader.ReadString('\n')

	// 去除换行符
	input = strings.TrimSuffix(input, "\n")

	// 调用反转字符串函数
	reversed := reverseString(input)

	// 输出反转后的结果
	fmt.Println("反转后的字符串：", reversed)
}
