package model

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sczhaoyu/panda/develop/config"
)

var (
	DB *xorm.Engine
)

func init() {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.DB("name").String(), config.DB("pwd").String(), config.DB("address").String(), config.DB("db").String())
	DB, _ = xorm.NewEngine(config.DB("type").String(), url)
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
