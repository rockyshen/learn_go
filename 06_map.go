package main

import "fmt"

func main() {
	// 创建 map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("初始 map:", ages)

	// 添加和修改元素
	ages["Charlie"] = 35
	ages["Bob"] = 31
	fmt.Println("修改后的 map:", ages)

	// 检查键是否存在
	age, exists := ages["David"]
	if exists {
		fmt.Println("David 的年龄:", age)
	} else {
		fmt.Println("David 不在 map 中")
	}

	// 删除元素
	delete(ages, "Alice")
	fmt.Println("删除 Alice 后的 map:", ages)

	// 遍历 map
	for name, age := range ages {
		fmt.Printf("%s 的年龄是 %d\n", name, age)
	}
}
