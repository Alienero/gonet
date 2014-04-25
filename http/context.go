// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"
	"sync"

	"github.com/astaxie/beego/session"
)

type Context struct {
	Sess           session.SessionStore
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	once           *sync.Once
}

// New a Context
func NewContext(w http.ResponseWriter, r *http.Request, sess session.SessionStore) *Context {
	return &Context{
		Sess:           sess,
		ResponseWriter: w,
		Request:        r,
		once:           new(sync.Once),
	}
}
func (cxt *Context) WriteString(s string) (int, error) {
	return cxt.ResponseWriter.Write([]byte(s))
}
func (cxt *Context) GetForm(key string) string {
	cxt.once.Do(func() { cxt.Request.ParseForm() })
	return cxt.Request.FormValue(key)
}
