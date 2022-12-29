package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, _ := sql.Open("mysql", "root:gV0qUTyimTLhiqhiK@tcp(10.91.122.5:3307)/jyfx")
	defer func() { _ = db.Close() }()

	//row := db.QueryRow()

	rows, err := db.Query("show tables")
	if err != nil {
		return
	}
	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err == nil {
			log.Println(name)
		} else {
			break
		}

	}
}
