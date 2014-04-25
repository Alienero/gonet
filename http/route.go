// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/astaxie/beego/session"
)

// A defaultmux
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
	// Gateways
	Group map[string]ControllerInterface
	// All routes
	Routes map[string]ControllerInterface
}

func NewDefaultMux() *DefaultMux {
	return &DefaultMux{
		Group:  make(map[string]ControllerInterface),
		Routes: make(map[string]ControllerInterface),
	}
}

// Implements the handler
func (mux *DefaultMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the session
	sess := Sessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	// Matching group routes
	for k, v := range mux.Group {
		if strings.HasPrefix(r.URL.Path, k) {
			method(v, w, r, sess)
			break
		}
	}

	// Matching routes
	for k, v := range mux.Routes {
		if k == r.URL.Path {
			method(v, w, r, sess)
			break
		}
	}
}

// Add a router
func (mux *DefaultMux) Add(match string, c ControllerInterface) {
	mux.Routes[match] = c
}

// Add a gateway
func (mux *DefaultMux) AddGateway(match string, c ControllerInterface) {
	mux.Group[match] = c
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
	mux.routes[match] = c
}

// Add a group(include his members). And if c is not nil adding the group's name to the gateway
func (mux *DefaultMux) AddGroup(match string, group *MuxGroup, c ControllerInterface) {
	if c != nil {
		mux.AddGateway(match, c)
	}
	for k, v := range group.routes {
		mux.Add(match+k, v)
	}
}

// New a ControllerInterface and Transfering the parameters
func method(v ControllerInterface, w http.ResponseWriter, r *http.Request, sess session.SessionStore) {
	c, ok := reflect.New(reflect.Indirect(reflect.ValueOf(v)).Type()).Interface().(ControllerInterface)
	if !ok {
		panic("controller is not ControllerInterface")
	}
	c.Set(w, r, sess)
	// Matching the method
	switch r.Method {
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
	}
}
