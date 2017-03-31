package rest

import (
	"mime"
	"net/http"
	"strings"
)

// ContentTypeCheckerMiddleware 校验  request header 的 Content-Type，如果类型不正确
// 则返回一个 StatusUnsupportedMediaType (415)的 HTTP 返回码
// 如果Content-Typer不为空，那么应该是 'application/json'
// Note: 如果 charset parameter 存在， 必须为 UTF-8.
type ContentTypeCheckerMiddleware struct{}

// MiddlewareFunc makes ContentTypeCheckerMiddleware implement the Middleware interface.
func (mw *ContentTypeCheckerMiddleware) MiddlewareFunc(handler HandlerFunc) HandlerFunc {

	return func(w ResponseWriter, r *Request) {

		mediatype, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		charset, ok := params["charset"]
		if !ok {
			charset = "UTF-8"
		}

		// per net/http doc, means that the length is known and non-null
		if r.ContentLength > 0 &&
			!(mediatype == "application/json" && strings.ToUpper(charset) == "UTF-8") {

			Error(w,
				"Bad Content-Type or charset, expected 'application/json'",
				http.StatusUnsupportedMediaType,
			)
			return
		}

		// call the wrapped handler
		handler(w, r)
	}
}
