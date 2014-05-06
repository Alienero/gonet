// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package htto implements a simple and userful http framwork
package http

// TODO add some config for the http
type http_config struct {
	// e.g. "https" or "http"
	Schem string

	IsGzip bool
}
