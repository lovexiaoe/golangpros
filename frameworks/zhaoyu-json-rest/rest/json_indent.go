package rest

import (
	"bufio"
	"encoding/json"
	"net"
	"net/http"
)

// JsonIndentMiddleware 提供缩进格式的json编码.
// 在开发中使用方便.
// 它通过中间件的responseWriter “子类“来实现 ，在程序为jsonIndentResponseWriter
// 替换了 writer.EncodeJson 和 writer.WriteJson 方法的实现,
type JsonIndentMiddleware struct {

	// prefix string, as in json.MarshalIndent
	Prefix string

	// indentation string, as in json.MarshalIndent
	Indent string
}

// MiddlewareFunc makes JsonIndentMiddleware implement the Middleware interface.
func (mw *JsonIndentMiddleware) MiddlewareFunc(handler HandlerFunc) HandlerFunc {

	if mw.Indent == "" {
		mw.Indent = "  "
	}

	return func(w ResponseWriter, r *Request) {

		writer := &jsonIndentResponseWriter{w, false, mw.Prefix, mw.Indent}
		// call the wrapped handler
		handler(writer, r)
	}
}

// Private responseWriter intantiated by the middleware.
// It implements the following interfaces:
// ResponseWriter
// http.ResponseWriter
// http.Flusher
// http.CloseNotifier
// http.Hijacker
type jsonIndentResponseWriter struct {
	ResponseWriter
	wroteHeader bool
	prefix      string
	indent      string
}

// 替换父级的EncodeJson方法，提供缩进功能。
func (w *jsonIndentResponseWriter) EncodeJson(v interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(v, w.prefix, w.indent)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Make sure the local EncodeJson and local Write are called.
// Does not call the parent WriteJson.
func (w *jsonIndentResponseWriter) WriteJson(v interface{}) error {
	b, err := w.EncodeJson(v)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// Call the parent WriteHeader.
func (w *jsonIndentResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.wroteHeader = true
}

// Make sure the local WriteHeader is called, and call the parent Flush.
// Provided in order to implement the http.Flusher interface.
func (w *jsonIndentResponseWriter) Flush() {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	flusher := w.ResponseWriter.(http.Flusher)
	flusher.Flush()
}

// Call the parent CloseNotify.
// Provided in order to implement the http.CloseNotifier interface.
func (w *jsonIndentResponseWriter) CloseNotify() <-chan bool {
	notifier := w.ResponseWriter.(http.CloseNotifier)
	return notifier.CloseNotify()
}

// Provided in order to implement the http.Hijacker interface.
func (w *jsonIndentResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := w.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}

// Make sure the local WriteHeader is called, and call the parent Write.
// Provided in order to implement the http.ResponseWriter interface.
func (w *jsonIndentResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	writer := w.ResponseWriter.(http.ResponseWriter)
	return writer.Write(b)
}
