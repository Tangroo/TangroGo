package main

import (
	"code/video18practice/calc"
	"fmt"
)

func main() {
	// 1. 打印菜单和提示信息，接收用户输入
	var (
		x      int
		y      int
		method string
	)
	fmt.Println("please enter first number:")
	fmt.Scanf("%d\n", &x)
	fmt.Println("Please enter the calculation method(+,-,*,/):")
	fmt.Scanf("%s\n", &method)
	fmt.Println("please enter first number:")
	fmt.Scanf("%d\n", &y)
	// 2. 接收用户输入，并判断用户输入是否合法，若合法则匹配对应的算法规则
	// 返回结果
	switch method {
	case "+":
		result := calc.Add(x, y)
		fmt.Printf("%d %s %d = %d", x, method, y, result)
	case "-":
		result := calc.Sub(x, y)
		fmt.Printf("%d %s %d = %d", x, method, y, result)
	case "*":
		result := calc.Mul(x, y)
		fmt.Printf("%d %s %d = %d", x, method, y, result)
	case "/":
		result := calc.Div(x, y)
		fmt.Printf("%d %s %d = %d", x, method, y, result)
	case "\\":
		result := calc.Div(x, y)
		fmt.Printf("%d %s %d = %d", x, method, y, result)
	}
	fmt.Println("\nplease enter any key to exit.")
	fmt.Scanf("haha")
}
