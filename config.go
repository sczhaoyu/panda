package panda

//Config说明:
//一个JSON字符串,传入Manager的配置信息
//cookieName: 客户端存储 cookie 的名字。
//enableSetCookie,omitempty: 是否开启SetCookie,omitempty这个设置
//gclifetime: 触发 GC 的时间。
//maxLifetime: 服务器端存储的数据的过期时间
//secure: 是否开启 HTTPS，在 cookie 设置的时候有 cookie.Secure 设置。
//sessionIDHashFunc: sessionID 生产的函数，默认是 sha1 算法。
//sessionIDHashKey: hash 算法中的 key。
//cookieLifeTime: 客户端存储的 cookie 的时间，默认值是 0，即浏览器生命周期。
//providerConfig: 配置信息，根据不同的引擎设置不同的配置信息，详细的配置请看下面的引擎设置

//globalSessions 有多个函数如下所示：

//SessionStart 根据当前请求返回 session 对象
//SessionDestroy 销毁当前 session 对象
//SessionRegenerateId 重新生成 sessionID
//GetActiveSession 获取当前活跃的 session 用户
//SetHashFunc 设置 sessionID 生成的函数
//SetSecure 设置是否开启 cookie 的 Secure 设置
var (
	SessionSwitch = false //是否启用SESSION
	SessionType   string  //存储方式 memory、file、mysql 或 redis。
	SessionConfig string  //Session配置
)
var (
	HttpSSL      bool   //是否启动SSL
	LocalAddress string //监听地址和端口
	HttpSslCert  string //SSL信息
	HttpSslKey   string //SSL信息
)

const (
	SESSION_MEMCACHE  = "memcache"
	SESSION_COUCHBASE = "couchbase"
	SESSION_MYSQL     = "mysql"
	SESSION_REDIS     = "redis"
	SESSION_POSTGRES  = "postgres"
	SESSION_MEMORY    = "memory"
)
