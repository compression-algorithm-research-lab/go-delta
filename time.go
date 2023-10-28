package delta

import (
	"fmt"
	"time"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// DefaultCompareToType 如果未指定delta的时候的比较对象的话，则默认将当前时间与序列相邻的上一个时间做delta比较
const DefaultCompareToType = CompareToPrevious

// ToTimeSliceDelta 对时间序列在毫秒单位做delta
func ToTimeSliceDelta(timeSlice []time.Time, compareToType ...CompareToType) []uint64 {

	if len(compareToType) == 0 {
		compareToType = append(compareToType, DefaultCompareToType)
	}

	deltaSlice := make([]uint64, len(timeSlice))
	for index, t := range timeSlice {

		// 第一个保持原样
		if index == 0 {
			deltaSlice[index] = uint64(t.UnixMilli())
			continue
		}

		//从第二个开始就要delta了
		switch compareToType[0] {
		case CompareToFirst:
			deltaSlice[index] = TimeDelta(t, timeSlice[0])
		case CompareToPrevious:
			deltaSlice[index] = TimeDelta(t, timeSlice[index-1])
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return deltaSlice
}

// FromTimeSliceDelta 从Time的delta恢复时间序列
func FromTimeSliceDelta(deltaSlice []uint64, compareToType ...CompareToType) []time.Time {

	if len(compareToType) == 0 {
		compareToType = append(compareToType, DefaultCompareToType)
	}

	timeSlice := make([]time.Time, len(deltaSlice))
	for index, delta := range deltaSlice {

		if index == 0 {
			timeSlice[index] = time.UnixMilli(int64(delta))
			continue
		}

		switch compareToType[0] {
		case CompareToFirst:
			timeSlice[index] = time.UnixMilli(int64(deltaSlice[0] + delta))
		case CompareToPrevious:
			timeSlice[index] = timeSlice[index-1].Add(time.Millisecond * time.Duration(delta))
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}

	}

	return timeSlice
}

// TimeDelta 求两个时间的毫秒差
func TimeDelta(t1 time.Time, t2 time.Time) uint64 {
	return uint64(t1.UnixMilli() - t2.UnixMilli())
}

// ------------------------------------------------ ---------------------------------------------------------------------

// ToTimestampSliceDelta 对时间戳做delta，这个时间戳可以根据自己需要，可以是10位的描述，也可以是13位的毫秒数
func ToTimestampSliceDelta(timestampSlice []uint64, compareToType ...CompareToType) []uint64 {

	if len(compareToType) == 0 {
		compareToType = append(compareToType, DefaultCompareToType)
	}

	deltaSlice := make([]uint64, len(timestampSlice))
	for index, timestamp := range timestampSlice {

		if index == 0 {
			deltaSlice[index] = timestamp
			continue
		}

		switch compareToType[0] {
		case CompareToFirst:
			deltaSlice[index] = timestamp - timestampSlice[0]
		case CompareToPrevious:
			deltaSlice[index] = timestamp - timestampSlice[index-1]
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return deltaSlice
}

// FromTimestampSliceDelta 从时间戳的delta恢复时间戳序列
func FromTimestampSliceDelta(deltaSlice []uint64, compareToType ...CompareToType) []uint64 {

	if len(compareToType) == 0 {
		compareToType = append(compareToType, DefaultCompareToType)
	}

	timeSlice := make([]uint64, len(deltaSlice))
	for index, delta := range deltaSlice {

		if index == 0 {
			timeSlice[index] = delta
			continue
		}

		switch compareToType[0] {
		case CompareToFirst:
			timeSlice[index] = deltaSlice[0] + delta
		case CompareToPrevious:
			timeSlice[index] = timeSlice[index-1] + delta
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}

	}
	return timeSlice
}

// ------------------------------------------------ ---------------------------------------------------------------------
