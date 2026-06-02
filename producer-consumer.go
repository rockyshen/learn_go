package main

import "fmt"

// 创建一个 channel （队列），容量是10
var stream = make(chan int, 10)
const n = 4

// 生产者
func produce() {
  for i := 0; ; i++ {
    fmt.Println("produce", i)   // 模拟生产一笔数据
    stream <- i                 // 往队列中生产一笔数据
  }
}

// 消费者
func consume() {
  for {
    x := <- stream                   // 从队列中消费一笔数据
    fmt.Println("consume", x)        // 模拟消费一笔数据
  }
}

func main() {
  for i := 0; i < n; i++ {
    go produce()
  }
  consume()
}
