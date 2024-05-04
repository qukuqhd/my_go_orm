package clause

import (
	"fmt"
	"strings"
)

// 生成器根据传入的一系列参数生成对应的sql语句
type generateor func(vals ...interface{}) (string, []interface{})

var generateors map[ClauseType]generateor

func init() {
	generateors = make(map[ClauseType]generateor)
	generateors[INSERT] = _insert
	generateors[VALUES] = _values
	generateors[SELECT] = _select
	generateors[LIMIT] = _limit
	generateors[WHERE] = _where
	generateors[ORDERBY] = _orderBy
	generateors[UPDATE] = _update
	generateors[DELETE] = _delete
	generateors[COUNT] = _count
}

// 按照需要的位数生成相应个占位
func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")

}

// 生成插入的语句
func _insert(vals ...interface{}) (string, []interface{}) {
	tableName := vals[0]                            //插入操作的第一个属性应该是表名称
	fields := strings.Join(vals[1].([]string), ",") //第二个属性应该是字段名称列表
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

// 生成value的内容
func _values(vals ...interface{}) (string, []interface{}) {
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	//开始的values语句应该有VALLUES
	bindStr = genBindVars(len(vals)) //获取占位符
	sql.WriteString("VALUES ")
	for i, val := range vals {
		v := val.([]interface{})
		sql.WriteString(fmt.Sprintf("%v", bindStr))
		if i != len(vals)-1 { //不是最后一项填充参数应该添加逗号
			sql.WriteString(",")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

// 生成select的子句
func _select(vals ...interface{}) (string, []interface{}) {
	table_name := vals[0]
	fileds := strings.Join(vals[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fileds, table_name), []interface{}{}
}

// 生成limit的子句
func _limit(vals ...interface{}) (string, []interface{}) {
	return "LIMIT ?", vals
}

// 生成where的子句
func _where(vals ...interface{}) (string, []interface{}) {
	desc, vars := vals[0], vals[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

// 生成orderby的子句
func _orderBy(vals ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", vals[0]), []interface{}{}
}

// 生成update的子句
func _update(vals ...interface{}) (string, []interface{}) {
	tableName := vals[0]
	m := vals[1].(map[string]interface{})
	var keys []string
	var vars []interface{}
	for k, v := range m {
		keys = append(keys, k)
		vars = append(vars, v)
	}
	return fmt.Sprintf("UPDATE %s SET %s ", tableName, strings.Join(keys, ",")), vars
}

// 生成delete的子句
func _delete(vals ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s ", vals[0]), []interface{}{}
}

// 生成count的子句
func _count(vals ...interface{}) (string, []interface{}) {
	return _select(vals[0], []string{"count(*)"})
}
