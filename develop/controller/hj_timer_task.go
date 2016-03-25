package controller

import (
	. "github.com/sczhaoyu/panda"
	"github.com/sczhaoyu/panda/develop/model"
	"strconv"
)

//TimerTask路由表
//HandlerRouter(CASUAL, "/timer_task/view", timerTaskView)
//HandlerRouter(CASUAL, "/timer_task/byid",timerTaskById)
//HandlerRouter(CASUAL, "/timer_task/list", timerTaskList)
//HandlerRouter(CASUAL, "/timer_task/update", timerTaskUpdate)
//HandlerRouter(CASUAL, "/timer_task/delete",timerTaskDelete)
//HandlerRouter(CASUAL, "/timer_task/save", timerTaskSave)

//保存TimerTask
//url:/timer_task/save
//HandlerRouter(CASUAL, "/timer_task/save", timerTaskSave)
func timerTaskSave(c *Controller) {
	var timerTask model.TimerTask
	err := c.ParseForm(&timerTask)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	err = timerTask.Save()
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson("保存成功！"))
}

//VIEW：TimerTask
//url:/timer_task/view
//HandlerRouter(CASUAL, "/timer_task/view", timerTaskView)
func timerTaskView(c *Controller) {
	c.Tpl = "hj_timer_task.html"
	c.Data["title"] = "TimerTask"
}

//根据ID获取信息TimerTask
//url:/timer_task/byid
//HandlerRouter(CASUAL, "/timer_task/byid", timerTaskById)
func timerTaskById(c *Controller) {
	var timerTask model.TimerTask
	err := c.ParseForm(&timerTask)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	ret, err := model.GetTimerTaskById(timerTask.Id)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson(ret))

}

//列表分页查询TimerTask
//url:/timer_task/list
//HandlerRouter(CASUAL, "/timer_task/list", timerTaskList)
func timerTaskList(c *Controller) {
	page, _ := strconv.Atoi(c.FormValue("page"))
	rows, _ := strconv.Atoi(c.FormValue("rows"))
	ret, count, err := model.FindTimerTaskPages(page, rows)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.WriteJson(GetPagination(count, page, rows, ret))
}

//更新TimerTask
//url:/timer_task/update
//HandlerRouter(CASUAL, "/timer_task/update", timerTaskUpdate)
func timerTaskUpdate(c *Controller) {
	var timerTask model.TimerTask
	err := c.ParseForm(&timerTask)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	err = timerTask.Update()
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson("更新成功！"))
}

//删除TimerTask
//url:/timer_task/delete
//HandlerRouter(CASUAL, "/timer_task/delete", timerTaskDelete)
func timerTaskDelete(c *Controller) {
	var timerTask model.TimerTask
	err := c.ParseForm(&timerTask)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	err = timerTask.Delete()
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson("删除成功！"))

}
