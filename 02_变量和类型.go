package main

import "fmt"

func main() {
	// 整数类型
	var intVar int = 42
	var int8Var int8 = 127
	var int64Var int64 = 9223372036854775807

	// 浮点数类型
	var float32Var float32 = 3.14
	var float64Var float64 = 3.14159265359

	// 布尔类型
	var boolVar bool = true

	// 字符串类型
	var stringVar string = "Golang Learning"

	// 类型推断
	shortIntVar := 100

	// 打印变量
	fmt.Printf("整数变量: %d, %d, %d\n", intVar, int8Var, int64Var)
	fmt.Printf("浮点数变量: %f, %f\n", float32Var, float64Var)
	fmt.Printf("布尔变量: %t\n", boolVar)
	fmt.Printf("字符串变量: %s\n", stringVar)
	fmt.Printf("类型推断变量: %d\n", shortIntVar)
}
