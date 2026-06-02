package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// ========== 基础：go 关键字启动 Goroutine ==========

func sayHello(name string, delay time.Duration) {
	fmt.Printf("[%s] %s: 开始...\n", time.Now().Format("15:04:05"), name)
	time.Sleep(delay) // 阻塞时，Go 调度器自动切换 Goroutine
	fmt.Printf("[%s] %s: 你好！（等了 %v）\n", time.Now().Format("15:04:05"), name, delay)
}

// ========== 顺序 vs 并发对比 ==========

func sequentialRun() {
	fmt.Println("\n=== 顺序执行 ===")
	start := time.Now()

	sayHello("任务A", 1*time.Second)
	sayHello("任务B", 2*time.Second)
	sayHello("任务C", 3*time.Second)

	fmt.Printf("总耗时: %.1f 秒\n", time.Since(start).Seconds())
}

func concurrentRun() {
	fmt.Println("\n=== 并发执行（go + sync.WaitGroup）===")
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(3) // 计数器 = 3

	go func() {
		defer wg.Done() // 完成时减1
		sayHello("任务A", 1*time.Second)
	}()

	go func() {
		defer wg.Done()
		sayHello("任务B", 2*time.Second)
	}()

	go func() {
		defer wg.Done()
		sayHello("任务C", 3*time.Second)
	}()

	wg.Wait() // 阻塞等待计数器归零
	fmt.Printf("总耗时: %.1f 秒\n", time.Since(start).Seconds())
}

// ========== 实战：并发网络请求（模拟） ==========

func fetchData(url string, delay time.Duration, result chan<- string) {
	fmt.Printf("  请求 %s ...\n", url)
	time.Sleep(delay) // 模拟网络延迟
	result <- fmt.Sprintf("数据-%s", url) // 发送结果到 channel
}

func fetchMultiple() {
	fmt.Println("\n=== 模拟并发网络请求 ===")
	start := time.Now()

	urls := []struct {
		url   string
		delay time.Duration
	}{
		{"api/users", 1 * time.Second},
		{"api/orders", 2 * time.Second},
		{"api/products", 1 * time.Second},
	}

	// 创建 buffered channel 收集结果
	results := make(chan string, len(urls))

	for _, u := range urls {
		go fetchData(u.url, u.delay, results) // 启动 Goroutine
	}

	// 收集结果
	var data []string
	for i := 0; i < len(urls); i++ {
		data = append(data, <-results) // 从 channel 接收（阻塞等待）
	}

	fmt.Printf("全部完成: %v\n", data)
	fmt.Printf("总耗时: %.1f 秒（不是 4 秒！）\n", time.Since(start).Seconds())
}

// ========== 进阶：Goroutine 间通信（Channel） ==========

func producer(ch chan<- int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i // 发送数据（channel 满时自动阻塞 Goroutine）
		fmt.Printf("  [生产] 产品-%d\n", i)
	}
	close(ch) // 生产完毕，关闭 channel
}

func consumer(ch <-chan int, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch { // range 自动检测 channel 关闭
		time.Sleep(800 * time.Millisecond)
		fmt.Printf("  [消费-%s] 处理 产品-%d\n", name, item)
	}
}

func producerConsumerDemo() {
	fmt.Println("\n=== Goroutine 间通信：Channel ===")
	ch := make(chan int, 2) // 有缓冲 channel（容量2）

	var wg sync.WaitGroup
	wg.Add(3)

	go producer(ch, 5, &wg)
	go consumer(ch, "A", &wg)
	go consumer(ch, "B", &wg)

	wg.Wait()
}

// ========== 进阶：Select 多路复用 ==========

func selectDemo() {
	fmt.Println("\n=== Select 多路复用 ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自 ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自 ch2"
	}()

	// 等待多个 channel，谁先就绪就处理谁
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("收到: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("收到: %s\n", msg2)
		}
	}
}

// ========== 避坑：同步阻塞 vs 让出调度 ==========

func cpuIntensive(id int) {
	// CPU 密集型：Goroutine 不会自动让出，需手动 yield
	sum := 0
	for i := 0; i < 1e8; i++ {
		sum += i
		if i%1e7 == 0 {
			runtime.Gosched() // 手动让出，让其他 Goroutine 运行
		}
	}
	fmt.Printf("  CPU 任务 %d 完成，sum=%d\n", id, sum)
}

func trapDemo() {
	fmt.Println("\n=== 陷阱：CPU 密集型不会自动切换 ===")
	// Go 1.14+ 有协作式抢占，但 CPU 密集仍可能饿死其他 Goroutine
	// 解决方案：用 runtime.Gosched() 或拆分为小任务
}

// ========== 主入口 ==========

func main() {
	fmt.Println("=" + "=================================================")
	fmt.Println("Go Goroutine 并发示例")
	fmt.Println("=" + "=================================================")

	// 1. 顺序执行
	sequentialRun()

	// 2. WaitGroup 并发
	concurrentRun()

	// 3. 模拟并发 IO
	fetchMultiple()

	// 4. Channel 通信
	producerConsumerDemo()

	// 5. Select 多路复用
	selectDemo()

	// 6. 避坑提示
	trapDemo()

	fmt.Println("\n" + "==================================================")
	fmt.Println("全部演示完成！")
	fmt.Println("==================================================")
}