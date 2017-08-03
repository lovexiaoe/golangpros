//HTTP routing 的查找树实现
//
//该查找树的实现支持:param和*splat参数.Strings 一般被用于表示HTTP routing 的路径。
//该查找树也为每个路径维护一张HTTP方法和路由关联的地图。
//
//你一般不会直接用到这个包。
package trie

import (
	"errors"
	"fmt"
)

//路径分割，默认使用/和.进行分割，返回前后两个字符串。
func splitParam(remaining string) (string, string) {
	i := 0
	for len(remaining) > i && remaining[i] != '/' && remaining[i] != '.' {
		i++
	}
	return remaining[:i], remaining[i:]
}

//使用 / 进行分割。
func splitRelaxed(remaining string) (string, string) {
	i := 0
	for len(remaining) > i && remaining[i] != '/' {
		i++
	}
	return remaining[:i], remaining[i:]
}

type node struct {
	HttpMethodToRoute map[string]interface{}

	Children       map[string]*node
	ChildrenKeyLen int

	ParamChild *node
	ParamName  string

	RelaxedChild *node
	RelaxedName  string

	SplatChild *node
	SplatName  string
}

//
func (n *node) addRoute(httpMethod, pathExp string, route interface{}, usedParams []string) error {

	//在HttpMethodToRoute中添加路径和方法的关联
	if len(pathExp) == 0 {
		// end of the path, leaf node, update the map
		if n.HttpMethodToRoute == nil {
			n.HttpMethodToRoute = map[string]interface{}{
				httpMethod: route,
			}
			return nil
		} else {
			if n.HttpMethodToRoute[httpMethod] != nil {
				return errors.New("node.Route already set, duplicated path and method")
			}
			n.HttpMethodToRoute[httpMethod] = route
			return nil
		}
	}

	token := pathExp[0:1]
	remaining := pathExp[1:]
	var nextNode *node

	//处理冒号的参数，添加参数名称到usedParams中，节点paramChild赋新的node，paramName为参数名称，nextNode赋节点的paramChild
	if token[0] == ':' {
		var name string
		name, remaining = splitParam(remaining)

		// 检测参数是否已经存在
		for _, e := range usedParams {
			if e == name {
				return errors.New(
					fmt.Sprintf("A route can't have two placeholders with the same name: %s", name),
				)
			}
		}
		usedParams = append(usedParams, name)

		if n.ParamChild == nil {
			n.ParamChild = &node{}
			n.ParamName = name
		} else {
			if n.ParamName != name {
				return errors.New(
					fmt.Sprintf(
						"Routes sharing a common placeholder MUST name it consistently: %s != %s",
						n.ParamName,
						name,
					),
				)
			}
		}
		nextNode = n.ParamChild
	} else if token[0] == '#' {
		//处理#号的参数，添加参数名称到usedParams中，节点RelaxedChild赋值新的node，RelaxedName为参数名称，nextNode赋节点的RelaxedChild
		var name string
		name, remaining = splitRelaxed(remaining)

		// 检测参数是否已经存在
		for _, e := range usedParams {
			if e == name {
				return errors.New(
					fmt.Sprintf("A route can't have two placeholders with the same name: %s", name),
				)
			}
		}
		usedParams = append(usedParams, name)

		if n.RelaxedChild == nil {
			n.RelaxedChild = &node{}
			n.RelaxedName = name
		} else {
			if n.RelaxedName != name {
				return errors.New(
					fmt.Sprintf(
						"Routes sharing a common placeholder MUST name it consistently: %s != %s",
						n.RelaxedName,
						name,
					),
				)
			}
		}
		nextNode = n.RelaxedChild
	} else if token[0] == '*' {
		//处理*号的参数，给节点的SplatChild和SplatName赋值，nextNode赋节点的SplatChild
		// *splat case
		name := remaining
		remaining = ""

		// 检测参数是否已经存在
		for _, e := range usedParams {
			if e == name {
				return errors.New(
					fmt.Sprintf("A route can't have two placeholders with the same name: %s", name),
				)
			}
		}

		if n.SplatChild == nil {
			n.SplatChild = &node{}
			n.SplatName = name
		}
		nextNode = n.SplatChild
	} else {
		// 一般情况的处理
		if n.Children == nil {
			n.Children = map[string]*node{}
			n.ChildrenKeyLen = 1
		}
		if n.Children[token] == nil {
			n.Children[token] = &node{}
		}
		nextNode = n.Children[token]
	}

	//递归调用处理节点
	return nextNode.addRoute(httpMethod, remaining, route, usedParams)
}

