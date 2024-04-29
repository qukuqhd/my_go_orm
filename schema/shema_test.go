package schema_test

import (
	"my_orm/dialect"
	"my_orm/schema"
	"testing"
)

type User struct {
	Name string `my_orm:"just some about my_orm tag info"`
	Age  int
}

var SqliteTestDialect, _ = dialect.GetDialect("sqlite3")

// 测试解析结构体为sqlite3模式对象
func TestParseSqlite3(t *testing.T) {
	schema := schema.Parse(&User{}, SqliteTestDialect) //解析得到sqlite3方言对应的模式对象
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Naem").Tag != "just some about my_orm tag info" {
		t.Fatal("failed to parse User struct")
	}
}
