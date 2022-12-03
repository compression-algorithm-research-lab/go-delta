package delta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToByteSliceDelta(t *testing.T) {
	bytes := []byte{1, 2, 3, 4, 5}
	//t.Log(bytes)
	deltaSlice := ToByteSliceDelta(bytes, CompareToLast)
	//t.Log(deltaSlice)
	originalBytes := FromByteSliceDelta(deltaSlice, CompareToLast)
	//t.Log(originalBytes)
	assert.Equal(t, bytes, originalBytes)
}
