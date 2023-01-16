package session

import (
	"gee-orm/clause"
	"reflect"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	//recordValues := make([]interface{}, 0)
	var recordValues []interface{}
	table := s.Model(values[0]).RefTable()
	s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
	for _, value := range values {
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	ret, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}

	return ret.RowsAffected()
}

// 传入一个切片指针，查询的结果保存在切片中s
// s := geeorm.NewEngine("sqlite3", "gee.db").NewSession()
// var users []User
// s.Find(&users);
func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	//返回 destSlice 的其中一个元素的 类型
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.WHERE, clause.LIMIT, clause.ORDERBY)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
