package delta

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToTimeSliceDelta(t *testing.T) {
	firstTime := time.Now()
	timeSlice := []time.Time{
		firstTime,
		firstTime.Add(time.Second * 1),
		firstTime.Add(time.Second * 2),
		firstTime.Add(time.Second * 3),
		firstTime.Add(time.Second * 4),
	}

	// 都与第一个做比较
	deltaSlice := ToTimeSliceDelta(timeSlice, CompareToFirst)
	assert.Equal(t, []uint64{
		uint64(firstTime.UnixMilli()),
		uint64(1000), uint64(2000), uint64(3000), uint64(4000),
	}, deltaSlice)

	// 都与上一个做比较
	deltaSlice = ToTimeSliceDelta(timeSlice, CompareToPrevious)
	assert.Equal(t, []uint64{
		uint64(firstTime.UnixMilli()),
		uint64(1000), uint64(1000), uint64(1000), uint64(1000),
	}, deltaSlice)
}

func TestFromTimeSliceDelta(t *testing.T) {
	firstTime := time.Unix(1698509341981, 0)
	timeSlice := []time.Time{
		firstTime,
		firstTime.Add(time.Second * 1),
		firstTime.Add(time.Second * 2),
		firstTime.Add(time.Second * 3),
		firstTime.Add(time.Second * 4),
	}

	// 都与第一个做比较
	deltaSlice := ToTimeSliceDelta(timeSlice, CompareToFirst)
	recoveredSlice := FromTimeSliceDelta(deltaSlice, CompareToFirst)
	assert.Equal(t, timeSlice, recoveredSlice)

	// 都与上一个作比较
	deltaSlice = ToTimeSliceDelta(timeSlice, CompareToPrevious)
	recoveredSlice = FromTimeSliceDelta(deltaSlice, CompareToPrevious)
	assert.Equal(t, timeSlice, recoveredSlice)
}

func TestToTimestampSliceDelta(t *testing.T) {
	firstTimestamp := uint64(time.Now().UnixMilli())
	timeSlice := []uint64{
		firstTimestamp,
		firstTimestamp + 1,
		firstTimestamp + 2,
		firstTimestamp + 3,
		firstTimestamp + 4,
	}

	// 都与第一个做比较
	deltaSlice := ToTimestampSliceDelta(timeSlice, CompareToFirst)
	assert.Equal(t, []uint64{
		firstTimestamp,
		uint64(1), uint64(2), uint64(3), uint64(4),
	}, deltaSlice)

	// 都与上一个做比较
	deltaSlice = ToTimestampSliceDelta(timeSlice, CompareToPrevious)
	assert.Equal(t, []uint64{
		firstTimestamp,
		uint64(1), uint64(1), uint64(1), uint64(1),
	}, deltaSlice)
}

func TestFromTimestampSliceDelta(t *testing.T) {
	firstTime := uint64(1698509341981)
	timeSlice := []uint64{
		firstTime,
		firstTime + 1,
		firstTime + 2,
		firstTime + 3,
		firstTime + 4,
	}

	// 都与第一个做比较
	deltaSlice := ToTimestampSliceDelta(timeSlice, CompareToFirst)
	recoveredSlice := FromTimestampSliceDelta(deltaSlice, CompareToFirst)
	assert.Equal(t, timeSlice, recoveredSlice)

	// 都与上一个作比较
	deltaSlice = ToTimestampSliceDelta(timeSlice, CompareToPrevious)
	recoveredSlice = FromTimestampSliceDelta(deltaSlice, CompareToPrevious)
	assert.Equal(t, timeSlice, recoveredSlice)
}
