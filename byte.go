package delta

// ToByteSliceDelta 不单独实现了，直接认为是int的一种特殊情况，为了调用方的代码可读性这里就给个别名
var ToByteSliceDelta = ToIntegerSliceDelta[byte]
var FromByteSliceDelta = FromIntegerSliceDelta[byte]

//// ToByteSliceDelta 字节数组
//func ToByteSliceDelta(bytes []byte, compareToType CompareToType) []byte {
//	deltaSlice := make([]byte, 0)
//	for index, byteValue := range bytes {
//		if index == 0 {
//			deltaSlice[index] = byteValue
//			continue
//		}
//		switch compareToType {
//		case CompareToFirst:
//			deltaSlice[index] = byteValue - bytes[0]
//		case CompareToLast:
//			deltaSlice[index] = byteValue - bytes[index-1]
//		default:
//			panic(fmt.Errorf("not support compare type: %#v", compareToType))
//		}
//	}
//	return deltaSlice
//}

//func FromByteSliceDelta(deltaSlice []byte, compareToType CompareToType) []byte {
//	bytes := make([]byte, 0)
//	for index, delta := range deltaSlice {
//		if index == 0 {
//			bytes[index] = delta
//			continue
//		}
//		switch compareToType {
//		case CompareToFirst:
//			bytes[index] = bytes[index-1] + delta
//		case CompareToLast:
//			bytes[index] = bytes[index-1] + delta
//		default:
//			panic(fmt.Errorf("not support compare type: %#v", compareToType))
//		}
//	}
//	return bytes
//}
