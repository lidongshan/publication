package main

import (
	"fmt"
	"testing"
)

func byteSliceToString() {
	sl := []byte{
		0xE4, 0xB8, 0xAD,
		0xE5, 0x9B, 0xBD,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
		0xEF, 0xBC, 0x8C,
		0xE5, 0x8C, 0x97,
		0xE4, 0xBA, 0xAC,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
	}

	_ = string(sl)
}

func stringToByteSlice() {
	s := "中国欢迎您，北京换欢您"
	_ = []byte(s)
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, byteSliceToString))
	fmt.Println(testing.AllocsPerRun(1, stringToByteSlice))
}
