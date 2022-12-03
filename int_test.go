package delta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnzipIntegerSlice(t *testing.T) {
	slice := make([]int, 0)
	for i := 0; i < 100; i++ {
		slice = append(slice, 10000+i*10)
	}
	//t.Log(slice)

	t1 := ToIntegerSliceDelta(slice, CompareToLast)
	//t.Log(t1)

	t2 := FromIntegerSliceDelta(t1, CompareToLast)
	//t.Log(t2)

	assert.Equal(t, slice, t2)
}
