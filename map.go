package iocv

import (
	"fmt"
	"reflect"
)

type MapContainer struct {
	name   string
	storge map[string]Object
}

// Registry 注册
func (m *MapContainer) Registry(name string, obj Object) {
	if name == "" {
		name = reflect.TypeOf(obj).String()
	}
	m.storge[name] = obj
}

func (m *MapContainer) Get(name string) any {
	return m.storge[name]
}

func (m *MapContainer) Init() error {
	for k, v := range m.storge {
		err := v.Init()
		if err != nil {
			return fmt.Errorf("%s init 错误: %s", k, err)
		}
		fmt.Println("init success 成功:", k)
	}
	return nil
}
