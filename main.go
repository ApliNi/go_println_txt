package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	fileName := strings.Replace(strings.TrimSuffix(os.Args[0], ".exe"), "-", ".", -1)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("文件 `%s` 不存在\n", fileName)
		return
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("读取文件时发生错误: `%s`\n", err)
		return
	}

	fmt.Println(string(content))
}
