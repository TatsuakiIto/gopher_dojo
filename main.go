package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("===== Process is start =====")
	// 引数を受け取る
	flag.Parse()
	args := flag.Args()

	// todo: ディレクトリ名を指定してファイルのパスのsliceを作る
	dirName := genDirName(args[0])
	readFiles(dirName)

	// todo: 変換後の画像の拡張を指定できるようにする

	fmt.Println("===== Process is end =====")
}

// 画像変換処理
func imageConv(filePath string) {
	// パスを指定して画像の読み込み
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open error:", err)
		return
	}
	defer f.Close()

	// 元画像のdecode
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	// 指定したファイル名で新規作成
	fso, err := os.Create("practice_test.png")
	if err != nil {
		fmt.Println("create error: ", err)
		return
	}
	defer fso.Close()

	png.Encode(fso, img)
}

// 引数をディレクトリー名に変換
func genDirName(s string) string{
	return "./" + s + "/"
}

// ディレクトリー内のファイルを読み取る処理
func readFiles(dirName string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("read err: ", err)
		return
	}
	for _, file := range files{
		filepath := "./test/" + file.Name()
		imageConv(filepath)
	}
}
