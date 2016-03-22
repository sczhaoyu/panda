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

//更新函数
func UpdateCtrFunc(tableName, modelName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//更新` + tf + `
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
func ViewCtrFunc(tableName, modelName string) string {
	psk := PSK(filert(tableName))
	tf := TF(filert(tableName))
	fun := `
	//列表分页查询` + tf + `
	func ` + psk + `View(c *Controller) {
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
	fun += CreateCtrFunc(tableName, modelPath)
	fun += ViewCtrFunc(tableName, modelPath)
	fun += UpdateCtrFunc(tableName, modelPath)
	fun += DeleteCtrFunc(tableName, modelPath)
	WriteFile(codePath, tableName+".go", fun)
}
