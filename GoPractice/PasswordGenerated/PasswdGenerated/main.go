package main

import (
	"fmt"
)

// func passGenerated(){
// 	rand.New()
// }

func main() {
	for {
		var x int
		var y int
		// 1. 询问用户需要生成那种类型的密码：
		// 1）纯数字密码
		// 2）小写字母+数字密码
		// 3) 大小写字母+数字密码
		// 4）大小写字母+数字+特殊字符密码
		// 5) 退出
		fmt.Println(`选项：
1 纯数字密码
2 小写字母 + 数字密码
3 大小写字母+数字密码
4 大小写字母+数字+特殊字符密码
5 退出`)
		fmt.Printf("请输入数字选择需要生成的密码类型：")
		fmt.Scanf("%d", &x)
		fmt.Println(x)
		// 询问用户需要输出的密码位数
		fmt.Printf("请输入密码位数：")
		fmt.Scanf("%d", &y)
	}
}