//节点压缩，如果子节点不为空，则对子结点做压缩
func (n *node) compress() {
	// *splat branch
	if n.SplatChild != nil {
		n.SplatChild.compress()
	}
	// :param branch
	if n.ParamChild != nil {
		n.ParamChild.compress()
	}
	// #param branch
	if n.RelaxedChild != nil {
		n.RelaxedChild.compress()
	}
	// main branch
	if len(n.Children) == 0 {
		return
	}
	// compressable ?
	canCompress := true
	for _, node := range n.Children {
		if node.HttpMethodToRoute != nil || node.SplatChild != nil || node.ParamChild != nil || node.RelaxedChild != nil {
			canCompress = false
		}
	}
	// compress
	if canCompress {
		merged := map[string]*node{}
		for key, node := range n.Children {
			for gdKey, gdNode := range node.Children {
				mergedKey := key + gdKey
				merged[mergedKey] = gdNode
			}
		}
		n.Children = merged
		n.ChildrenKeyLen++
		n.compress()
		// continue
	} else {
		for _, node := range n.Children {
			node.compress()
		}
	}
}

//在格式打印的前面添加padding个空格，
func printFPadding(padding int, format string, a ...interface{}) {
	for i := 0; i < padding; i++ {
		fmt.Print("-")
	}
	fmt.Printf(format, a...)
}

// 节点的调试打印 ，按层级显示
func (n *node) printDebug(level int) {
	level++
	// *splat branch
	if n.SplatChild != nil {
		printFPadding(level, "*splat\n")
		n.SplatChild.printDebug(level)
	}
	// :param branch
	if n.ParamChild != nil {
		printFPadding(level, ":param\n")
		n.ParamChild.printDebug(level)
	}
	// #param branch
	if n.RelaxedChild != nil {
		printFPadding(level, "#relaxed\n")
		n.RelaxedChild.printDebug(level)
	}
	// main branch
	for key, node := range n.Children {
		printFPadding(level, "\"%s\"\n", key)
		node.printDebug(level)
	}
}

//自己添加的打印 HttpMethodToRoute map 的方法
func (n *node) printFHttpMethodToRoute(level int) {
	level++
	// *splat branch
	if n.SplatChild != nil {
		printFPadding(level, "\"%s\" -splat\n", n.HttpMethodToRoute)
		n.SplatChild.printFHttpMethodToRoute(level)
	}
	// :param branch
	if n.ParamChild != nil {
		printFPadding(level, "\"%s\" -param\n", n.HttpMethodToRoute)
		n.ParamChild.printFHttpMethodToRoute(level)
	}
	// #param branch
	if n.RelaxedChild != nil {
		printFPadding(level, "\"%s\" -relaxed\n", n.HttpMethodToRoute)
		n.RelaxedChild.printFHttpMethodToRoute(level)
	}
	// main branch
	for _, node := range n.Children {
		printFPadding(level, "\"%s\"\n", node.HttpMethodToRoute)
		node.printFHttpMethodToRoute(level)
	}
}

//---定义route递归查找的工具
type paramMatch struct {
	name  string
	value string
}

type findContext struct {
	paramStack []paramMatch
	matchFunc  func(httpMethod, path string, node *node)
}

//---

func newFindContext() *findContext {
	return &findContext{
		paramStack: []paramMatch{},
	}
}

//在参数堆中push一个元素。
func (fc *findContext) pushParams(name, value string) {
	fc.paramStack = append(
		fc.paramStack,
		paramMatch{name, value},
	)
}

//在参数堆中pop一个元素。
func (fc *findContext) popParams() {
	fc.paramStack = fc.paramStack[:len(fc.paramStack)-1]
}

//将参数堆转换为map。
func (fc *findContext) paramsAsMap() map[string]string {
	r := map[string]string{}
	for _, param := range fc.paramStack {
		if r[param.name] != "" {
			// this is checked at addRoute time, and should never happen.
			panic(fmt.Sprintf(
				"placeholder %s already found, placeholder names should be unique per route",
				param.name,
			))
		}
		r[param.name] = param.value
	}
	return r
}

