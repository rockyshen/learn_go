package main

import "fmt"

// 自定义错误
type CustomError struct {
	Message string
	Code    int
}

// 函数，实现error包下的Error()方法
func (e *CustomError) Error() string {
	return fmt.Sprintf("错误 %d: %s", e.Code, e.Message)
}

// 显式错误处理
func divid(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &CustomError{
			Message: "除数不能为零",
			Code:    1001,
		}
	}
	return a / b, nil
}

func recoverFromPanic() {
	// 使用recover 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
		}
	}()
	// 模拟会发生异常的汗水
	dangerousOperation()
}

func dangerousOperation() {
	// 模拟一个可能导致 panic 的操作
	panic("发生了严重错误")
}

func main() {
	res, err := divid(10, 2)
	if err != nil {
		// 如果是CustomError类型的错误，走这里
		if customErr, ok := err.(*CustomError); ok {
			fmt.Printf("自定义错误: %s (错误码: %d)\n", customErr.Message, customErr.Code)
		} else { // 其他普通报错，走这里
			fmt.Println("普通错误：", err)
		}
	} else {
		fmt.Println("结果：", res)
	}

	// 调用可能导致panic的函数，但是通过cover捕获
	recoverFromPanic()
}
