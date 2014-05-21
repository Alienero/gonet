package main

import (
	"net/http"

	myhttp "github.com/Alienero/gonet/http"
)

func main() {
	myhttp.Mux.Add("/hello/", &indexPage{})
	myhttp.Run(":8081")
}

type indexPage struct {
	myhttp.Controller
}

func (c *indexPage) Get() {
	handle := http.StripPrefix("/hello/", http.FileServer(http.Dir(`D:\git\helloJs`)))
	handle.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}
