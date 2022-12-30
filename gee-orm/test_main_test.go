package main

import (
	"strconv"
	"testing"
)

// import (
//
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"testing"
//
// )
//
//	func Test_Engine(t *testing.T) {
//		engine, _ := NewEngine("mysql", "root:root@tcp(localhost:3306)/geeorm")
//
//		defer engine.Close()
//		s := engine.NewSession()
//		_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
//		_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
//		_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
//		result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
//		count, _ := result.RowsAffected()
//		fmt.Printf("Exec success, %d affected\n", count)
//
//		rows, err := s.Raw("Select * from User").QueryRows()
//		if err != nil {
//			return
//		}
//		var value string
//		for rows.Next() {
//			err := rows.Scan(&value)
//			if err != nil {
//				return
//			}
//			println(value)
//		}
//
// }
func TestName(t *testing.T) {
	var a []string
	for i := 0; i < 10; i++ {
		a = append(a, strconv.Itoa(i))
	}
	for _, s := range a {
		println(s)

	}
}

type TUser2 struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {

	engine, _ := NewEngine("mysql", "root:root@tcp(localhost:3306)/geeorm")
	defer engine.Close()
	s := engine.NewSession()
	s.Model(&TUser2{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
