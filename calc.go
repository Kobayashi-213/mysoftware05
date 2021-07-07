package main

import (
	"fmt"

	"strconv"
)

//入力からスライスを作成
func makeSlice() []string {
	//var x int
	slice := make([]string, 5)
	fmt.Scanln(&slice[0], &slice[1], &slice[2], &slice[3], &slice[4])

	return slice
}

//答えを表示
func displayAnswer(i int) int {
	//標準出力
	fmt.Println(i)
	return i
}

//文字列からx,yを数値として抜き出す
func getXY(s []string) (int, int, string) {
	//input  1 3 +
	//return 1 3
	A := s[0]
	B := s[1]
	z := s[2]
	//fmt.Println(A)
	//fmt.Println(B)
	x, _ := strconv.Atoi(A)
	y, _ := strconv.Atoi(B)

	return x, y, z
}

//2つの数値と四則演算子から計算結果を返す
func calculat(x int, y int, cal string) int {
	answer := 0
	if cal == "+" {
		answer = x + y
	}
	if cal == "-" {
		answer = x - y
	}
	if cal == "*" {
		answer = x * y
	}
	if cal == "/" {

	}
	return answer
}

func main() {
	slice := makeSlice()
	fmt.Println(slice)
}
