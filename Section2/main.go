package main

import "fmt"

// 変数
func main() {
	// 明示的な定義
	var variable string = "string"
	fmt.Println(variable)

	var t, f = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "golang"
	)

	fmt.Println(i2, s2)

	// 変数の型だけ定義して値を格納しない場合はその型の初期値が入る
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	// 暗黙的な定義
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)
}
