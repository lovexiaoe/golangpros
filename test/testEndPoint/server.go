// 提供一个简单的web service.
package main

import (
	"log"
	"net/http"

	"github.com/lovexiaoe/golangpros/test/testEndPoint/handlers"
)

// main is the entry point for the application.
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
