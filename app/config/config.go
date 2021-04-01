package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	NameDB     = "db_mhsApi"
	NameDriver = "mysql"
	PassDB     = "123456789"
	UserName   = "saul"
)

func InitMysql() (db *sql.DB) {
	dsn := fmt.Sprintf("%v:%v@/%v?parseTime=true", UserName, PassDB, NameDB)
	conn, err := sql.Open(NameDriver, dsn)

	if err != nil {
		log.Fatal(err)
		return
	}
	return conn
}
