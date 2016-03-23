package auto_code

import (
	"os"
	"path/filepath"
	"strings"
)

//创建函数
func CreateCtrFunc(tableName, modelName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//保存` + tf + `
	//url:/` + filert(tableName) + `/save
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/save", ` + psk + `Save)
	func ` + psk + `Save(c *Controller) {
		var ` + psk + ` ` + modelName + `.` + tf + `
		err:=c.ParseForm(&` + psk + `)
		if err!=nil {
			c.Write(ToJson(err))
			return
		}
		err=` + psk + `.Save()
		if err!=nil {
			c.Write(ToJson(err))
			return 
		}
		c.Write(ToJson("保存成功！"))
	}`
	return fun
}

//根据ID获取函数
func ByIdCtrFunc(tableName, modelName, idName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//根据ID获取信息` + tf + `
	//url:/` + filert(tableName) + `/byid
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/byid", ` + psk + `ById)
	func ` + psk + `ById(c *Controller) {
	     var  ` + psk + ` ` + modelName + `.` + tf + `
	     err:=c.ParseForm(&` + psk + `)
	     if err!=nil {
	     	 c.Write(ToJson(err))
	     	 return 
	     }
	     ret,err:=` + modelName + `.Get` + tf + `ById(` + psk + `.` + TF(idName) + `)
	     if err!=nil {
	     	 c.Write(ToJson(err))
	     	 return 
	     }
		c.Write(ToJson(ret))
		 
	}`
	return fun
}

//更新函数
func UpdateCtrFunc(tableName, modelName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//更新` + tf + `
	//url:/` + filert(tableName) + `/update
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/update", ` + psk + `Update)
	func ` + psk + `Update(c *Controller) {
		var ` + psk + ` ` + modelName + `.` + tf + `
		err:=c.ParseForm(&` + psk + `)
		if err!=nil {
			c.Write(ToJson(err))
			return
		}
		err=` + psk + `.Update()
		if err!=nil {
			c.Write(ToJson(err))
			return 
		}
		c.Write(ToJson("更新成功！"))
	}`
	return fun
}

//列表函数
func ListCtrFunc(tableName, modelName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//列表分页查询` + tf + `
	//url:/` + filert(tableName) + `/list
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/list", ` + psk + `List)
	func ` + psk + `List(c *Controller) {
	   page, _ := strconv.Atoi(c.FormValue("page"))
	   rows, _ := strconv.Atoi(c.FormValue("rows"))
	   ret,count,err:= ` + modelName + `.Find` + tf + `Pages(page, rows)
	   if err!=nil {
	   	    c.Write(ToJson(err))
			return 
	   }
	   c.WriteJson(GetPagination(count, page, rows,ret))
	}`
	return fun
}

//删除函数
func DeleteCtrFunc(tableName, modelName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//删除` + tf + `
	//url:/` + filert(tableName) + `/delete
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/delete", ` + psk + `Delete)
	func ` + psk + `Delete(c *Controller) {
		var ` + psk + `   ` + modelName + `.` + tf + `
		err:=c.ParseForm(&` + psk + `)
		if err!=nil {
			c.Write(ToJson(err))
			return
		}
		err= ` + psk + `.Delete()
		if err!=nil {
			c.Write(ToJson(err))
			return
		}
		c.Write(ToJson("删除成功！"))
		 
	}`
	return fun
}
func ViewCtrFunc(tableName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//VIEW：` + tf + `
	//url:/` + filert(tableName) + `/view
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/view", ` + psk + `View)
    func ` + psk + `View(c *Controller) {
	c.Tpl = "` + tableName + `.html"
	c.Data["title"] = "` + tf + `"
    }
  `
	return fun
}

//@codePath代码生成路径
//@key主键
//@tableName 表名
//@modelPath 数据库绝对路径
func CreateController(codePath, key, tableName, modelPath string) {
	mp, _ := filepath.Abs(modelPath)
	mp = strings.Replace(mp, "\\", "/", -1)
	gp := strings.Replace(os.Getenv("GOPATH"), "\\", "/", -1)
	mp = strings.Trim(strings.Replace(mp, gp+"/src/", "", -1), "/")

	modelPath = strings.Trim(modelPath, "/")

	fun := `
package ` + strings.Trim(codePath, "/") + `

import(
   . "github.com/sczhaoyu/panda"
  "` + mp + `"
   "strconv" 
)

`
	psk := PSK(filert(tableName))
	fun += `
	//` + TF(filert(tableName)) + `路由表
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/view", ` + psk + `View)
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/byid",` + psk + `ById)
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/list", ` + psk + `List)
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/update", ` + psk + `Update)
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/delete",` + psk + `Delete)
	//HandlerRouter(CASUAL, "/` + filert(tableName) + `/save", ` + psk + `Save)
    `
	fun += CreateCtrFunc(tableName, modelPath)
	fun += ViewCtrFunc(tableName)
	fun += ByIdCtrFunc(tableName, modelPath, key)
	fun += ListCtrFunc(tableName, modelPath)
	fun += UpdateCtrFunc(tableName, modelPath)
	fun += DeleteCtrFunc(tableName, modelPath)
	WriteFile(codePath, tableName+".go", fun)
}
