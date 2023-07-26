package foo

const (
	Max = 100
	min = 1 //ほかのパッケージから参照不可
)

func ReturnMin() int {
	return min
}
