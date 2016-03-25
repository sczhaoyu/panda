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
	SetStaticFolder("/public/.*")
	//TimerTask路由表
	HandlerRouter(CASUAL, "/timer_task/view", timerTaskView)
	HandlerRouter(CASUAL, "/timer_task/byid", timerTaskById)
	HandlerRouter(CASUAL, "/timer_task/list", timerTaskList)
	HandlerRouter(CASUAL, "/timer_task/update", timerTaskUpdate)
	HandlerRouter(CASUAL, "/timer_task/delete", timerTaskDelete)
	HandlerRouter(CASUAL, "/timer_task/save", timerTaskSave)
}
