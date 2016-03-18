package controller

import (
	. "github.com/sczhaoyu/panda"
)

func index(c *Controller) {
	c.ParseFiles("index.html")
	c.Data["title"] = "开发工具"
}
