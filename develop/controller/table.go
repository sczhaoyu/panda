package controller

import (
	. "github.com/sczhaoyu/panda"
	"github.com/sczhaoyu/panda/develop/model"
	"strconv"
)

func tableFind(c *Controller) {
	page, _ := strconv.Atoi(c.FormValue("page"))
	rows, _ := strconv.Atoi(c.FormValue("rows"))
	ret, count, err := model.FindTable(page, rows)
	if err != nil {
		c.WriteJson(err)
		return
	}

	c.WriteJson(GetPagination(count, page, rows, ret))
}
