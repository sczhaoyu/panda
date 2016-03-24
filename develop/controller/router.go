package controller

import (
	. "github.com/sczhaoyu/panda"
)

func init() {
	HandlerRouter(CASUAL, "/", index)
	HandlerRouter(CASUAL, "/table/find", tableFind)
	HandlerRouter(CASUAL, "/table/crud", crud)
	HandlerRouter(CASUAL, "/column/find", loadColumns)
	HandlerRouter(POST, "/auto/code", autoCode)
	//
	SetStaticFolder("/public/.*")
}
