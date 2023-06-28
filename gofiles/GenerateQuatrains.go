package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	nouns   = []string{"春风", "夏雨", "秋叶", "冬雪", "明月", "繁星", "青山", "碧水", "红尘", "白云"}
	verbs   = []string{"行", "坐", "立", "卧", "飞", "游", "闻", "望", "思", "忆"}
	adjectives = []string{"美丽", "温柔", "忧伤", "宁静", "浪漫", "幽默", "深情", "激情", "自由", "清新"}
	places  = []string{"山顶", "江畔", "田野", "庭院", "海边", "花丛", "云端", "树下", "窗前", "梦中"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成每行绝句的内容
	var poem []string
	for i := 0; i < 4; i++ {
		// 随机选择名词、动词、形容词和地点
		noun := getRandomElement(nouns)
		verb := getRandomElement(verbs)
		adjective := getRandomElement(adjectives)
		place := getRandomElement(places)

		// 组合成一句绝句
		line := fmt.Sprintf("%s%s，%s%s，%s%s%s。", noun, verb, adjective, noun, place, noun, verb)
		poem = append(poem, line)
	}

	// 输出生成的绝句诗
	fmt.Println(strings.Join(poem, "\n"))
}

// 从切片中随机选择一个元素
func getRandomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}