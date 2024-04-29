package mygoorm

import (
	"database/sql"
	"my_orm/log"
	"my_orm/session"
)

type Engine struct {
	db *sql.DB
}

// 创建engine根据数据库的驱动类型以及数据库的资源路径
func NewEngine(driver, source string) (*Engine, error) {
	log.Infof("db conn driver: %s, source: %s", driver, source)
	//尝试连接
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	//尝试ping databas
	err = db.Ping()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	e := &Engine{db: db}
	log.Info("db conn success")
	return e, nil
}

// 根据engine创建会话
func (e *Engine) CreateSession() *session.Session {
	return session.NewSession(e.db)
}
func (e *Engine) Close() {
	err := e.db.Close()
	if err != nil {
		log.Errorf("db close error: %v", err)
	} else {
		log.Info("db close success")
	}
}
