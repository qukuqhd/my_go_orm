package session

import (
	"my_orm/log"
	"reflect"
)

// 钩子容器sql执行的生命周期
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// 调用钩子函数的方法
func (s *Session) CallMethod(method string, val interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method) //获取对象里面对应的方法
	if val != nil {
		fm = reflect.ValueOf(val).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
}
