// 演示如何给代码写一个使用样例，通过godoc生成文档后会显示在文档中.

//描述了对 /sendjson endpoint的调用例子。
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// 样例是基于存在的函数或者方法，必须以Example+函数名命名.
//output标签定义了程序期待的输出，会显示在godoc的OutPut一栏，
//test测试框架会比较测试的结果，如果匹配，则通过，如果不匹配，则测试失败。
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(u)
	// Output:
	// {Bill bill@ardanstudios.com}
}
