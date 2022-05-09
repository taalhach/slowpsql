package database

import (
	"fmt"

	"github.com/taalhach/slowpsql/internal/server/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbSession struct {
	*gorm.DB
}

var Db DbSession

func MustConnectDB(cfg *configs.DatabaseConfig) {
	db, err := gorm.Open(postgres.Open(cfg.ConnString()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("db connection configs failed, the error is '%v'", err))
	}

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)

	// start debug mod
	if cfg.ShowSql {
		db.Debug()
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	Db = DbSession{
		db,
	}
}

//EnablePgStatStatementsExt enable pg_stat_statements extension
func EnablePgStatStatementsExt() error {
	err := Db.Exec("CREATE  EXTENSION IF NOT EXISTS pg_stat_statements;").Error
	return err
}
