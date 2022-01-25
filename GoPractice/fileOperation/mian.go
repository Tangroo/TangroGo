package main

import (
	"fmt"
	"io"
	"os"
)

// 1. open source.txt file
// 2. create dest.txt file
// 3. read

func CopyFile(srcFileName, destFileName string) (written int64, err error) {
	srcfile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("Open %s faild!\nError:%v\n", srcFileName, err)
		return
	}
	defer srcfile.Close()
	destfile, err := os.OpenFile(destFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err!=nil{
		fmt.Printf("Open %s faild!\nError:%v\n", destFileName, err)
		return
	}
	defer destfile.Close()
	return io.Copy(srcfile,destfile)
}

func main() {

}
