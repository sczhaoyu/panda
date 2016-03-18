package controller

import (
	. "github.com/sczhaoyu/panda"
)

func init() {
	HandlerRouter(CASUAL, "/", index)
	SetStaticFolder("/public/.*")
}
