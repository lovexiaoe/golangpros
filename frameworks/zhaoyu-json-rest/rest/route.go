package rest

import (
	"strings"
)

// Route定义了一个被 router使用的route。它可以被直接实例化，或者以以下快捷方法使用：
// rest.Get, rest.Post, rest.Put, rest.Patch and rest.Delete.
type Route struct {

	// 任意的 HTTP 方法，它会被大写来避免一些错误。
	HttpMethod string

	// 一个类似 "/resource/:id.json" 的字符串。
	// Placeholders 支持如下:
	// :paramName 匹配任意以 '/' 或者 '.' 开头的字符
	// #paramName 匹配任意以 '/' 开头的字符
	// *paramName 匹配任意以“定义字符串”结尾的字符串。
	// (placeholder 的名称必须在每个 PathExp 中唯一)
	PathExp string

	// 当 route 被获取时，方法将会被执行
	Func HandlerFunc
}

// MakePath 根据Route 和 pathParams 参数生成对应的 path
// 为了在 反向路径解析中使用。
func (route *Route) MakePath(pathParams map[string]string) string {
	path := route.PathExp
	for paramName, paramValue := range pathParams {
		paramPlaceholder := ":" + paramName
		relaxedPlaceholder := "#" + paramName
		splatPlaceholder := "*" + paramName
		r := strings.NewReplacer(paramPlaceholder, paramValue, splatPlaceholder, paramValue, relaxedPlaceholder, paramValue)
		path = r.Replace(path)
	}
	return path
}

// 返回HEAD方法的route
// Head 是实例化一个 HEAD route 的快捷方法。
// 查看Route对象的参数定义，你会发现它等同于 &Route{"HEAD", pathExp, handlerFunc}
func Head(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "HEAD",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}

// 返回Get方法的route
// Get is a shortcut method that instantiates a GET route. See the Route object the parameters definitions.
// Equivalent to &Route{"GET", pathExp, handlerFunc}
func Get(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "GET",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}

// 返回Post方法的route
// Post is a shortcut method that instantiates a POST route. See the Route object the parameters definitions.
// Equivalent to &Route{"POST", pathExp, handlerFunc}
func Post(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "POST",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}

// 返回Put方法的route
// Put is a shortcut method that instantiates a PUT route.  See the Route object the parameters definitions.
// Equivalent to &Route{"PUT", pathExp, handlerFunc}
func Put(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "PUT",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}

// 返回Patch方法的route
// Patch is a shortcut method that instantiates a PATCH route.  See the Route object the parameters definitions.
// Equivalent to &Route{"PATCH", pathExp, handlerFunc}
func Patch(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "PATCH",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}

// 返回Delete方法的route
// Delete is a shortcut method that instantiates a DELETE route. Equivalent to &Route{"DELETE", pathExp, handlerFunc}
func Delete(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "DELETE",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}

// 返回Options方法的route
// Options is a shortcut method that instantiates an OPTIONS route.  See the Route object the parameters definitions.
// Equivalent to &Route{"OPTIONS", pathExp, handlerFunc}
func Options(pathExp string, handlerFunc HandlerFunc) *Route {
	return &Route{
		HttpMethod: "OPTIONS",
		PathExp:    pathExp,
		Func:       handlerFunc,
	}
}
