package main

import (
	"fmt"
	"os"
)

func getRand(baseStr string,)

func main() {

	var x int
	var y int
	var number string
	number = "0123456789"
	for {
		fmt.Println(`选项：
1 纯数字密码
2 小写字母 + 数字密码
3 大小写字母 + 数字密码
4 大小写字母 + 数字 + 特殊字符密码
5 退出`)
		fmt.Printf("请输入数字选择需要生成的密码类型：")
		fmt.Scanln(&x)
		// 询问用户需要输出的密码位数
		fmt.Printf("请输入密码位数：")
		fmt.Scanln(&y)
		// fmt.Printf("%T", y)
		switch x {
		case 1:
			fmt.Println(x)
		case 2:
			fmt.Println(x)
		case 3:
			fmt.Println(x)
		case 4:
			fmt.Println("Sorry,Not supported")
		case 5:
			fmt.Println("See you!")
			os.Exit(0)
		}
	}
}
