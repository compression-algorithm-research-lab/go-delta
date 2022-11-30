package delta

import (
	"reflect"
)

// ZipStructSlice 对Struct序列进行压缩
func ZipStructSlice[T any](slice []T) {

}

func diff[T any](a, b T) *StructDelta {
	reflectA := reflect.ValueOf(a)
	reflectB := reflect.ValueOf(b)
	
	switch reflectA.Type().Kind() {
	case reflect.String:

	}
}

type StructDelta struct {
	Init          any
	FieldDeltaMap map[string]*FieldDelta
}

type FieldDelta struct {
	FieldName  string
	FieldValue string
}
