package controller

import (
	"errors"

	. "github.com/sczhaoyu/panda"
	"github.com/sczhaoyu/panda/develop/auto_code"
	"path/filepath"
)

type PrmAutoCode struct {
	Pk      string `form:"pk"`      //主键
	Action  string `form:"action"`  //控制器代码生成目录
	View    string `form:"view"`    //模板代码生成目录
	Model   string `form:"model"`   //数据库文件生成目录
	Table   string `form:"table"`   //表名
	IdFiled string `form:"idFiled"` //ID字段名
}

func autoCode(c *Controller) {
	var p PrmAutoCode
	err := c.ParseForm(&p)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	//数据库文件生成路径
	path, err := filepath.Abs(p.Model)
	if err != nil {
		c.Write(ToJson(errors.New("数据库目录不存在!")))
		return
	}
	auto_code.Path = path
	auto_code.PackageName = filepath.Base(path)
	auto_code.GetTableInfo(p.Table)
	//生成控制层代码
	auto_code.CreateController(p.Action, p.IdFiled, p.Table, p.Model)
	//生成view代码
	auto_code.CreateHtml(p.Table, p.View, p.IdFiled)
	c.Write(ToJson("生成完毕！"))
}
