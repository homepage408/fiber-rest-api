package configuration

import (
	"database/sql"
	"fiber-rest-api/helper"
	"fmt"
	"time"
)

func NewDb(params *Configuration) *sql.DB {

	mysql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", params.Database.Username, params.Database.Password, params.Database.Host, params.Database.Port, params.Database.Name)
	fmt.Println("Mysql :", mysql)
	db, err := sql.Open("mysql", mysql)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
