// 快速安装 RESTful JSON API 的方法
//
// 地址：http://ant0ine.github.io/go-json-rest/
//
// Go-Json-Rest 是一个构建在net/http之上的轻量级应用，它可以很容易地构建 RESTful JSON API
// 它使用树的实现提供了快速和扩展性强的 request 路由。帮助处理json 请求和响应。并提供了一些功能性
// 中间件，如CORS, Auth, Gzip,Status 等等。
//
// 示例:
//
//      package main
//
//      import (
//              "github.com/ant0ine/go-json-rest/rest"
//              "log"
//              "net/http"
//      )
//
//      type User struct {
//              Id   string
//              Name string
//      }
//
//      func GetUser(w rest.ResponseWriter, req *rest.Request) {
//              user := User{
//                      Id:   req.PathParam("id"),
//                      Name: "Antoine",
//              }
//              w.WriteJson(&user)
//      }
//
//      func main() {
//              api := rest.NewApi()
//              api.Use(rest.DefaultDevStack...)
//              router, err := rest.MakeRouter(
//                      rest.Get("/users/:id", GetUser),
//              )
//              if err != nil {
//                      log.Fatal(err)
//              }
//              api.SetApp(router)
//              log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
//      }
//
//
package rest
