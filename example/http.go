package main

import (
	myhttp "github.com/Alienero/gonet/http"
)

func main() {
	group := myhttp.NewMuxGroup()
	group.Add("/hello", &hello{})
	myhttp.Mux.AddGroup("/login", group)
	myhttp.Run(":8080")
}

type hello struct {
	myhttp.Controller
}

func (c *hello) Get() {
	c.Ctx.WriteString("hello http框架")
}
func (c *hello) Prepare() int {
	if c.Ctx.GetSess().Get("name") != "hello" {
		// 请登录
		c.Ctx.WriteString("请登录")
		if name := c.Ctx.GetForm("name"); name == "hello" {
			c.Ctx.GetSess().Set("name", "hello")
			c.Ctx.WriteString("登录成功")
		} else {
			c.Ctx.WriteString("用户名错误")
		}
		return myhttp.STOP
	}
	return myhttp.PASS

}
