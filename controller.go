package panda

import (
	"encoding/json"
	"fmt"
	. "github.com/sczhaoyu/panda/session"
	"html/template"
	"net/http"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//控制器
type Controller struct {
	*http.Request                              //请求信息
	http.ResponseWriter                        //输出信息
	Data                map[string]interface{} //渲染时候的参数
	Tpl                 string                 //渲染使用的模板
	SessionManager      *Manager               //sesion管理器
}

//创建控制器
func newController(r *http.Request, w http.ResponseWriter) *Controller {
	var c Controller
	c.Request = r
	c.ResponseWriter = w
	c.Data = make(map[string]interface{}, 0)
	c.Request.ParseForm()
	//加入session管理器
	if SessionSwitch && panda.SessionManager != nil {
		c.SessionManager = panda.SessionManager
	}
	return &c
}

//获取Session
func (c *Controller) GetSession(key string) interface{} {
	sess, err := c.SessionManager.SessionStart(c.ResponseWriter, c.Request)
	defer sess.SessionRelease(c.ResponseWriter)
	if err != nil {
		return nil
	}
	return sess.Get(key)
}

//设置SESSION
func (c *Controller) SetSession(key string, val interface{}) error {
	sess, err := c.SessionManager.SessionStart(c.ResponseWriter, c.Request)
	defer sess.SessionRelease(c.ResponseWriter)
	err = sess.Set(key, val)
	if err != nil {
		return err
	}
	return nil
}

//删除SESSION
func (c *Controller) DeleteSession(key string) error {
	sess, err := c.SessionManager.SessionStart(c.ResponseWriter, c.Request)
	defer sess.SessionRelease(c.ResponseWriter)
	err = sess.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

//销毁用户的SESSION
func (c *Controller) DestroySession() {
	c.SessionManager.SessionDestroy(c.ResponseWriter, c.Request)
}
func (c *Controller) NotFound() {
	notFound(c.ResponseWriter)
}
func (c *Controller) Write(b []byte) {
	c.ResponseWriter.Header().Add("content-type", "text/html;charset=utf-8")
	c.ResponseWriter.Write(b)
}
func (c *Controller) WriteJson(ret interface{}) {
	c.ResponseWriter.Header().Add("access-control-allow-origin", "*")
	c.ResponseWriter.Header().Add("content-type", "application/json;charset=utf-8")
	c.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With")
	data, _ := json.Marshal(ret)
	c.ResponseWriter.Write(data)
}

//渲染模板
func (c *Controller) render() {
	c.Tpl = ViewPath + c.Tpl
	ret := []string{c.Tpl}
	ret = append(ret, templates[c.Tpl]...)
	t, err := template.New(filepath.Base(c.Tpl)).Funcs(PandaTplFuncMap).ParseFiles(ret...)
	if err != nil {
		c.Write([]byte(err.Error()))
		return
	}
	c.ResponseWriter.Header().Add("content-type", "text/html;charset=utf-8")
	err = t.Execute(c.ResponseWriter, c.Data)
	if err != nil {
		ERROR.Println(err)
	}
}

//将参数转化为一个结构
func (c *Controller) ParseForm(obj interface{}) error {
	form := c.Request.Form
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !isStructPtr(objT) {
		return fmt.Errorf("%v must be  a struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()

	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		fieldT := objT.Field(i)
		tags := strings.Split(fieldT.Tag.Get("form"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			continue
		} else {
			tag = tags[0]
		}
		value := form.Get(tag)
		if len(value) == 0 {
			continue
		}

		switch fieldT.Type.Kind() {
		case reflect.Bool:
			if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
				fieldV.SetBool(true)
				continue
			}
			if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
				fieldV.SetBool(false)
				continue
			}
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			fieldV.SetString(value)
		case reflect.Struct:
			switch fieldT.Type.String() {
			case "time.Time":
				format := time.RFC3339
				if len(tags) > 1 {
					format = tags[1]
				}
				t, err := time.Parse(format, value)
				if err != nil {
					return err
				}
				fieldV.Set(reflect.ValueOf(t))
			}
		case reflect.Slice:
			if fieldT.Type == sliceOfInts {
				formVals := form[tag]
				fieldV.Set(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(int(1))), len(formVals), len(formVals)))
				for i := 0; i < len(formVals); i++ {
					val, err := strconv.Atoi(formVals[i])
					if err != nil {
						return err
					}
					fieldV.Index(i).SetInt(int64(val))
				}
			} else if fieldT.Type == sliceOfStrings {
				formVals := form[tag]
				fieldV.Set(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf("")), len(formVals), len(formVals)))
				for i := 0; i < len(formVals); i++ {
					fieldV.Index(i).SetString(formVals[i])
				}
			}
		}
	}
	return nil
}
