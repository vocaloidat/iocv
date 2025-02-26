package iocv

type Container interface {
	// Registry 注册对象
	Registry(name string, obj Object)
	// Get 获取对象 返回对象
	Get(name string) any
	// Init 初始化容器
	Init() error
}

type Object interface {
	Init() error
}
