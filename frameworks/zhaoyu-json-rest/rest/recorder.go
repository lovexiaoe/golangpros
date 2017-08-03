package rest

import (
	"bufio"
	"net"
	"net/http"
)

// RecorderMiddleware 保存了一个response的http 状态码和写入byte多少的记录
// 可以通过 handler的request.Env["STATUS_CODE"].(int)
// 和request.Env["BYTES_WRITTEN"].(int64)取得。
type RecorderMiddleware struct{}

// RecorderMiddleware 实现 Middleware 接口.
func (mw *RecorderMiddleware) MiddlewareFunc(h HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, r *Request) {

		writer := &recorderResponseWriter{w, 0, false, 0}

		// call the handler
		h(writer, r)

		r.Env["STATUS_CODE"] = writer.statusCode
		r.Env["BYTES_WRITTEN"] = writer.bytesWritten
	}
}

// responseWriter 被recorder 中间件实例化
// 记录了response HTTP 的状态码。
// 实现了以下接口:
// ResponseWriter
// http.ResponseWriter
// http.Flusher
// http.CloseNotifier
// http.Hijacker
type recorderResponseWriter struct {
	ResponseWriter
	statusCode   int
	wroteHeader  bool
	bytesWritten int64
}

// 记录 状态码。
func (w *recorderResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.statusCode = code
	w.wroteHeader = true
}

// write Json.
func (w *recorderResponseWriter) WriteJson(v interface{}) error {
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

// 确保 local WriteHeader 被调用, 并调用父级的 Flush.
// 实现 http.Flusher 接口.
func (w *recorderResponseWriter) Flush() {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	flusher := w.ResponseWriter.(http.Flusher)
	flusher.Flush()
}

// 实现 http.CloseNotifier 接口。
func (w *recorderResponseWriter) CloseNotify() <-chan bool {
	notifier := w.ResponseWriter.(http.CloseNotifier)
	return notifier.CloseNotify()
}

// 实现 http.Hijacker 接口.
func (w *recorderResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := w.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}

// 实现 http.ResponseWriter 接口.
func (w *recorderResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	writer := w.ResponseWriter.(http.ResponseWriter)
	written, err := writer.Write(b)
	w.bytesWritten += int64(written)
	return written, err
}
