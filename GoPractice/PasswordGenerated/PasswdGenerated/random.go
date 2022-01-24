// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	s1 := "3"
// 	s2 := os.Args[0]
// 	fmt.Printf("%s\n%T\n", s2, s2)
// 	i, err := strconv.Atoi(s1)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("%d,%T", i, i)
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	length := 16
	if len(os.Args) > 1 {
		arg := os.Args[1]
		i, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Println("参数转换失败")
			return
		}
		if i < 4 || i > 30 {
			fmt.Println("密码长度介于4~30之间")
			return
		}
		length = int(i)
	}
	baseStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseSymbol := "!@#$%^&*()[]{}+-*/_=."
	fmt.Println("-----简单密码-----")
	for i := 0; i < 5; i++ {
		fmt.Println(getRandStr(baseStr, length))
	}
	fmt.Println("-----复杂密码-----")
	for i := 0; i < 5; i++ {
		fmt.Println(getRandStr(baseStr+baseSymbol, length))
	}
}

func getRandStr(baseStr string, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}
