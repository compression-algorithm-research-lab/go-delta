package delta

import (
	"fmt"
	"github.com/golang-infrastructure/go-gtypes"
)

// ToIntegerSliceDelta delta压缩整数，只适合正整数一般
func ToIntegerSliceDelta[T gtypes.Integer](intSlice []T, compareToType CompareToType) []T {
	deltaSlice := make([]T, len(intSlice))
	for index := range intSlice {
		if index == 0 {
			deltaSlice[index] = intSlice[index]
			continue
		}
		switch compareToType {
		case CompareToFirst:
			deltaSlice[index] = intSlice[index] - intSlice[0]
		case CompareToLast:
			deltaSlice[index] = intSlice[index] - intSlice[index-1]
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return deltaSlice
}

// FromIntegerSliceDelta 解压缩delta
func FromIntegerSliceDelta[T gtypes.Integer](deltaSlice []T, compareToType CompareToType) []T {
	intSlice := make([]T, len(deltaSlice))
	for index, delta := range deltaSlice {
		if index == 0 {
			intSlice[index] = delta
			continue
		}
		switch compareToType {
		case CompareToFirst:
			intSlice[index] = deltaSlice[0] + delta
		case CompareToLast:
			intSlice[index] = intSlice[index-1] + delta
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return intSlice
}
