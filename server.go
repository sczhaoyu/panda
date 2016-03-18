package panda

import (
	"fmt"
	"github.com/agtorre/gocolorize"
	. "github.com/sczhaoyu/panda/session"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//服务器
type Panda struct {
	Port           int          //启动端口
	Server         *http.Server //http服务
	HttpSSL        bool         //是否启动SSL
	NetWork        string       //工作模式
	LocalAddress   string       //监听地址和端口
	HttpSslCert    string       //SSL信息
	HttpSslKey     string       //SSL信息
	SessionManager *Manager     //Session管理
}

var (
	panda Panda //管理全局
	//错误日志颜色样式
	colors = map[string]gocolorize.Colorize{
		"error": gocolorize.NewColor("red:black"),
	}
	error_log = pandaLogs{c: colors["error"], w: os.Stderr}
	ERROR     = log.New(&error_log, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
)

func Run() {
	//tcp模式
	panda.NetWork = "tcp"
	panda.init()
	panda.Server = &http.Server{
		Addr:         panda.LocalAddress,
		Handler:      http.HandlerFunc(handle),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Start port:%d[%s]\n", panda.Port, time.Now().Local().Format("2006-01-02 15:04:05"))
	}()
	if panda.HttpSSL {
		if err := panda.Server.ListenAndServeTLS(panda.HttpSslCert, panda.HttpSslKey); err != nil {
			ERROR.Println("Start SSL error:", err)
		}

	} else {
		listen, err := net.Listen(panda.NetWork, panda.LocalAddress)
		if err != nil {
			ERROR.Println("listen server error:", err)
		}
		if err = panda.Server.Serve(listen); err != nil {
			ERROR.Println("server start error:", err)
		}
	}

}

//初始化服务地址
func (p *Panda) init() {
	//启动地址检测
	if LocalAddress == "" {
		p.LocalAddress = ":5200"
	} else {
		p.LocalAddress = LocalAddress
	}
	panda.HttpSSL = HttpSSL
	panda.HttpSslCert = HttpSslCert
	panda.HttpSslKey = HttpSslKey
	parts := strings.SplitN(p.LocalAddress, ":", 2)
	p.LocalAddress = parts[0] + ":" + parts[1]
	port, err := strconv.Atoi(parts[1])

	if err != nil {
		panic(fmt.Sprintf("port is error %v", parts[1]))
	}
	p.Port = port
	//session
	if SessionSwitch {
		//SESSION开启，未发现配置信息，加入默认的内存配置
		if SessionType == "" {
			SessionType = SESSION_MEMORY
			SessionConfig = `{"cookieName":"GOSESSIONID","gclifetime":3600}`
		}
		sess, err := NewManager(SessionType, SessionConfig)
		if err != nil {
			panic(err)
		}
		go sess.GC()
		panda.SessionManager = sess

	}
	//session end
}

type pandaLogs struct {
	c gocolorize.Colorize
	w io.Writer
}

func (p *pandaLogs) Write(b []byte) (n int, err error) {
	return p.w.Write([]byte(p.c.Paint(string(b))))
}
