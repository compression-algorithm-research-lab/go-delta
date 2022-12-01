package delta

import "time"

// ------------------------------------------------ ---------------------------------------------------------------------

// TimeSliceDelta 对时间序列做delta
func TimeSliceDelta(timeSlice []time.Time, compareToType CompareToType) []uint64 {
	result := make([]uint64, len(timeSlice))
	for index, t := range timeSlice {
		if index == 0 {
			result[index] = uint64(t.UnixMilli())
			continue
		}
		if compareToType == CompareToFirst {
			result[index] = TimeDelta(t, timeSlice[0])
		} else if compareToType == CompareToLast {
			result[index] = TimeDelta(t, timeSlice[index-1])
		}
	}
	return result
}

// FromTimeSliceDelta 从Time的delta恢复时间序列
func FromTimeSliceDelta(deltaSlice []uint64, compareToType CompareToType) []time.Time {
	result := make([]time.Time, 0)
	for index, delta := range deltaSlice {
		if index == 0 {
			result[index] = time.UnixMilli(int64(delta))
			continue
		}
		if compareToType == CompareToFirst {
			result[index] = time.UnixMilli(int64(deltaSlice[0] + delta))
		} else if compareToType == CompareToLast {
			result[index] = result[index-1].Add(time.Millisecond * time.Duration(delta))
		}
	}
	return result
}

func TimeDelta(t1 time.Time, t2 time.Time) uint64 {
	return uint64(t1.UnixMilli() - t2.UnixMilli())
}

// ------------------------------------------------ ---------------------------------------------------------------------

// TimestampSliceDelta 对时间戳做delta
func TimestampSliceDelta(timestampSlice []uint64, compareToType CompareToType) []uint64 {
	result := make([]uint64, len(timestampSlice))
	for index, timestamp := range timestampSlice {
		if index == 0 {
			result[index] = timestamp
			continue
		}
		if compareToType == CompareToFirst {
			result[index] = timestampSlice[index] - timestampSlice[0]
		} else if compareToType == CompareToLast {
			result[index] = timestampSlice[index] - timestampSlice[index-1]
		}
	}
	return result
}

// FromTimestampSliceDelta 从时间戳的delta恢复时间戳序列
func FromTimestampSliceDelta(deltaSlice []uint64, compareToType CompareToType) []uint64 {
	result := make([]uint64, 0)
	for index, delta := range deltaSlice {
		if index == 0 {
			result[index] = delta
			continue
		}
		if compareToType == CompareToFirst {
			result[index] = deltaSlice[0] + delta
		} else if compareToType == CompareToLast {
			result[index] = result[index-1] + delta
		}
	}
	return result
}

// ------------------------------------------------ ---------------------------------------------------------------------