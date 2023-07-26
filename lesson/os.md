package main

import (
	"fmt"
	"os"
)

//os

func main() {
	//Exit
	// os.Exit(1)
	// fmt.Println("start")

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}

```
package main

import (

	"log"
	"os"
)

//os

func main() {
	//ファイル操作
	//Open
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
}

```
```
package main

import (
	"os"
)

func main() {
	// ファイルを作成する
	f, err := os.Create("foo.txt")
	if err != nil {
		// エラーが発生した場合は、エラーを表示してプログラムを終了する
		panic(err)
	}

	// ファイルに"Hello"と書き込む
	// Write関数は、引数に指定したバイトスライスをファイルに書き込む関数です。
	// []byte("Hello\n")は、文字列"Hello"をバイトスライスに変換しています。
	f.Write([]byte("Hello\n"))

	// ファイルの指定位置に"Golang"と書き込む
	// WriteAt関数は、第1引数に指定したバイトスライスを、第2引数で指定したオフセット位置に書き込む関数です。
	// ここでは、ファイルの6番目の位置に"Golang"を書き込んでいます（先頭を0と数えます）。
	f.WriteAt([]byte("Golang"), 6) //6文字目に書き込む

	// ファイルの末尾に移動して"Yaah"と書き込む
	// Seek関数は、ファイル内の読み書き位置を変更する関数です。
	// 第1引数に移動する位置を、第2引数に位置の基準を指定します。ここではファイルの末尾を基準にしています。
	// ファイルの末尾に移動した後、WriteString関数を使用して文字列"Yaah"をファイルに追記しています。
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")

	// ファイルを閉じる
	// ファイル操作が終わったら、必ずファイルを閉じることが重要です。
	// defer文を使用することで、関数の最後でファイルを閉じるようにしています。
	defer f.Close()
}


//foo.txt
	// Hello
	// GolangYaah

```
