package panda

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const (
	GET    = "GET"  //GET请求
	POST   = "POST" //POST请求
	CASUAL = ""     //任意请求
)

//路由器
type Router struct {
	Method string            //GET|POST请求方式
	Path   string            //请求路径
	Action func(*Controller) //控制器
}

var (
	//路由器列表
	routers map[string]*Router = make(map[string]*Router, 0)
	//静态资源目录列表,表达式
	staticFolder []string                                      = make([]string, 0, 0)
	StaticHandle func(http.ResponseWriter, *http.Request) bool = nil
	NotHandle    func(*Controller)                             = nil //默认处理器
)

//添加路由器
func HandlerRouter(method, url string, f func(*Controller)) {
	_, ok := routers[url]
	if ok {
		//路由已存在，抛出异常
		panic(errors.New(fmt.Sprintf("url:%sexist?", url)))
	}
	if f == nil {
		//函数不能为空，抛出异常
		panic(errors.New("func is nul?"))
	}
	var r Router
	r.Action = f
	r.Path = url
	r.Method = strings.ToUpper(method)
	routers[url] = &r
}

//静态资源目录检查
func checkStaticFolder(w http.ResponseWriter, r *http.Request, url string) bool {

	for i := 0; i < len(staticFolder); i++ {
		if isOk, _ := regexp.MatchString(staticFolder[i], url); isOk {
			//运行静态文件处理拦截器
			if StaticHandle != nil {
				b := StaticHandle(w, r)
				if !b {
					return isOk
				}
			}
			http.ServeFile(w, r, strings.TrimLeft(url, "/"))
			return isOk

		}
	}
	return false
}

//设置静态目录
func SetStaticFolder(folderName ...string) {
	staticFolder = append(staticFolder, folderName...)
}
func handle(w http.ResponseWriter, r *http.Request) {
	//匹配URL
	reg := regexp.MustCompile(`\?.*`)
	url := reg.ReplaceAllString(r.RequestURI, "")
	//检查是否为静态资源目录
	if checkStaticFolder(w, r, url) {
		return
	}
	val, ok := routers[url]

	if !ok {
		notFound(r, w)
		return
	}
	if val.Method != "" && strings.ToUpper(r.Method) != val.Method {
		notFound(r, w)
		return
	}

	c := newController(r, w)
	//拦截器调用
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			c.Write([]byte(fmt.Sprintf("%s", recoverErr)))
		}
	}()
	interceptorRun(c, val.Action)

}

//路由过滤器
func RouterFilter(c *Controller) {

}
func notFound(r *http.Request, w http.ResponseWriter) {
	if NotHandle != nil {
		NotHandle(newController(r, w))
	} else {

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Fount 404!"))
	}

}
