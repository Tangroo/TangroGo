package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// var name, age = "tangro", 18

// var (
// 	name = "tang"
// 	age  = 20
// )

func foo() (int, string) {
	return 10, "regan"
}

func main() {
	num := rand.Int63()
	fmt.Printf("rand.Int63:%d\n", num)
	nu := strconv.Itoa(int(num))
	le := len(nu)
	fmt.Println(le)
	now := time.Now().UnixNano()
	fmt.Printf("time.Now().UnixNano:%d\n", now)
	newsource := rand.New(rand.NewSource(now + num))
	fmt.Printf("rand.New:%d\n", newsource)
	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y)
}
