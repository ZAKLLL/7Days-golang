package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql8 struct {
}

var _ Dialect = (*mysql8)(nil)

func init() {
	RegisterDialect("mysql", &mysql8{})
}

func (s *mysql8) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

func (s *mysql8) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select TABLE_NAME from information_schema.TABLES where TABLE_SCHEMA = (SELECT DATABASE()) and TABLE_NAME = ?", args
}
