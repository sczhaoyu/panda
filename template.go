package panda

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	//模板对应关系对照
	templates map[string][]string = make(map[string][]string, 0)
)

//读取模板之间的全部对照关系
func readTemplates() {
	//模板名称对应文件名
	var templateName map[string]string = make(map[string]string, 0)
	filepath.Walk(ViewPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		//替换目录中的\符号，防止出错
		path = strings.Replace(path, "\\", "/", -1)
		if err == nil {
			ret := string(data)
			//分析当前文件是否有模板名称
			var r = regexp.MustCompile(`{{define.*}}`)
			//提取出所有的表达式标签
			names := r.FindStringSubmatch(ret)
			//整理对照关系
			for i := 0; i < len(names); i++ {
				//清除标签信息，只保留名称

				reg := regexp.MustCompile(`({{.*define)|"|}}| `)
				name := reg.ReplaceAllString(names[i], "")
				//判断html的标签命名是否冲突，如果冲突抛出异常
				_, ok := templateName[name]
				if ok {
					panic(errors.New("path:" + path + "  tagName [" + name + "] already exist?"))
				}
				//保存标签与文件名字的对照关系
				templateName[name] = path
				//---end
			}
			//开始读取模板文件中需要对应的标签信息
			r = regexp.MustCompile(`{{template.*}}`)
			//提取出所有的表达式标签
			names = r.FindAllString(ret, -1)
			for i := 0; i < len(names); i++ {
				//清除标签信息，只保留名称
				reg := regexp.MustCompile(`({{.*template)|"|\.|}}| `)
				name := reg.ReplaceAllString(names[i], "")
				templates[path] = append(templates[path], name)
				//---end
			}
		}
		return nil
	})
	//文件对照
	for k, v := range templates {
		for i := 0; i < len(v); i++ {
			val, ok := templateName[v[i]]
			if ok {
				templates[k][i] = val
			}
		}
		//生成所有对应模板文件
		tpls := templates[k]
		var ret []string = make([]string, 0, 0)
		for i := 0; i < len(tpls); i++ {
			findTplsPaths(tpls[i], &ret)

		}
		templates[k] = append(templates[k], ret...)
		//生成所有对应模板文件  END
	}

}

//递归查找模板
func findTplsPaths(tplPath string, paths *[]string) {
	tpls := templates[tplPath]
	if len(tpls) == 0 {
		return
	}
	*paths = append(*paths, tpls...)
	for i := 0; i < len(tpls); i++ {
		//递归判断里面的模板是否有加载路径
		tmp := templates[tpls[i]]
		for j := 0; j < len(tmp); j++ {
			findTplsPaths(tmp[j], paths)
		}
	}

}
