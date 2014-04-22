// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fileutil

import (
	"os"
	"runtime"
	"strings"
)

func CreateFile(path string) (file *os.File, err error) {
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, `\`, "/", -1)
	}

	if i := strings.LastIndex(path, "/"); i == -1 {
		return os.Create(path)
	} else {
		err = os.MkdirAll(path[:i+1], 0777)
		if err != nil {
			return nil, err
		}
		return os.Create(path)
	}
}
