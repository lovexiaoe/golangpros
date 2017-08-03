package rest

import (
	"net/http"
)

//HandlerFunc定义了一个处理函数，是在go-json-rest中http.HandlerFunc的等效方法
type HandlerFunc func(ResponseWriter, *Request)

//实现了该接口的对象会在框架栈中被用作一个app。
//这个app是框架栈的栈顶元素，框架栈中的其它元素是middleware
type App interface {
	AppFunc() HandlerFunc
}

//对HandleFunc的一个封装
//AppSimple是一个适配器类型，可以使用一个简单的方法更容易的写出一个app。
//如:rest.NewApi(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) { ... }))
type AppSimple HandlerFunc

//AppFunc 使 AppSimple 实现 App 接口更简单.
func (as AppSimple) AppFunc() HandlerFunc {
	//将as转换成HandlerFunc，返回。
	return HandlerFunc(as)
}

//定义了一个中间件接口，一个对象实现该接口并封装一个HandlerFunc，就可以在中间件栈中使用了。
type Middleware interface {
	MiddlewareFunc(handler HandlerFunc) HandlerFunc
}

// MiddlewareSimple 是一个适配器类型， 便于使用一个简单的方法定义一个中间件。
// 如: api.Use(rest.MiddlewareSimple(func(h HandlerFunc) Handlerfunc { ... }))
type MiddlewareSimple func(handler HandlerFunc) HandlerFunc

// MiddlewareSimple 实现 Middleware 接口。
func (ms MiddlewareSimple) MiddlewareFunc(handler HandlerFunc) HandlerFunc {
	return ms(handler)
}

//在一组中间件中倒序调用MiddlewareFunc方法，返回一个多层嵌套的HandlerFunc方法用于执行。
// This can be used to wrap a set of middlewares, post routing, on a per Route basis.
func WrapMiddlewares(middlewares []Middleware, handler HandlerFunc) HandlerFunc {
	wrapped := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i].MiddlewareFunc(wrapped)
	}
	return wrapped
}

// 处理 net/http 和 go-json-rest 对象之间的转换。
// 它实例化了 rest.Request and rest.ResponseWriter, ...
func adapterFunc(handler HandlerFunc) http.HandlerFunc {

	return func(origWriter http.ResponseWriter, origRequest *http.Request) {

		// instantiate the rest objects
		request := &Request{
			origRequest,
			nil,
			map[string]interface{}{},
		}

		writer := &responseWriter{
			origWriter,
			false,
		}

		// call the wrapped handler
		handler(writer, request)
	}
}
