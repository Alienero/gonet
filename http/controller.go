// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"

	"github.com/astaxie/beego/session"
)

// TODO session & form quick

type Controller struct {
	// A connection context
	Ctx *Context
}

type ControllerInterface interface {
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()

	Set(w http.ResponseWriter, r *http.Request)
	Prepare() int
	finished()
}

func NewController(w http.ResponseWriter, r *http.Request, sess session.SessionStore) *Controller {
	return &Controller{
		Ctx: NewContext(w, r, sess),
	}
}
func (c *Controller) Set(w http.ResponseWriter, r *http.Request) {
	c.Ctx = &Context{
		ResponseWriter: w,
		Request:        r,
	}
}
func (c *Controller) Prepare() int {
	return 0
}
func (c *Controller) finished() {
	c.Ctx.finished()
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
