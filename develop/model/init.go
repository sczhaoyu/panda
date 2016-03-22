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
	DB, _ = xorm.NewEngine("mysql", "root:root@tcp(10.0.0.252:3306)/hj?charset=utf8")
	DB.ShowSQL = true
}
func NoData(b bool) error {
	if b {
		return nil
	}
	return errors.New("not null")
}

//错误消息定义
func NoDataMsg(b bool, msg string) error {
	if b {
		return nil
	}
	return errors.New(msg)
}
