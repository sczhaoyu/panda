package controller

import (
	. "github.com/sczhaoyu/panda"
	"github.com/sczhaoyu/panda/develop/model"
	"strconv"
)

//AuthCode路由表
//HandlerRouter(CASUAL, "/auth_code/view", authCodeView)
//HandlerRouter(CASUAL, "/auth_code/byid",authCodeById)
//HandlerRouter(CASUAL, "/auth_code/list", authCodeList)
//HandlerRouter(CASUAL, "/auth_code/update", authCodeUpdate)
//HandlerRouter(CASUAL, "/auth_code/delete",authCodeDelete)
//HandlerRouter(CASUAL, "/auth_code/save", authCodeSave)

//保存AuthCode
//url:/auth_code/save
//HandlerRouter(CASUAL, "/auth_code/save", authCodeSave)
func authCodeSave(c *Controller) {
	var authCode model.AuthCode
	err := c.ParseForm(&authCode)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	err = authCode.Save()
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson("保存成功！"))
}

//VIEW：AuthCode
//url:/auth_code/view
//HandlerRouter(CASUAL, "/auth_code/view", authCodeView)
func authCodeView(c *Controller) {
	c.Tpl = "hj_auth_code.html"
	c.Data["title"] = "AuthCode"
}

//根据ID获取信息AuthCode
//url:/auth_code/byid
//HandlerRouter(CASUAL, "/auth_code/byid", authCodeById)
func authCodeById(c *Controller) {
	var authCode model.AuthCode
	err := c.ParseForm(&authCode)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	ret, err := model.GetAuthCodeById(authCode.Id)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson(ret))

}

//列表分页查询AuthCode
//url:/auth_code/list
//HandlerRouter(CASUAL, "/auth_code/list", authCodeList)
func authCodeList(c *Controller) {
	page, _ := strconv.Atoi(c.FormValue("page"))
	rows, _ := strconv.Atoi(c.FormValue("rows"))
	ret, count, err := model.FindAuthCodePages(page, rows)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.WriteJson(GetPagination(count, page, rows, ret))
}

//更新AuthCode
//url:/auth_code/update
//HandlerRouter(CASUAL, "/auth_code/update", authCodeUpdate)
func authCodeUpdate(c *Controller) {
	var authCode model.AuthCode
	err := c.ParseForm(&authCode)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	err = authCode.Update()
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson("更新成功！"))
}

//删除AuthCode
//url:/auth_code/delete
//HandlerRouter(CASUAL, "/auth_code/delete", authCodeDelete)
func authCodeDelete(c *Controller) {
	var authCode model.AuthCode
	err := c.ParseForm(&authCode)
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	err = authCode.Delete()
	if err != nil {
		c.Write(ToJson(err))
		return
	}
	c.Write(ToJson("删除成功！"))

}
