package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DB *xorm.Engine
)

func init() {
	DB, _ = xorm.NewEngine("mysql", "")
	DB.ShowSQL = true
}
func noData(b bool) error {
	if b {
		return nil
	}
	return errors.New("not null")
}