type Match struct {
	// Same Route as in AddRoute
	Route interface{}
	// map of params matched for this result
	Params map[string]string
}

func (n *node) find(httpMethod, path string, context *findContext) {

	if n.HttpMethodToRoute != nil && path == "" {
		context.matchFunc(httpMethod, path, n)
	}

	if len(path) == 0 {
		return
	}

	// *splat branch
	if n.SplatChild != nil {
		context.pushParams(n.SplatName, path)
		n.SplatChild.find(httpMethod, "", context)
		context.popParams()
	}

	// :param branch
	if n.ParamChild != nil {
		value, remaining := splitParam(path)
		context.pushParams(n.ParamName, value)
		n.ParamChild.find(httpMethod, remaining, context)
		context.popParams()
	}

	// #param branch
	if n.RelaxedChild != nil {
		value, remaining := splitRelaxed(path)
		context.pushParams(n.RelaxedName, value)
		n.RelaxedChild.find(httpMethod, remaining, context)
		context.popParams()
	}

	// main branch
	length := n.ChildrenKeyLen
	if len(path) < length {
		return
	}
	token := path[0:length]
	remaining := path[length:]
	if n.Children[token] != nil {
		n.Children[token].find(httpMethod, remaining, context)
	}
}

type Trie struct {
	root *node
}

//初始化一个空节点的树
func New() *Trie {
	return &Trie{
		root: &node{},
	}
}

//给一个树添加一个路径节点
func (t *Trie) AddRoute(httpMethod, pathExp string, route interface{}) error {
	return t.root.addRoute(httpMethod, pathExp, route, []string{})
}

// 为了减少树的大小 ，必须在最后一次添加路径后压缩。
func (t *Trie) Compress() {
	t.root.compress()
}

// 树的打印方法。
func (t *Trie) PrintDebug() {
	fmt.Print("<trie>\n")
	t.root.printDebug(0)
	fmt.Print("</trie>\n")
}

//自己添加
// 树的HttpMethodToRoute打印方法。
func (t *Trie) PrintHttpMethodToRoute() {
	fmt.Print("===MethodToRoute=====\n")
	t.root.printFHttpMethodToRoute(0)
	fmt.Print("===MethodToRoute=====\n")
}

// 给出一个路径的http方法，返回所有的匹配路径
func (t *Trie) FindRoutes(httpMethod, path string) []*Match {
	context := newFindContext()
	matches := []*Match{}
	context.matchFunc = func(httpMethod, path string, node *node) {
		if node.HttpMethodToRoute[httpMethod] != nil {
			// path and method match, found a route !
			matches = append(
				matches,
				&Match{
					Route:  node.HttpMethodToRoute[httpMethod],
					Params: context.paramsAsMap(),
				},
			)
		}
	}
	t.root.find(httpMethod, path, context)
	return matches
}

// 和FindRoutes方法相同，但是返回一个bool值，用于指示路径是否是匹配的。在返回405时非常有用。
func (t *Trie) FindRoutesAndPathMatched(httpMethod, path string) ([]*Match, bool) {
	context := newFindContext()
	pathMatched := false
	matches := []*Match{}
	context.matchFunc = func(httpMethod, path string, node *node) {
		pathMatched = true
		if node.HttpMethodToRoute[httpMethod] != nil {
			// path and method match, found a route !
			matches = append(
				matches,
				&Match{
					Route:  node.HttpMethodToRoute[httpMethod],
					Params: context.paramsAsMap(),
				},
			)
		}
	}
	t.root.find(httpMethod, path, context)
	return matches, pathMatched
}

// 给出一个路径，不限定http方法，返回匹配的路径
func (t *Trie) FindRoutesForPath(path string) []*Match {
	context := newFindContext()
	matches := []*Match{}
	context.matchFunc = func(httpMethod, path string, node *node) {
		params := context.paramsAsMap()
		for _, route := range node.HttpMethodToRoute {
			matches = append(
				matches,
				&Match{
					Route:  route,
					Params: params,
				},
			)
		}
	}
	t.root.find("", path, context)
	return matches
}
