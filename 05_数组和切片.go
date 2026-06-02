package main

import "fmt"

func main() {
	// 数组：固定长度
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	fmt.Println("数组:", numbers)

	// 数组也可以一开始指定默认值
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("数组，声明时提供默认值：", arr)

	// 切片：动态长度
	fruits := [5]string{"apple", "banana", "orange"}
	fmt.Println("切片:", fruits)

	// 切片操作
	fmt.Println("切片第一个元素:", fruits[0])
	fmt.Println("切片长度:", len(fruits))

	fmt.Println("切片容量：", cap(fruits))

	// 切片追加元素
	//fruits = append(fruits, "grape")
	//fmt.Println("追加后的切片:", fruits)

	// 切片切割
	subFruits := fruits[1:3]
	fmt.Println("子切片:", subFruits)

	// 使用 make 创建切片
	numbers2 := make([]int, 5, 10)
	fmt.Printf("使用 make 创建的切片: %v, 长度: %d, 容量: %d\n", numbers2, len(numbers2), cap(numbers2))
}
