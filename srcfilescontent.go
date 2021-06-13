package main

// set Content of Go-Sourcefiles

var routergo string = `package main


func SetRouterPaths() {

	//Set your Router-Paths here and link them with a handler-func from handlers.go

	router.GET("/", Start)

}`

var handlersgo string = `package main

import "net/http"

//Define your Handler-Funcs like the Start-Func below

func Start(w http.ResponseWriter, r *http.Request) {
	
}`

var maingo string = `package main

import (
	"log"
	"net/http"

	"github.com/bouk/httprouter"
)

//Set the port the app will listen to

var port string = "8000"

/////////////////////////////////////

var router = httprouter.New()

func main() {

	SetRouterPaths()

	log.Fatal(http.ListenAndServe(":"+port, router))

}
`
