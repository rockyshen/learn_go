package main

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Haha interface {
	Area() float64
	Perimeter() float64
}

// 圆形结构体
type Circle struct {
	Radius float64
}

// 面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 矩形结构体
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 接受接口作为参数的函数
func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	// 声明一个圆
	circle := Circle{Radius: 5}
	// 声明一个矩形
	rectangle := Rectangle{Width: 4, Height: 6}

	// 多态
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}
