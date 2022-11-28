package delta

import (
	"testing"
)

func TestUnzipIntegerSlice(t *testing.T) {
	slice := make([]int, 0)
	for i := 0; i < 100; i++ {
		slice = append(slice, 10000-i*10)
	}
	t1 := ZipIntegerSlice(slice)
	t.Log(t1)

	t2 := UnzipIntegerSlice(t1)
	t.Log(t2)
}
