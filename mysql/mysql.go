package mysql

import (
	"fmt"
	"minipro/setting"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//Init the database connection
func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.Max_conns)      //Set max open connections
	db.SetMaxIdleConns(cfg.Max_idle_conns) //Set max idle connections
	return
}

//close the database connection
func Close() {
	_ = db.Close()
}
