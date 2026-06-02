package main

import "fmt"

// 简单函数
func add(a, b int) int {
	return a + b
}

// 多返回值函数
func divideAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

// 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 函数作为参数和返回值
func applyOperation(operation func(int, int) int, a, b int) int {
	return operation(a, b)
}

func main() {
	// 简单函数调用
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// 多返回值函数
	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Printf("10 / 3 = %d, 余数 = %d\n", quotient, remainder)

	// 可变参数函数
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("1 到 5 的和: %d\n", total)

	// 函数作为参数
	multiplyResult := applyOperation(func(a, b int) int {
		return a * b
	}, 4, 5)
	fmt.Printf("4 * 5 = %d\n", multiplyResult)
}
