package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func main() {
	// 创建 JavaScript 解释器
	vm := otto.New()

	// 执行 JavaScript 代码
	_, err := vm.Run(`
		function greet() {
			return "Hello from JavaScript!";
		}
		greet();
	`)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取 JavaScript 函数的返回值
	result, err := vm.Call("greet", nil)
	if err != nil {
		fmt.Println("Error calling function:", err)
		return
	}

	// 输出结果
	fmt.Println(result.String()) // 输出：Hello from JavaScript!
}
