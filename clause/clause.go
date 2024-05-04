package clause

import "strings"

type ClauseType int

const (
	INSERT ClauseType = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql    map[ClauseType]string        //各个类别的sql子句
	sqlVar map[ClauseType][]interface{} //各个类别的sql子句的参数
}

// 根据传入的参数和类型生成对应的sql子句
func (c *Clause) Set(name ClauseType, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[ClauseType]string)
		c.sqlVar = make(map[ClauseType][]interface{})
	}
	sql, vars := generateors[name](vars...) //生成器根据传入的一系列参数生成对应的sql语句
	c.sql[name] = sql
	c.sqlVar[name] = vars
}

//构造sql语句
func (c *Clause) Build(orders ...ClauseType) (string, []interface{}) {
	var sqls []string
	var sqlvars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order]; ok { //存在对应类型的sql子句就取出
			sqls = append(sqls, sql)
			sqlvars = append(sqlvars, c.sqlVar[order]...)
		}
	}
	return strings.Join(sqls, " "), sqlvars
}
