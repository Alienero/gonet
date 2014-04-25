// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"
)

func Run(addr string) {
	// TODO read the config file
	http.ListenAndServe(addr, Mux)
}

// func Add(match string, c ControllerInterface) {
// 	Mux.Add(match, c)
// }
// func AddGateway(match string, c ControllerInterface) {
// 	Mux.AddGateway(match, c)
// }
// func AddGroup(match string, group *MuxGroup, c ControllerInterface) {
// 	Mux.AddGroup(match, group, c)
// }
