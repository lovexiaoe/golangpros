package rest

import (
	"time"
)

// TimerMiddleware 计算出某handler执行花费的时间。
//  handler可以通过request.Env["ELAPSED_TIME"].(*time.Duration)取得结果。
// 或者通过request.Env["START_TIME"].(*time.Time)也可。
type TimerMiddleware struct{}

// 实现中间件接口，计算中间件执行的时间
func (mw *TimerMiddleware) MiddlewareFunc(h HandlerFunc) HandlerFunc {
	return func(w ResponseWriter, r *Request) {

		start := time.Now()
		r.Env["START_TIME"] = &start

		// call the handler
		h(w, r)

		end := time.Now()
		elapsed := end.Sub(start)
		r.Env["ELAPSED_TIME"] = &elapsed
	}
}
