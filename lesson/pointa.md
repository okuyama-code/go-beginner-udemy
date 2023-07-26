ポインタは、変数のメモリアドレスを保持するデータ型であり、その変数にアクセスするための手段として機能します。

```
package main

import "fmt"

//ポインタ

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {  //sはスライスを渡している
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n) //100

	fmt.Println(&n) //0xc0000a6058  メモリ上のアドレス

	Double(n)

	fmt.Println(n) //100 //200にこれではならない

	var p *int = &n //ポインタ型の宣言(*データ型) アドレスを渡す　＆はアドレス演算子と呼ばれている
	fmt.Println(p) //0xc000018098
	fmt.Println(*p) //100  アドレスがさす実体を表示

	// *p =300 //実体のｐ
	// fmt.Println(n) //300に置き換わる　ｐとｎは同じアドレスを指しているため 同じアドレスのｎも値が変わる仕組み
	// n = 200
	// fmt.Println(*p) //200 アドレスを共有できる

	Doublev2(&n) //アドレスを渡す ポインタ型を使うことでアドレスで渡して参照渡し(コピーしないで渡している)
	fmt.Println(n) //200

	Doublev2(p) //var p *int = &n
	fmt.Println(*p) //実体のpを渡す　　400

	var sl []int = []int{1, 2, 3} //参照型のデータ構造の場合はもともと参照渡しの機能を持っている。だからポインタ型を使わなくても値を書き換えることができる

	Doublev3(sl)
	fmt.Println(sl) //[2 4 6]
}

 ```
