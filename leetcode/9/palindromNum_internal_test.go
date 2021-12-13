package pal

import "testing"

const palInt = 98765432123456789
const notPalInt = 123456789123456789

func BenchmarkStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindromeStr(palInt)
	}
}

func BenchmarkNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindromeNum(palInt)
	}
}

func Test_isPalindromeNum(t *testing.T) {
	if !isPalindromeNum(palInt) {
		t.Error("!")
	}
	if isPalindromeNum(notPalInt) {
		t.Error("!")
	}
}

func Test_isPalindromeStr(t *testing.T) {
	if !isPalindromeStr(palInt) {
		t.Error("!")
	}
	if isPalindromeStr(notPalInt) {
		t.Error("!")
	}
}
