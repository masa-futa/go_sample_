package main

import "fmt"

// 定数
const n int = 1

// 複数宣言可能
// 右辺が省略可能
const (
	m = 1
	s
	ss
	sss
)

// 変数
var i int = 1

func main() {
	// if - swiftと変わらない感じ
	// if i == 1 {

	// } else if i == 2 {

	// } else {

	// }

	for i := 0; i <= 100; i = i + 1 {
		if i%2 == 1 {
			fmt.Print(i)
			fmt.Print("奇数")
		} else {
			fmt.Print(i)
			fmt.Print("偶数")
		}
	}

	// 	// 特徴
	// 	// 代入分が記述できる（ブロック内でのみ使用可能）
	// 	if i := 2; i == 2 {

	// 	}

	// 	switch i {
	// 	case 1:
	// 		fmt.Println("1")
	// 	default:
	// 		fmt.Println("default")
	// 	}
	// 	// if分のように式が使用可能
	// 	switch {
	// 	case i == 1:

	// 	}
	// 	// Goには繰り返し構文がforしかない
	// 	for i := 0; i <= 100; i = i + 1 {

	// 	}
	// 	// breakにてループ抜け出し
	// 	for i <= 100 {
	// 		break
	// 	}
	// 	// ラベル指定にて、抜け出す範囲を指定できる
	// LOOP:
	// 	for {
	// 		switch {
	// 		case i%2 == 1:
	// 			break LOOP
	// 		}
	// 	}

	// 	// for i, v := range []int{1, 2, 3} {

	// 	// }

	// 	fmt.Println("Hello")
	// 	fmt.Println(ss)
}
