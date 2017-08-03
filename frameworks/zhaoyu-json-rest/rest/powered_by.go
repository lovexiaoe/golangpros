package rest

const xPoweredByDefault = "go-json-rest"

// PoweredByMiddleware 向HTTP response中添加 "X-Powered-By" header。
type PoweredByMiddleware struct {

	// 如果指定它, 则作为response header中 "X-Powered-By" 的值。
	// 默认 "go-json-rest".
	XPoweredBy string
}

// PoweredByMiddleware 实现 Middleware 接口
func (mw *PoweredByMiddleware) MiddlewareFunc(h HandlerFunc) HandlerFunc {

	poweredBy := xPoweredByDefault
	if mw.XPoweredBy != "" {
		poweredBy = mw.XPoweredBy
	}

	return func(w ResponseWriter, r *Request) {

		w.Header().Add("X-Powered-By", poweredBy)

		// call the handler
		h(w, r)

	}
}
