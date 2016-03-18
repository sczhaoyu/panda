package panda

const (
	BEFORE = iota //函数运行前
	AFTER         //函数运行后
	PANIC         //异常拦截器
)

type interceptor struct {
	F    func(*Controller) bool //运行的函数
	When int                    //运行前或者运行后执行（0,1）
}

var interceptors []interceptor = make([]interceptor, 0, 0)

//函数运行前
func interceptorRun(c *Controller, f func(*Controller)) {
	b := interceptorFilert(c, BEFORE)
	if b {
		f(c)
		//判断是否运行渲染
		if len(c.tpl) > 0 {
			c.render()
		}
	}
	//函数运行后
	defer interceptorFilert(c, AFTER)
}

func interceptorFilert(c *Controller, when int) bool {
	for i := 0; i < len(interceptors); i++ {
		if interceptors[i].When == when {
			b := interceptors[i].F(c)
			if !b {
				return false
			}
		}
	}
	return true
}
func AddInterceptor(f func(*Controller) bool, when int) {
	var i interceptor
	i.F = f
	i.When = when
	interceptors = append(interceptors, i)
}
