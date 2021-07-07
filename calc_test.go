package main

import "testing"

func Test01(t *testing.T) {
	expected := 42
	result := displayAnswer(42)
	if result != expected {
		t.Errorf("test01 Error")
	}
}

func Test02(t *testing.T){
	expected := [1 1 +]
	result := makeSlice()
	if result != expected {
		t.Errorf("test02 Error")
	}
}

func Test04(t *testing.T) {
	expected := 8
	result := calculat(6, 2, "+")
	if result != expected {
		t.Errorf("test04 Error")
	}
}

//文字列からx,yを数値として抜き出す 川原
func Test03(t *testing.T) {
	expX := 1
	expY := 3
	resultX, resultY := getXY("13+")
	if (expX != resultX) || (expY != resultY) {
		t.Errorf("test03 Error")
	}
}
