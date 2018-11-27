package db

import (

	// mysqldriver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Connect public method
func Connect() (db *xorm.Engine, err error) {
	db, err = xorm.NewEngine("mysql", "root:123@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}
