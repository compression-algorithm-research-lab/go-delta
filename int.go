package delta

import (
	"fmt"
	"github.com/golang-infrastructure/go-gtypes"
)

// ZipIntegerSlice delta压缩整数，只适合正整数一般
func ZipIntegerSlice[T gtypes.Integer](intSlice []T, compareToType CompareToType) []T {
	deltaSlice := make([]T, len(intSlice))
	for index, x := range intSlice {
		if index == 0 {
			deltaSlice[index] = x
			continue
		}
		var delta T
		switch compareToType {
		case CompareToFirst:
			delta = intSlice[index] - intSlice[0]
		case CompareToLast:
			delta = intSlice[index] - intSlice[index-1]
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
		deltaSlice[index] = delta
	}
	return deltaSlice
}

// UnzipIntegerSlice 解压缩delta
func UnzipIntegerSlice[T gtypes.Integer](deltaSlice []T, compareToType CompareToType) []T {
	intSlice := make([]T, len(deltaSlice))
	for index, delta := range deltaSlice {
		if index == 0 {
			intSlice[index] = delta
			continue
		}
		var data T
		switch compareToType {
		case CompareToFirst:
			data = deltaSlice[0] + delta
		case CompareToLast:
			data = intSlice[index-1] + delta
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
		intSlice[index] = data
	}
	return intSlice
}
