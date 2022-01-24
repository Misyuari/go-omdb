package config

import (
	"database/sql"
	"github.com/misyuari/go-omdb/helper"
	"time"
)

func NewDB() {
	db, err := sql.Open(GlobalConfig.Env.DbDriver, GlobalConfig.Env.DbUri)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	GlobalConfig.DBConnection = db
}
