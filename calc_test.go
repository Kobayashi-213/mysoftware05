package main

import "testing"

func Test01(t *testing.T) {
	expected := 42
	result := displayAnswer(42)
	if result != expected {
		t.Errorf("test01 Error")
	}
}

/*func Test02(t *testing.T){
	expected := [1 1 +]
	result := makeSlice()
	if result != expected {
		t.Errorf("test02 Error")
	}
}*/

//スライスからx,yを数値として抜き出す 川原
func Test03(t *testing.T) {
	expX := 1
	expY := 3
	str := []string{"1", "3", "+"}
	resultX, resultY := getXY(str)
	if (expX != resultX) || (expY != resultY) {
		t.Errorf("test03 Error")
	}
}

//2つの数値と"+"から計算結果を返す
func Test04(t *testing.T) {
	expected := 8
	result := calculat(6, 2, "+")
	if result != expected {
		t.Errorf("test04 Error")
	}
}

//2つの数値と"-"から計算結果を返す
func Test05(t *testing.T) {
	expected := 4
	result := calculat(6, 2, "-")
	if result != expected {
		t.Errorf("test05 Error")
	}
}

//2つの数値と"*"から計算結果を返す
func Test06(t *testing.T) {
	expected := 12
	result := calculat(6, 2, "*")
	if result != expected {
		t.Errorf("test06 Error")
	}
}
