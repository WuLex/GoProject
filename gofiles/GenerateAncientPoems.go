package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	verbs   = []string{"行", "坐", "立", "卧", "飞", "游", "闻", "望", "思", "忆"}
	objects = []string{"春风", "夏雨", "秋叶", "冬雪", "明月", "繁星", "青山", "碧水", "红尘", "白云"}
	places  = []string{"山顶", "江畔", "田野", "庭院", "海边", "花丛", "云端", "树下", "窗前", "梦中"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成古诗的行数
	lines := rand.Intn(3) + 2

	// 生成每行古诗的内容
	var poem []string
	for i := 0; i < lines; i++ {
		// 随机选择动词、物体和地点
		verb := getRandomElement(verbs)
		object := getRandomElement(objects)
		place := getRandomElement(places)

		// 组合成一句古诗
		line := fmt.Sprintf("%s%s，%s%s%s。", verb, object, place, object, verb)
		poem = append(poem, line)
	}

	// 输出生成的古诗
	fmt.Println(strings.Join(poem, "\n"))
}

// 从切片中随机选择一个元素
func getRandomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}