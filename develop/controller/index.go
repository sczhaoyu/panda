package controller

import (
	. "github.com/sczhaoyu/panda"
	"time"
)

func index(c *Controller) {
	c.ParseFiles("index.html")
	c.Data["index"] = "开发工具"
	c.Data["now"] = time.Now().Local()
}
