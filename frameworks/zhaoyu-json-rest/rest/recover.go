package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

// RecoverMiddleware 捕获在执行的HandleFunc 中发生的panic 错误。
// 并将它们转成 500 错误码，response返回。
type RecoverMiddleware struct {

	// 自定义logger记录 panic 错误,
	// 可选项, 默认的日志为 log.New(os.Stderr, "", 0)
	Logger *log.Logger

	// json格式标记。
	EnableLogAsJson bool

	//如果为true,当一个panic发生时，将在500的response body中打印错误信息和堆栈。
	EnableResponseStackTrace bool
}

// RecoverMiddleware 实现 Middleware 接口。
func (mw *RecoverMiddleware) MiddlewareFunc(h HandlerFunc) HandlerFunc {

	// set the default Logger
	if mw.Logger == nil {
		mw.Logger = log.New(os.Stderr, "", 0)
	}

	return func(w ResponseWriter, r *Request) {

		// catch user code's panic, and convert to http response
		defer func() {
			if reco := recover(); reco != nil {
				trace := debug.Stack()

				// log the trace
				message := fmt.Sprintf("%s\n%s", reco, trace)
				mw.logError(message)

				// write error response
				if mw.EnableResponseStackTrace {
					Error(w, message, http.StatusInternalServerError)
				} else {
					Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}
		}()

		// call the handler
		h(w, r)
	}
}

func (mw *RecoverMiddleware) logError(message string) {
	if mw.EnableLogAsJson {
		record := map[string]string{
			"error": message,
		}
		b, err := json.Marshal(&record)
		if err != nil {
			panic(err)
		}
		mw.Logger.Printf("%s", b)
	} else {
		mw.Logger.Print(message)
	}
}
