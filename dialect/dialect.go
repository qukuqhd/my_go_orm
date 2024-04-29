package dialect

import "reflect"

var dialectMap = map[string]Dialect{} //不同的数据库类型对应不同的dialect

type Dialect interface { //定义dialect接口，用于将go的数据类型转换为对应的数据库类型消除类型不一致的问题
	DataTypeOf(dataType reflect.Value) string             //转换go的数据类型为对应连接数据库的类型
	TableExistSql(tableName string) (string, interface{}) //返回表是否存在的sql
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
