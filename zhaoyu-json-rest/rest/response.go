package rest

import (
	"bufio"
	"encoding/json"
	"net"
	"net/http"
)

// ResponseWriter 接口专用来处理 JSON HTTP 响应。
// 注意，被框架实例化的 responseWriter对应也实现了 可以通过类型断言的其它许多接口
// 如：http.ResponseWriter, http.Flusher, http.CloseNotifier, http.Hijacker.
type ResponseWriter interface {

	//和 http.ResponseWriter 接口中的方法相同
	Header() http.Header

	// 使用 EncodeJson 生成内容, 如果内容还没有写完，在headers中写http.StatusOK。然后写内容。
	// Content-Type 被设置成 "application/json"。
	WriteJson(v interface{}) error

	// 把数据结构编译成JSON, 主要用于在中间件中封装 ResponseWriter
	EncodeJson(v interface{}) ([]byte, error)

	// Similar to the http.ResponseWriter interface, with additional JSON related
	// headers set.
	WriteHeader(int)
}

//定义了错误名称。可以在启动server之前被更改。
var ErrorFieldName = "Error"

// 在response中生成一个json的错误信息，
func Error(w ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	err := w.WriteJson(map[string]string{ErrorFieldName: error})
	if err != nil {
		panic(err)
	}
}

//生成一个404，没有找到的错误
func NotFound(w ResponseWriter, r *Request) {
	Error(w, "Resource not found", http.StatusNotFound)
}

// 在资源处理器中被实例化的私有responseWriter。
// 它实现了下面接口:
// ResponseWriter
// http.ResponseWriter
// http.Flusher
// http.CloseNotifier
// http.Hijacker
type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (w *responseWriter) WriteHeader(code int) {
	if w.Header().Get("Content-Type") == "" {
		// Per spec, UTF-8 is the default, and the charset parameter should not
		// be necessary. But some clients (eg: Chrome) think otherwise.
		// Since json.Marshal produces UTF-8, setting the charset parameter is a
		// safe option.
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	w.ResponseWriter.WriteHeader(code)
	w.wroteHeader = true
}

func (w *responseWriter) EncodeJson(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Encode the object in JSON and call Write.
func (w *responseWriter) WriteJson(v interface{}) error {
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

// Provided in order to implement the http.ResponseWriter interface.
func (w *responseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

// Provided in order to implement the http.Flusher interface.
func (w *responseWriter) Flush() {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	flusher := w.ResponseWriter.(http.Flusher)
	flusher.Flush()
}

// Provided in order to implement the http.CloseNotifier interface.
func (w *responseWriter) CloseNotify() <-chan bool {
	notifier := w.ResponseWriter.(http.CloseNotifier)
	return notifier.CloseNotify()
}

// Provided in order to implement the http.Hijacker interface.
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker := w.ResponseWriter.(http.Hijacker)
	return hijacker.Hijack()
}
