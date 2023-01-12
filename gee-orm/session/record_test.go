package session

import (
	"gee-orm/dialect"
	"testing"
)

type User struct {
	Age  int
	Name string
}

func Test_Insert(t *testing.T) {
	dialect, _ := dialect.GetDialect("mysql")
	s := Session{dialect: dialect}

	_, _ = s.Insert(User{Age: 23, Name: "zhangkun"}, User{Age: 123, Name: "asdqw"})
}
