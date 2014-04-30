// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"io"
	"net/http"

	"github.com/astaxie/beego/session"
)

type Context struct {
	sess session.SessionStore

	Writer io.Writer

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

// New a Context
func NewContext(w http.ResponseWriter, r *http.Request, sess session.SessionStore) *Context {
	return &Context{
		sess:           sess,
		ResponseWriter: w,
		Request:        r,
	}
}
func (cxt *Context) WriteString(s string) (int, error) {
	return cxt.ResponseWriter.Write([]byte(s))
}
func (cxt *Context) GetForm(key string) string {
	if cxt.Request.Form == nil {
		cxt.Request.ParseForm()
	}
	return cxt.Request.FormValue(key)
}
func (cxt *Context) GetSess() session.SessionStore {
	if cxt.sess == nil {
		// Get the session
		cxt.sess = Sessions.SessionStart(cxt.ResponseWriter, cxt.Request)
	}
	return cxt.sess
}

// Redirect does redirection to localurl with http header status code.
// It sends http response header directly.
func (ctx *Context) Redirect(status int, localurl string) {
	ctx.ResponseWriter.Header().Add("Location", localurl)
	ctx.ResponseWriter.WriteHeader(status)
}

func (cxt *Context) finished() {
	if cxt.sess != nil {
		cxt.sess.SessionRelease(cxt.ResponseWriter)
	}
}

type Writer struct {
	w      io.Writer
	encode string
}

func NewWriter(w io.Writer, encode string) *Writer {
	return &Writer{w, encode}
}
func (this *Writer) Write(data []byte) (int, error) {
	// TODO encode
	return this.w.Write(data)
}

type Reader struct {
	body   io.ReadCloser
	encode string
}

func NewReader(body io.ReadCloser, encode string) *Reader {
	return &Reader{body, encode}
}
func (this *Reader) Read(data []byte) (int, error) {
	// TODO encode
	return this.Read(data)
}
