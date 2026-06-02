package main

import "fmt"

func main() {
	// if-else 语句
	x := 10
	if x > 5 {
		fmt.Println("x 大于 5")
	} else if x < 5 {
		fmt.Println("x 小于 5")
	} else {
		fmt.Println("x 等于 5")
	}

	// switch 语句
	switch x {
	case 10:
		fmt.Println("x 是 10")
		fallthrough
	case 5:
		fmt.Println("x 是 5") //继续 执行下一个 case 语句
	default:
		fmt.Println("x 是其他值")
	}

	// for 循环
	// 传统循环
	for i := 0; i < 5; i++ {
		fmt.Printf("传统循环: %d\n", i)
	}

	// while 风格的循环
	j := 0
	for j < 3 {
		fmt.Printf("while 风格循环: %d\n", j)
		j++
	}

	// 无限循环
	k := 0
	for {
		fmt.Printf("无限循环计数: %d\n", k)
		k++
		if k >= 3 {
			break
		}
	}
}
