package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func main() {
	// todo: パスをハードコードしているので、引数で指定できるようにする
	// todo: 変換後の画像の拡張を指定できるようにする
	imageConv()
}

// 画像変換処理
func imageConv() {
	// パスを指定して画像の読み込み
	f, err := os.Open("./test/practice01.jpg")
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

	fso, err := os.Create("practice_test.png")
	if err != nil {
		fmt.Println("create error: ", err)
		return
	}
	defer fso.Close()

	png.Encode(fso, img)
}

//func main() {
//
//	// 引数をパースして引数以下のファイルを読み込めるようにする
//	flag.Parse()
//	args := flag.Args()
//	dirName := genDirName(args[0])
//
//	readFiles(dirName)
//}
//
//// 引数で入力画像を受け取る
//func genDirName(s string) string {
//	dirName := "./" + s
//	return dirName
//}
//
//// ディレクトリ内のファイルを読み取る処理
//func readFiles(dir string) {
//	files, _ := ioutil.ReadDir(dir)
//
//	for _, file := range files {
//		// 元画像の読み込み
//		f, _ := os.Open("./test/practice01.jpg")
//
//
//		img, _, _ := image.Decode(file)
//		defer file.Close()
//
//		bound := img.Bounds()
//		fmt.Printf("Go-Logo_LightBlue.jpg\tX : %d px, Y : %d px\n", bound.Dx(), bound.Dy())
//	}
//}
