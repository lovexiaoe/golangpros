/*
ForceSSL
Demonstrate how to use the ForceSSL Middleware to force HTTPS on requests to a go-json-rest API.

For the purposes of this demo, we are using HTTP for all requests and checking the X-Forwarded-Proto header to see if it is set to HTTPS (many routers set this to show what type of connection the client is using, such as Heroku). To do a true HTTPS test, make sure and use http.ListenAndServeTLS with a valid certificate and key file.

Additional documentation for the ForceSSL middleware can be found here.

curl demo:

curl -i 127.0.0.1:8080/
curl -H "X-Forwarded-Proto:https" -i 127.0.0.1:8080/
*/

package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jadengore/go-json-rest-middleware-force-ssl"
)

func main() {
	api := rest.NewApi()
	api.Use(&forceSSL.Middleware{
		TrustXFPHeader:     true,
		Enable301Redirects: false,
		Message:            "We are unable to process your request over HTTP.",
	})
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"body": "Hello World!"})
	}))

	// For the purposes of this demo, only HTTP connections accepted.
	// For true HTTPS, use ListenAndServeTLS.
	// https://golang.org/pkg/net/http/#ListenAndServeTLS
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
