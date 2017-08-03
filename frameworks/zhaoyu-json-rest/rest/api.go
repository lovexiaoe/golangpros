package rest

import (
	"net/http"
)

//api定义了一个中间件的stack和app
type Api struct {
	stack []Middleware
	app   App
}

// NewApi 实例化一个新的Api对象。 the Middleware stack is empty, and the App is nil.
func NewApi() *Api {
	return &Api{
		stack: []Middleware{},
		app:   nil,
	}
}

//use方法，向栈中压入多个中间件
func (api *Api) Use(middlewares ...Middleware) {
	api.stack = append(api.stack, middlewares...)
}

// SetApp sets the App in the Api object.
func (api *Api) SetApp(app App) {
	api.app = app
}

// MakeHandler 包装了所有栈中的 Middlewares 和 App , 返回一个可调用的http.Handler方法。
// 如果栈中的 Middleware 和 App 都是空的, 则HandlerFunc 方法什么都不做.
func (api *Api) MakeHandler() http.Handler {
	var appFunc HandlerFunc
	if api.app != nil {
		appFunc = api.app.AppFunc()
	} else {
		appFunc = func(w ResponseWriter, r *Request) {}
	}
	return http.HandlerFunc(
		adapterFunc(
			WrapMiddlewares(api.stack, appFunc),
		),
	)
}

// 定义一个方便开发用的中间件栈. 其中包括:
// console friendly logging, JSON indentation, error stack strace in the response.
var DefaultDevStack = []Middleware{
	&AccessLogApacheMiddleware{},
	&TimerMiddleware{},
	&RecorderMiddleware{},
	&PoweredByMiddleware{},
	&RecoverMiddleware{
		EnableResponseStackTrace: true,
	},
	&JsonIndentMiddleware{},
	&ContentTypeCheckerMiddleware{},
}

// 定义一个方便正式产品用的中间件栈。 其中包括:
// Apache CombinedLogFormat logging, gzip compression.
var DefaultProdStack = []Middleware{
	&AccessLogApacheMiddleware{
		Format: CombinedLogFormat,
	},
	&TimerMiddleware{},
	&RecorderMiddleware{},
	&PoweredByMiddleware{},
	&RecoverMiddleware{},
	&GzipMiddleware{},
	&ContentTypeCheckerMiddleware{},
}

// 定义一个默认的中间件栈。
var DefaultCommonStack = []Middleware{
	&TimerMiddleware{},
	&RecorderMiddleware{},
	&PoweredByMiddleware{},
	&RecoverMiddleware{},
}
