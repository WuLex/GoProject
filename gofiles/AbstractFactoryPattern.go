package main

import "fmt"

// 抽象产品接口
type Shape interface {
	Draw()
}

// 具体产品类型：圆形
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}

// 具体产品类型：正方形
type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}

// 抽象工厂接口
type ShapeFactory interface {
	CreateShape() Shape
}

// 具体工厂类型：圆形工厂
type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

// 具体工厂类型：正方形工厂
type SquareFactory struct{}

func (sf *SquareFactory) CreateShape() Shape {
	return &Square{}
}

// 使用工厂创建产品
func CreateAndDrawShape(factory ShapeFactory) {
	shape := factory.CreateShape()
	shape.Draw()
}

func main() {
	// 使用圆形工厂创建圆形
	circleFactory := &CircleFactory{}
	CreateAndDrawShape(circleFactory)

	// 使用正方形工厂创建正方形
	squareFactory := &SquareFactory{}
	CreateAndDrawShape(squareFactory)
}
