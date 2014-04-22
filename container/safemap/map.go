// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package safemap

import (
	"sync"
)

type SMap struct {
	lock    *sync.RWMutex
	element map[interface{}]interface{}
}

func NewSMap() *SMap {
	return &SMap{new(sync.RWMutex), make(map[interface{}]interface{})}
}
func (s *SMap) Get(key interface{}) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.element[key]
}
func (s *SMap) Set(key interface{}, vaule interface{}) {
	s.lock.Lock()
	s.element[key] = vaule
	s.lock.Unlock()
}
func (s *SMap) Delete(key interface{}) {
	s.lock.Lock()
	delete(s.element, key)
	s.lock.Unlock()
}
