package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gee-orm/dialect"
	"gee-orm/log"
	"gee-orm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	targetDialect, ok := dialect.GetDialect(driver)
	if !ok {
		return nil, errors.New(fmt.Sprintf("dirver:%s 's dialect dosen't exist", driver))
	}
	e = &Engine{db: db, dialect: targetDialect}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
	}
	log.Info("Close dataBase success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
