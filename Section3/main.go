package main

import "fmt"

func main() {
	// intの後の数字はビット数を表す
	var i int8 = 100
	fmt.Println(i)

	// 現在の型を調べる
	fmt.Printf("%T\n", i)

	// 型変換
	fmt.Printf("%T\n", int32(i))

	// float32はあまりつかわない(デフォルトfloat64)

	// byte型
	// 文字列に変換
	byteA := []byte{72, 73}
	fmt.Println(string(byteA))

	// 文字列からbyteに変換
	c := []byte("HI")
	fmt.Println(c)

	// 配列型
	// 要素数を変更することはできない
	var arr1 [3]int
	fmt.Println(arr1) // 結果:000
	// [3]int　型が左の状態になるので要素数を変更することができない

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 要素の数を自動でカウントするように設定
	arr3 := [...]string{"A"}
	fmt.Println(arr3)

	// 操作
	fmt.Println(arr1[0])

	// interface
	var x interface{}
	fmt.Println(x)

	// 型変換

}
