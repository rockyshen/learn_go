package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() string {
	return fmt.Sprintf("我是 %s, 今年 %d 岁", p.Name, p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	alice := Person{Name: "Alice", Age: 25}
	result := alice.SayHello()
	fmt.Println(result)

	alice.HaveBirthday()
	result2 := alice.SayHello()
	fmt.Println(result2)
}
