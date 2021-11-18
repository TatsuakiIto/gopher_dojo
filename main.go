package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main()  {
	flag.Parse()
	fmt.Println(flag.Args())
}

// ディレクトリ内のファイルを読み取る処理
func readFiles()  {
	files, _ := ioutil.ReadDir("./test")

	// 本当はエラーハンドリングする

	var names []string
	for _, file := range files {
		newName := strings.Replace(file.Name(), ".txt", ".csv", 1)
		names = append(names, newName)
	}

	fmt.Println(names)
}
