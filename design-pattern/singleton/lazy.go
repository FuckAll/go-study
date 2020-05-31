package singleton

var instanceLazy = new(Example)

// GetInstanceLazy 懒汉模式，不会有并发问题
func GetInstanceLazy() *Example {
	return instanceLazy
}
