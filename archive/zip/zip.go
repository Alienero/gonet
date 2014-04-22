// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gzip implements reading and writing ZIP archives.
package zip

import (
	"archive/zip"
	"io"
	"os"

	mystack "github.com/Alienero/gonet/container/stack"
	"github.com/Alienero/gonet/fileutil"
)

// The method implements compress src file to ZIP archive named dst
func CompressFile(dst, src string) error {
	stack := mystack.NewStack()
	f_dst, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f_dst.Close()
	write := zip.NewWriter(f_dst)
	defer write.Close()
	stack.Push("")
	var temp string
	for {
		// for the stack
		if stack.Len() > 0 {
			parent := stack.Pop().(string)
			d, err := os.Open(src + "/" + parent)
			if err != nil {
				return err
			}
			// read all file of the directory
			files, _ := d.Readdir(0)
			for _, file := range files {
				if file.IsDir() {
					// push the directoy into stacktemp := ""
					if temp = ""; parent != "" {
						temp = parent + "/"
					}
					// creat the directory file head
					_, err = write.Create(temp + file.Name() + "/")
					if err != nil {
						break
					}
					stack.Push(temp + file.Name())
				} else {
					// Compress the file
					if temp = ""; parent != "" {
						temp = parent + "/"
					}
					var f io.Writer
					f, err = write.Create(temp + file.Name())
					if err != nil {
						break
					}
					var srcFile *os.File
					srcFile, err = os.Open(src + "/" + parent + "/" + file.Name())
					if err != nil {
						break
					}

					_, err = io.Copy(f, srcFile)
					srcFile.Close()
					if err != nil {
						break
					}
				}
			}
			d.Close()
			if err != nil {
				return err
			}
		} else {
			break
		}
	}
	return nil
}

// The method implements decompress src ZIP archive to file named dst
func DeCompressFile(dst, src string) error {
	err := os.MkdirAll(dst, 0777)
	if err != nil {
		return err
	}
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		// Creat the file
		temp := dst + "/" + f.Name
		if rs := []rune(temp); rs[len(rs)-1] == '/' {
			err = os.Mkdir(temp, 0777)
			if err != nil {
				return err
			}
			continue
		}
		file, err := fileutil.CreateFile(temp)
		if err != nil {
			return err
		}
		_, err = io.Copy(file, rc)
		file.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
