package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//入力からスライスを作成
func makeSlice() {
	slices := []string{}
	for {
		//標準入力から受け取る
		slice := fmt.Scan(&x)
		slices = append(slices, slice)
		if slice == "\n" {
			break
		}
	}
}

//答えを表示
func displayAnswer(i int) int {
	//標準出力
	fmt.Println(i)
	return 42
}

//文字列からx,yを数値として抜き出す
func getXY(s string) (int, int) {
	//input  "13+"
	//return 1,3
	A := s[0:1]
	B := s[1:2]
	//fmt.Println(A)
	//fmt.Println(B)
	x, _ := strconv.Atoi(A)
	y, _ := strconv.Atoi(B)

	return x, y
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
	makeSlice()
}
