package main

import (
	"log"
	"mygolang/zhaoyu-json-rest/rest"
	"net/http"
)

func main() {
	api := rest.NewApi()
	statusMw := &rest.StatusMiddleware{}
	api.Use(statusMw)
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/.status", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(statusMw.GetStatus())
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
