package session

import (
	"my_orm/clause"
	"reflect"
)

// 插入对象的数据到会话连接的数据库
func (s *Session) Insert(vals ...interface{}) (int64, error) {
	recordVals := make([]interface{}, 0)
	for _, val := range recordVals { //遍历要插入的变量
		//生成对象对应的模式
		table := s.Model(val).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordVals = append(recordVals, table.RecordValues(val)) //转换sql语句的参数
	}
	s.clause.Set(clause.VALUES, recordVals...)                //设置所有的参数
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES) //生成sql
	result, err := s.Raw(sql, vars...).Exec()                 //执行
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 根据传入对象查找数据库
func (s *Session) Find(val interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(val))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}
	for rows.Next() { //遍历查询结果集合
		dest := reflect.New(destType).Elem()
		var vals []interface{}
		for _, name := range table.FieldNames {
			vals = append(vals, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(vals...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

// 根据传入对象更新数据
func (s *Session) Update(kv ...interface{}) (int64, error) {
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 删除数据库记录
func (s *Session) Delete() (int64, error) {
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.Raw(sql, vars).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// 获取会话连接的数据库的表的记录条数
func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, vars...).QueryRow()
	var tmp int64
	if err := row.Scan(&tmp); err != nil {
		return 0, err
	}
	return tmp, nil
}
