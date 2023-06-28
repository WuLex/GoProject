package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	// 发起 HTTP 请求获取网页内容
	resp, err := http.Get("https://blog.csdn.net/admans?type=blog")
	if err != nil {
		fmt.Println("HTTP 请求失败：", err)
		return
	}
	defer resp.Body.Close()

	// 读取网页内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网页内容失败：", err)
		return
	}

	// 使用正则表达式提取网页中的链接
	re := regexp.MustCompile(`<a\s+[^>]*href=["\']([^"\']+)["\']`)
	matches := re.FindAllStringSubmatch(string(body), -1)

	// 打印提取到的链接
	for _, match := range matches {
		fmt.Println(match[1])
	}
}