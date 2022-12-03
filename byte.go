package delta

// ToByteSliceDelta 不单独实现了，直接认为是int的一种特殊情况，为了调用方的代码可读性这里就给个别名
var ToByteSliceDelta = ToIntegerSliceDelta[byte]
var FromByteSliceDelta = FromIntegerSliceDelta[byte]
