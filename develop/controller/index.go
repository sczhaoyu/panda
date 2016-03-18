package controller

import (
	. "github.com/sczhaoyu/panda"
	"time"
)

func index(c *Controller) {
	c.Tpl = "view/index.html"
	c.Data["index"] = "开发工具"
	c.Data["now"] = time.Now().Local()
}
