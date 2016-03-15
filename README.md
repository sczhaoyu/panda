 

<pre>
    <code>  
package main

import (
    "github.com/sczhaoyu/panda"
)

func index(c *panda.Controller) {
    c.SetSession("hello", "world") //设置session
    c.Write([]byte("hello world"))
}
func interceptor(c *panda.Controller) bool {
    c.Write([]byte("hello interceptor"))
    return true
}
func interceptorAfter(c *panda.Controller) bool {
    c.Write([]byte("hello interceptorAfter"))
    return true
}
func main() {
    panda.LocalAddress = ":8080"
    panda.HandlerRouter(panda.CASUAL, "/", index)
    panda.HandlerRouter(panda.POST, "/post", index)
    panda.HandlerRouter(panda.GET, "/get", index)
    panda.AddInterceptor(interceptor, panda.BEFORE)
    panda.AddInterceptor(interceptor, panda.AFTER)
    panda.SessionSwitch = true //session开启
    panda.Run()
}
    </code>
</pre>
