package delta

import (
	"fmt"
	"time"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// ToTimeSliceDelta 对时间序列做delta
func ToTimeSliceDelta(timeSlice []time.Time, compareToType CompareToType) []uint64 {
	deltaSlice := make([]uint64, len(timeSlice))
	for index, t := range timeSlice {
		if index == 0 {
			deltaSlice[index] = uint64(t.UnixMilli())
			continue
		}
		switch compareToType {
		case CompareToFirst:
			deltaSlice[index] = TimeDelta(t, timeSlice[0])
		case CompareToLast:
			deltaSlice[index] = TimeDelta(t, timeSlice[index-1])
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return deltaSlice
}

// FromTimeSliceDelta 从Time的delta恢复时间序列
func FromTimeSliceDelta(deltaSlice []uint64, compareToType CompareToType) []time.Time {
	timeSlice := make([]time.Time, 0)
	for index, delta := range deltaSlice {
		if index == 0 {
			timeSlice[index] = time.UnixMilli(int64(delta))
			continue
		}
		switch compareToType {
		case CompareToFirst:
			timeSlice[index] = time.UnixMilli(int64(deltaSlice[0] + delta))
		case CompareToLast:
			timeSlice[index] = timeSlice[index-1].Add(time.Millisecond * time.Duration(delta))
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return timeSlice
}

// TimeDelta 求两个时间的差
func TimeDelta(t1 time.Time, t2 time.Time) uint64 {
	return uint64(t1.UnixMilli() - t2.UnixMilli())
}

// ------------------------------------------------ ---------------------------------------------------------------------

// ToTimestampSliceDelta 对时间戳做delta
func ToTimestampSliceDelta(timestampSlice []uint64, compareToType CompareToType) []uint64 {
	deltaSlice := make([]uint64, len(timestampSlice))
	for index, timestamp := range timestampSlice {
		if index == 0 {
			deltaSlice[index] = timestamp
			continue
		}
		switch compareToType {
		case CompareToFirst:
			deltaSlice[index] = timestamp - timestampSlice[0]
		case CompareToLast:
			deltaSlice[index] = timestamp - timestampSlice[index-1]
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return deltaSlice
}

// FromTimestampSliceDelta 从时间戳的delta恢复时间戳序列
func FromTimestampSliceDelta(deltaSlice []uint64, compareToType CompareToType) []uint64 {
	timeSlice := make([]uint64, 0)
	for index, delta := range deltaSlice {
		if index == 0 {
			timeSlice[index] = delta
			continue
		}
		switch compareToType {
		case CompareToFirst:
			timeSlice[index] = deltaSlice[0] + delta
		case CompareToLast:
			timeSlice[index] = timeSlice[index-1] + delta
		default:
			panic(fmt.Errorf("not support compare type: %#v", compareToType))
		}
	}
	return timeSlice
}

// ------------------------------------------------ ---------------------------------------------------------------------
