package singleton

type Example struct {
	name string
}

var instance *Example

// GetInstance 懒汉模式，会有并发问题
func GetInstance() *Example {
	if instance == nil {
		instance = new(Example)
	}
	return instance
}
