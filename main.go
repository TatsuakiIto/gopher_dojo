package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main()  {

	// 引数をパースして引数以下のファイルを読み込めるようにする
	flag.Parse()
	args := flag.Args()
	dirName := genDirName(args[0])

	readFiles(dirName)

}

// 引数を一つ渡して、ディレクトリ名を返す
func genDirName(s string) string{
	dirName := "./" + s
	return dirName
}

// ディレクトリ内のファイルを読み取る処理
func readFiles(dir string) {
	files, _ := ioutil.ReadDir(dir)

	// 本当はエラーハンドリングする

	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}

	fmt.Println(names)
}

// TODO: fileを渡してjpegの拡張子をpngに変える処理を作る
func changeExtension() {

}
