package session

import "gee-orm/clause"

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
