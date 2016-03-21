package controller

import (
	. "github.com/sczhaoyu/panda"
	"github.com/sczhaoyu/panda/develop/model"
	"strconv"
)

func tableFind(c *Controller) {
	page, _ := strconv.Atoi(c.FormValue("page"))
	rows, _ := strconv.Atoi(c.FormValue("rows"))
	ret, count, err := model.FindTable("hj", page, rows)
	if err != nil {
		c.WriteJson(err)
		return
	}

	p := GetPagination(count, page, rows)
	p.Ret = ret
	c.WriteJson(p)
}
