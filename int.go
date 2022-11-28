package delta

import "github.com/golang-infrastructure/go-gtypes"

// ZipIntegerSlice delta压缩整数，只适合正整数一般
func ZipIntegerSlice[T gtypes.Integer](intSlice []T) []T {
	result := make([]T, len(intSlice))
	for index, x := range intSlice {
		if index == 0 {
			result[index] = x
			continue
		}
		delta := intSlice[index-1] - intSlice[index]
		result[index] = delta
	}
	return result
}

// UnzipIntegerSlice 解压缩delta
func UnzipIntegerSlice[T gtypes.Integer](intSlice []T) []T {
	result := make([]T, len(intSlice))
	for index, x := range intSlice {
		if index == 0 {
			result[index] = x
			continue
		}
		delta := result[index-1] - intSlice[index]
		result[index] = delta
	}
	return result
}
