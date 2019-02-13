package mysql

import (
	"database/sql"
	"fmt"
	// "runtime"

	"github.com/go-gorp/gorp"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

//Init ...
func init() {
	dbinfo := ""
	addr := ""

	user := "root"
	pass := "3wes4r55"
	addr = "47.106.115.69:3307" //测试服地址

	name := "alipay"
	dbinfo = fmt.Sprintf("%s:%s@tcp(%s)/%s?interpolateParams=true", user, pass, addr, name)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dataSourceName)
	// defer db.Close()
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)
	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap, nil
}

func CloseDB() {

}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
