package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 输入要搜索的关键词
	keyword := "golang tutorial"

	// 构建YouTube搜索URL
	searchURL := "https://www.youtube.com/results"
	queryParams := url.Values{}
	queryParams.Set("search_query", keyword)
	searchURL += "?" + queryParams.Encode()

	// 发送HTTP GET请求
	resp, err := http.Get(searchURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 解析HTML页面
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 提取视频信息
	doc.Find(".yt-lockup-content").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find(".yt-lockup-title").Text())
		views := strings.TrimSpace(s.Find(".yt-lockup-meta-info li:first-child").Text())

		fmt.Println("Title:", title)
		fmt.Println("Views:", views)
		fmt.Println("-------------------------")
	})
}