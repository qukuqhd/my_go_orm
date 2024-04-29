package schema

import (
	"go/ast"
	"my_orm/dialect"
	"reflect"
)

// 字段结构体
type Field struct {
	Name string //字段名称
	Type string //字段类型
	Tag  string //字段的标签
}

// 模式转换的结构体
type Schema struct {
	Model      interface{}       //要转换的模型
	Name       string            //模型的名称
	Fields     []*Field          //各个字段结构体
	FieldNames []string          //各个字段的名称
	filedMap   map[string]*Field //字段映射
}

func (s *Schema) GetField(name string) *Field {
	return s.filedMap[name]
}

// 根据选择的数据库方言来解析结构体为对应的模式
func Parse(model interface{}, dialect dialect.Dialect) *Schema {
	modeleType := reflect.Indirect(reflect.ValueOf(model)).Type() //获取model的类型
	//创建模式对象
	schema := &Schema{
		Model:    model,
		Name:     modeleType.Name(),
		filedMap: make(map[string]*Field),
	}
	for i := 0; i < modeleType.NumField(); i++ { //遍历结构体的字段
		p := modeleType.Field(i) //获取当前的字段
		//过滤掉非导出和非公开的字段
		if !p.Anonymous && ast.IsExported(p.Name) {
			//创建字段
			filed := &Field{
				Name: p.Name,
				Type: dialect.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			//寻找标签
			if v, ok := p.Tag.Lookup("my_orm"); ok {
				filed.Tag = v
			}
			//添加字段信息到模式对象
			schema.Fields = append(schema.Fields, filed)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.filedMap[p.Name] = filed
		}
	}
	return schema
}
