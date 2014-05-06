// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/astaxie/beego/session"
)

const (
	// Go on exe
	PASS = iota //b=0
	// Stop
	STOP //c=1
)

// A defaultmux user the go net ctrollor
var Mux *DefaultMux

func init() {
	Mux = NewDefaultMux()
	InitDefaultSessions()
}

// Just a beego session manager
var Sessions *session.Manager

// Initialization the session manager
func InitDefaultSessions() {
	var err error
	Sessions, err = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	if err != nil {
		println(err.Error())
	}
	go Sessions.GC() // new a goroutinue to clean the session whitch timeout
}

type DefaultMux struct {
	// All routes
	Routes     map[string]ControllerInterface
	FileRoutes map[string]ControllerInterface
}

func NewDefaultMux() *DefaultMux {
	return &DefaultMux{
		Routes:     make(map[string]ControllerInterface),
		FileRoutes: make(map[string]ControllerInterface),
	}
}

// Implements the handler
func (mux *DefaultMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Matching routes
	// 性能优化
	v := mux.Routes[r.URL.Path]
	if v == nil {
		// 精确匹配失败，进行模糊匹配
		// find the static routes
		if !strings.HasSuffix(r.URL.Path, "/") {
			for k, tv := range mux.FileRoutes {
				if strings.HasPrefix(r.URL.Path, k) {
					v = tv
					println("Match :", k)
					break
				}
			}
		}
	}
	if v != nil {
		fmt.Println("Routes:", "path->", r.URL.Path)
		ctx := NewContext(w, r)
		r.Body = ctx.In
		method(v, ctx)
	} else {
		http.Error(w, "Not found", 404)
	}
}

// New a ControllerInterface and Transfering the parameters
func method(v ControllerInterface, ctx *Context) {
	// Get a controller
	c, ok := reflect.New(reflect.Indirect(reflect.ValueOf(v)).Type()).Interface().(ControllerInterface)
	if !ok {
		panic("Controller is not ControllerInterface")
	}
	c.Set(ctx)
	defer c.finished()
	// Do the Prepare
	if c.Prepare() == STOP {
		return
	}
	// Matching the method
	switch ctx.Request.Method {
	// Call ControllerInterface's method,not Controller's method
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
	default:
		// TODO use the user custom method name
	}
}

// Add a router (warnning you can't use match("/") as the route)
func (mux *DefaultMux) Add(match string, c ControllerInterface) {
	// TODO: index
	// fix code style like golang
	println("add:", match)
	switch {
	case match == "/":
		mux.Routes[match] = c
	case strings.HasSuffix(match, "/"):
		mux.FileRoutes[match] = c
	default:
		mux.Routes[match] = c
	}
}

// A Group e.g. "/admin/login", "/admin/info" all belong to "/admin" group
type MuxGroup struct {
	routes map[string]ControllerInterface
}

func NewMuxGroup() *MuxGroup {
	return &MuxGroup{make(map[string]ControllerInterface)}
}

// Add the group's member
func (mux *MuxGroup) Add(match string, c ControllerInterface) {
	if strings.HasSuffix(match, "/") {
		panic(`connot use "/" in the end of match`)
	}
	mux.routes[match] = c
}

// Add a group(include his members).
func (mux *DefaultMux) AddGroup(match string, group *MuxGroup) {
	for k, v := range group.routes {
		if match == "/" {
			mux.Add(k, v)
		} else {
			mux.Add(match+k, v)
		}
	}
}
