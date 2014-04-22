// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"crypto/sha512"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestZip(t *testing.T) {
	CompressFile("D:/test1.zip", `C:\Users\Yi\Desktop\jieyasuo`)
	DeCompressFile("D:/Test", `D:/test1.zip`)
	CompressFile("D:/test.zip", "D:/Test")

	sha1st, err := fileSha512(`D:/test1.zip`)
	if err != nil {
		t.Error(err)
	}
	sha2nd, err := fileSha512("D:/test.zip")
	if err != nil {
		t.Error(err)
	}
	// check the both
	if sha1st != sha2nd {
		t.Errorf("Failed! 1:%s 2:%s", sha1st, sha2nd)
	}
	os.Remove("D:/test1.zip")
	os.Remove("D:/test.zip")
	os.RemoveAll("D:Test/")
}

func fileSha512(path string) (string, error) {
	h := sha512.New()
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	io.Copy(h, f)
	f.Close()
	return strings.Replace(fmt.Sprintf("% x", h.Sum(nil)), " ", "", -1), nil
}
