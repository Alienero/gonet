// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"

	"github.com/astaxie/beego/session"
)

type Controller struct {
	Ctx *Context
}
type Context struct {
	Sess           session.SessionStore
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type ControllerInterface interface {
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Set(w http.ResponseWriter, r *http.Request, sess session.SessionStore)
	MatchMethod()
}

func NewController(w http.ResponseWriter, r *http.Request, sess session.SessionStore) *Controller {
	return &Controller{
		Ctx: &Context{
			Sess:           sess,
			ResponseWriter: w,
			Request:        r,
		},
	}
}
func (c *Controller) Set(w http.ResponseWriter, r *http.Request, sess session.SessionStore) {
	c.Ctx = &Context{
		Sess:           sess,
		ResponseWriter: w,
		Request:        r,
	}
}
func (c *Controller) MatchMethod() {
	switch c.Ctx.Request.Method {
	case "GET":
		c.Get()
	case "POST":
		c.Post()
	case "DELETE":
		c.Delete()
	case "PUT":
		c.Put()
	case "PATCH":
		c.Patch()
	case "OPTIONS":
		c.Options()
	}
}
func (c *Controller) Get() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

// Post adds a request function to handle POST request.
func (c *Controller) Post() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

// Delete adds a request function to handle DELETE request.
func (c *Controller) Delete() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

// Put adds a request function to handle PUT request.
func (c *Controller) Put() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

// Head adds a request function to handle HEAD request.
func (c *Controller) Head() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

// Patch adds a request function to handle PATCH request.
func (c *Controller) Patch() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

// Options adds a request function to handle OPTIONS request.
func (c *Controller) Options() {
	http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
}
