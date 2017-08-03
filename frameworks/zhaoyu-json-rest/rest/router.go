package rest

import (
	"errors"
	"mygolang/zhaoyu-json-rest/rest/trie"
	"net/http"
	"net/url"
	"strings"
)

type router struct {
	Routes                 []*Route
	disableTrieCompression bool
	index                  map[*Route]int
	trie                   *trie.Trie
}

// MakeRouter 返回 router app. 给出一个Route的集合。按照Routers相关的顺序。
// 它将请求分发给第一个匹配到的route 的 HandlerFunc。it dispatches the request to the
func MakeRouter(routes ...*Route) (App, error) {
	r := &router{
		Routes: routes,
	}
	err := r.start()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// 处理REST路由，运行用户代码
// Handle the REST routing and run the user code.
func (rt *router) AppFunc() HandlerFunc {
	return func(writer ResponseWriter, request *Request) {

		// find the route
		route, params, pathMatched := rt.findRouteFromURL(request.Method, request.URL)
		if route == nil {

			if pathMatched {
				// no route found, but path was matched: 405 Method Not Allowed
				Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// no route found, the path was not matched: 404 Not Found
			NotFound(writer, request)
			return
		}

		// a route was found, set the PathParams
		request.PathParams = params

		// run the user code
		handler := route.Func
		handler(writer, request)
	}
}

// 每一个新request会调用，性能是非常重要的。
func escapedPath(urlObj *url.URL) string {
	// 使用？分割为两个字符串。
	parts := strings.SplitN(urlObj.RequestURI(), "?", 2)
	return parts[0]
}

var preEscape = strings.NewReplacer("*", "__SPLAT_PLACEHOLDER__", "#", "__RELAXED_PLACEHOLDER__")

var postEscape = strings.NewReplacer("__SPLAT_PLACEHOLDER__", "*", "__RELAXED_PLACEHOLDER__", "#")

// pathExp url转义。
// 这个方法仅在初始化时调用。
func escapedPathExp(pathExp string) (string, error) {

	// PathExp validation
	if pathExp == "" {
		return "", errors.New("empty PathExp")
	}
	if pathExp[0] != '/' {
		return "", errors.New("PathExp must start with /")
	}
	if strings.Contains(pathExp, "?") {
		return "", errors.New("PathExp must not contain the query string")
	}

	// Get the right escaping
	// XXX a bit hacky

	pathExp = preEscape.Replace(pathExp)

	urlObj, err := url.Parse(pathExp)
	if err != nil {
		return "", err
	}

	// get the same escaping as find requests
	pathExp = urlObj.RequestURI()

	pathExp = postEscape.Replace(pathExp)

	return pathExp, nil
}

// 这个方法校验Routes 并准备对应的trie数据结构。
// 它需要在Routes被定义之后，并在查找Routers之前调用。
// 根据排序，如果有多个Route匹配，则第一个将被使用。
func (rt *router) start() error {

	rt.trie = trie.New()
	rt.index = map[*Route]int{}

	for i, route := range rt.Routes {

		// PathExp url编码转义.
		pathExp, err := escapedPathExp(route.PathExp)
		if err != nil {
			return err
		}

		// insert in the Trie
		err = rt.trie.AddRoute(
			strings.ToUpper(route.HttpMethod), // work with the HttpMethod in uppercase
			pathExp,
			route,
		)
		if err != nil {
			return err
		}

		// index
		rt.index[route] = i
	}

	if rt.disableTrieCompression == false {
		rt.trie.Compress()
	}

	return nil
}

// 根据matches中的routeIndex，返回最早定义的route。
func (rt *router) ofFirstDefinedRoute(matches []*trie.Match) *trie.Match {
	minIndex := -1
	var bestMatch *trie.Match

	for _, result := range matches {
		route := result.Route.(*Route)
		routeIndex := rt.index[route]
		if minIndex == -1 || routeIndex < minIndex {
			minIndex = routeIndex
			bestMatch = result
		}
	}

	return bestMatch
}

// 根据一个URL对象及参数，返回第一个匹配的route。
// Return the first matching Route and the corresponding parameters for a given URL object.
func (rt *router) findRouteFromURL(httpMethod string, urlObj *url.URL) (*Route, map[string]string, bool) {

	// lookup the routes in the Trie
	matches, pathMatched := rt.trie.FindRoutesAndPathMatched(
		strings.ToUpper(httpMethod), // work with the httpMethod in uppercase
		escapedPath(urlObj),         // work with the path urlencoded
	)

	// short cuts
	if len(matches) == 0 {
		// no route found
		return nil, nil, pathMatched
	}

	if len(matches) == 1 {
		// one route found
		return matches[0].Route.(*Route), matches[0].Params, pathMatched
	}

	// multiple routes found, pick the first defined
	result := rt.ofFirstDefinedRoute(matches)
	return result.Route.(*Route), result.Params, pathMatched
}

// 根据url字符串,返回第一个匹配的route及参数。
func (rt *router) findRoute(httpMethod, urlStr string) (*Route, map[string]string, bool, error) {

	// parse the url
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		return nil, nil, false, err
	}

	route, params, pathMatched := rt.findRouteFromURL(httpMethod, urlObj)
	return route, params, pathMatched, nil
}
