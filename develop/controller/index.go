package controller

import (
	. "github.com/sczhaoyu/panda"
	"github.com/sczhaoyu/panda/develop/model"
	"strconv"
)

func index(c *Controller) {
	c.Tpl = "index.html"
	c.Data["title"] = "开发工具"
}
func crud(c *Controller) {
	c.Tpl = "crud.html"
	c.Data["title"] = "代码生成"
	c.Data["name"] = c.FormValue("name")
	ret, err := model.FindColumns("hj", c.FormValue("name"))
	if err == nil {
		c.Data["columns"] = ret
	}
}
func loadColumns(c *Controller) {
	page, _ := strconv.Atoi(c.FormValue("page"))
	rows, _ := strconv.Atoi(c.FormValue("rows"))
	ret, err := model.FindColumns("hj", c.FormValue("name"))
	if err != nil {
		c.WriteJson(err)
		return
	}
	c.WriteJson(GetPagination(int64(len(ret)), page, rows, ret))
}
