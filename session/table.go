package session

import (
	"fmt"
	"my_orm/log"
	"my_orm/schema"
	"reflect"
	"strings"
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

// 创建会话包含模式对应的表
func (s *Session) CreateTable() error {
	table_schema := s.RefTable() //获取模式对象
	var columns []string         //sql语句描述表结构每一个字段的一行
	for _, field := range table_schema.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")                                                   //为一个表字段的描述添加逗号分隔形成对所有字段的描述
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table_schema.Name, desc)).Exec() //格式化为最终的sql语句执行
	return err
}

// 删除会话包含模式对应的表
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

// 判断模式对应的表是否存在
func (s *Session) HasTable() bool {
	sql, val := s.dialect.TableExistSql(s.RefTable().Name)
	row := s.Raw(sql, val).QueryRow()
	var tmp string
	row.Scan(tmp)
	return tmp == s.RefTable().Name
}
