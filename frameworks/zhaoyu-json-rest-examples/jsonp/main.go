package main

import (
	"log"
	"mygolang/zhaoyu-json-rest/rest"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.JsonpMiddleware{
		CallbackNameKey: "cb",
	})
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
