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
	HandlerRouter(CASUAL, "/auth_code/view", authCodeView)
	HandlerRouter(CASUAL, "/auth_code/byid", authCodeById)
	HandlerRouter(CASUAL, "/auth_code/list", authCodeList)
	HandlerRouter(CASUAL, "/auth_code/update", authCodeUpdate)
	HandlerRouter(CASUAL, "/auth_code/delete", authCodeDelete)
	HandlerRouter(CASUAL, "/auth_code/save", authCodeSave)
	SetStaticFolder("/public/.*")
}
