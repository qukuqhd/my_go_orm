package session

import (
	"my_orm/log"
	"my_orm/schema"
	"reflect"
)

func (s *Session) Model(modle interface{}) *Session {
	//如果还没有设置模式或者是模式不一致，则设置模式
	if s.refTble == nil || reflect.TypeOf(modle) != reflect.TypeOf(s.refTble) {
		s.refTble = schema.Parse(modle, s.dialect) //根据选择的数据库方言来解析结构体为对应的模式
	}
	return s
}
func (s *Session) RefTable() *schema.Schema {
	if s.refTble == nil {
		log.Error("model is not set")
	}
	return s.refTble
}
