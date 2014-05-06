package main

import (
	myhttp "github.com/Alienero/gonet/http"
)

func main() {
	myhttp.Mux.Add("/hello/", &hello{})
	myhttp.Run(":8081")
}

type hello struct {
	myhttp.Controller
}

func (c *hello) Get() {
	c.Ctx.WriteString("hello http框架")
}
