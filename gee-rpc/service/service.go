package service

import (
	"reflect"
	"sync/atomic"
)

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
	numCalls  uint64
}

func (m *methodType) NumCalls() uint64 {
	//原子获取
	return atomic.LoadUint64(&m.numCalls)
}
func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	//arg may be a pointer type， or a value Type
	if m.ArgType.Kind() == reflect.Ptr {
		//指针实例值
		argv = reflect.New(m.ArgType.Elem())
	} else {
		//非指针实例值
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}

// 根据返回值类型，设置返回值
func (m *methodType) newReplyv() reflect.Value {
	//reply must be a pointer type
	replyv := reflect.New(m.ReplyType.Elem())
	switch m.ReplyType.Elem().Kind() {
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return replyv
}

type service struct {
	name   string                 //name 即映射的结构体的名称
	typ    reflect.Type           //typ 是结构体的类型
	rcvr   reflect.Value          //rcvr 即结构体的实例本身,保留 rcvr 是因为在调用时需要 rcvr 作为第 0 个参数
	method map[string]*methodType //method 是 map 类型，存储映射的结构体的所有符合条件的方法。
}

func newService(rcvr interface{}) *service {
	s := new(service)

}
