package auto_code

import (
	"fmt"
	"github.com/sczhaoyu/panda/develop/model"
	"os"
	"path/filepath"
	"text/template"
)

func CreateHtml(tableName, viewPath, idFiled string) {
	t, err := template.ParseFiles("auto_code/html.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	path, err := filepath.Abs(viewPath)
	if err == nil {
		//加载全部列
		var m map[string]interface{} = make(map[string]interface{}, 0)
		cols, err := model.FindColumns(tableName)
		//处理名称
		for i := 0; i < len(cols); i++ {

			if cols[i].Key == "PRI" {
				cols[i].Name = PSK(filert(cols[i].TableName)) + TF(cols[i].Name)
			} else {
				cols[i].Name = PSK(cols[i].Name)
			}

		}
		if err == nil {
			m["cols"] = cols
			m["table"] = filert(tableName)
			m["id"] = PSK(filert(tableName)) + TF(idFiled)
			m["header"] = `{{template "header" .}}`
			f, _ := os.Create(path + tableName + ".html")
			defer f.Close()
			err = t.Execute(f, m)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
