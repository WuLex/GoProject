package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // Redis密码，如果没有设置可以为空
		DB:       0,                // Redis数据库索引
	})

	// 检查是否与Redis服务器建立连接
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Redis:", pong)

	// 操作字符串数据
	err = client.Set("name", "John", 0).Err() // 设置键为"name"，值为"John"
	if err != nil {
		log.Fatal(err)
	}
	name, err := client.Get("name").Result() // 获取键为"name"的值
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name:", name)

	// 操作列表数据
	err = client.RPush("fruits", "apple").Err()  // 在列表"fruits"末尾添加元素"apple"
	err = client.RPush("fruits", "banana").Err() // 在列表"fruits"末尾添加元素"banana"
	if err != nil {
		log.Fatal(err)
	}
	fruits, err := client.LRange("fruits", 0, -1).Result() // 获取列表"fruits"的所有元素
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fruits:", fruits)

	// 操作有序集合数据
	err = client.ZAdd("scores", redis.Z{Score: 90, Member: "Alice"}).Err()  // 添加成员"Alice"和对应的分数90到有序集合"scores"
	err = client.ZAdd("scores", redis.Z{Score: 80, Member: "Bob"}).Err()     // 添加成员"Bob"和对应的分数80到有序集合"scores"
	if err != nil {
		log.Fatal(err)
	}
	scores, err := client.ZRangeWithScores("scores", 0, -1).Result() // 获取有序集合"scores"的所有成员和对应的分数
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("scores:")
	for _, score := range scores {
		fmt.Println(score.Member, ":", score.Score)
	}

	// 操作哈希数据
	err = client.HSet("user", "name", "John").Err()     // 设置哈希"user"的字段"name"的值为"John"
	err = client.HSet("user", "age", "30").Err()        // 设置哈希"user"的字段"age"的值为"30"
	err = client.HSet("user", "email", "john@example.com").Err() // 设置哈希"user"的字段"email"的值为"john@example.com"
	
		err = client.HSet("user", "address", "123 Main St").Err() // 设置哈希"user"的字段"address"的值为"123 Main St"
	if err != nil {
		log.Fatal(err)
	}
	user, err := client.HGetAll("user").Result() // 获取哈希"user"的所有字段和对应的值
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:")
	for field, value := range user {
		fmt.Println(field, ":", value)
	}

	// 删除数据
	err = client.Del("name").Err()   // 删除键为"name"的数据
	err = client.Del("fruits").Err() // 删除键为"fruits"的数据
	err = client.Del("scores").Err() // 删除键为"scores"的数据
	err = client.Del("user").Err()   // 删除键为"user"的数据
	if err != nil {
		log.Fatal(err)
	}

	// 关闭Redis客户端连接
	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}
}