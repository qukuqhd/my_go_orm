package session

import (
	"database/sql"
	"my_orm/log"
	"strings"
)

// 数据库会话结构体
type Session struct {
	db      *sql.DB         // 数据库连接
	sql     strings.Builder // sql语句
	sqlVars []interface{}   //sql 参数
}

// 根据连接来创建会话
func NewSession(db *sql.DB) *Session {
	return &Session{db: db}
}

// 清空会话的sql
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// 获取会话的数据库连接
func (s *Session) DB() *sql.DB {
	return s.db
}

// 设置会话的sql
func (s *Session) Raw(sql string, values ...interface{}) {
	s.sql.WriteString(sql)                   //添加sql
	s.sql.WriteString(" ")                   //添加空格，保证sql语句可以正确地被三百执行
	s.sqlVars = append(s.sqlVars, values...) //添加sql语句参数到切片末尾
}

// 执行数据库会话的sql
func (s *Session) Exec() (sql.Result, error) {
	defer s.Clear() //移除执行了的sql
	log.Info(s.sql.String(), s.sqlVars)
	result, err := s.DB().Exec(s.sql.String(), s.sqlVars...) //执行sql
	if err != nil {
		log.Error(err)
	}
	return result, err
}

// 执行数据库会话的sql进行查询单条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql, s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// 执行数据库会话的sql进行查询多条记录
func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql, s.sqlVars)
	rows, err := s.DB().Query(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}
