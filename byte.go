package delta

// ByteSliceDelta 字节数组
func ByteSliceDelta(bytes []byte) []byte {
	result := make([]byte, 0)
	for index, value := range bytes {
		if index == 0 {
			result[index] = value
		} else {
			result[index] = bytes[index-1] - value
		}
	}
	return result
}

func ByteSliceFromDelta(bytes []byte) []byte {
	result := make([]byte, 0)
	for index, value := range bytes {
		if index == 0 {
			result[index] = value
		} else {
			result[index] = result[index-1] + value
		}
	}
	return result
}
