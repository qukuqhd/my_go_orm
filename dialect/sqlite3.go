/**
* 对sqlite3的支持
**/
package dialect

import (
	"reflect"
	"time"
)

type sqlite3 struct {
}

// DataTypeOf implements Dialect.
func (s sqlite3) DataTypeOf(dataType reflect.Value) string {
	switch dataType.Kind() { //根据传入的数据类型返回对应的数据库类型
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Uint64, reflect.Int64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := dataType.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	//没有匹配的数据类型，直接panic说明还没有支持该类型
	panic("unsupported type: " + dataType.Type().Name())
}

// TableExistSql implements Dialect.
func (s sqlite3) TableExistSql(tableName string) (string, interface{}) {
	args := []interface{}{tableName} //传入的表名转换为[]interface{}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}

var _ Dialect = sqlite3{}

// 初始化注册sqlite方言支持
func init() {
	RegisterDialect("sqlite3", sqlite3{})
}
