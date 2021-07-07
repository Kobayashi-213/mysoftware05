package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
)

//入力からスライスを作成
func makeSlice() {
	var stdin = bytes.NewBuffer([]byte("1 1 +"))
	scanner := bufio.NewScanner(stdin)
	scanner.Split(bufio.ScanWords)
	slices := []string{}
	for scanner.Scan() {
		slice := scanner.Text()
		slices = append(slices, slice)
	}
	fmt.Println(slices)
}

//答えを表示
func displayAnswer(i int) int {
	fmt.Println(i)
}

//文字列からx,yを数値として抜き出す
func getXY(s string) (x int, y int) {
	//input  1 3 +
	//return 1 3

	x = strconv.Itoa(s[0])
	y = strconv.Itoa(s[1])

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

}
