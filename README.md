gonet
=====
### gonet is a simple and useful go network framework
 It's support high-performance http server,include a  proxy server and a http framework
 And providing a high-performance socket framework

Quick Start
===========
##http server
```golang
package main

import (
	myhttp "github.com/Alienero/gonet/http"
)

func main() {
	group := myhttp.NewMuxGroup()
	group.Add("/hello", &hello{})
	myhttp.Mux.AddGroup("/good", group, &index{})
	myhttp.Run(":80")
}

type hello struct {
	myhttp.Controller
}

func (c *hello) Get() {
	c.Ctx.WriteString("hello ")
}

type index struct {
	myhttp.Controller
}

func (c *index) Get() {
	c.Ctx.WriteString("曼曼猪")
}
```
