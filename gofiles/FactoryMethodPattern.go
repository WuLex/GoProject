package main

import "fmt"

// Product 是产品接口，定义了产品的方法
type Product interface {
	Use() string
}

// ConcreteProduct 是具体的产品类型
type ConcreteProduct struct{}

// Use 实现了产品接口的方法
func (p *ConcreteProduct) Use() string {
	return "使用具体的产品"
}

// Factory 是抽象工厂接口，定义了创建产品的方法
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactory 是具体的工厂类型
type ConcreteFactory struct{}

// CreateProduct 实现了抽象工厂接口的方法，返回具体的产品实例
func (f *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{}
}

func main() {
	// 创建具体的工厂实例
	factory := &ConcreteFactory{}

	// 使用工厂创建产品
	product := factory.CreateProduct()

	// 使用产品
	result := product.Use()

	// 输出结果
	fmt.Println(result)
}
