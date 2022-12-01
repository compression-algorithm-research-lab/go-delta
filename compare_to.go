package delta

// CompareToType delta比较的时候，当前元素是与哪个元素比较
type CompareToType int

const (

	// CompareToFirst 将当前元素与序列的第一个元素比较
	CompareToFirst CompareToType = iota

	// CompareToLast 将当前元素与上一个元素做比较
	CompareToLast
)
