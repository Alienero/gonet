// Copyright © 2014 Alienero. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gzip implements a simple stack
package stack

import "sync"

type Stack struct {
	element []interface{}
}

// It's not thread safe
func NewStack() *Stack {
	return &Stack{element: make([]interface{}, 0)}
}
func (s *Stack) Len() int { return len(s.element) }
func (s *Stack) Pop() interface{} {
	if s.Len() < 1 {
		return nil
	}
	old := s.element
	n := len(old)
	x := old[n-1]
	s.element = old[0 : n-1]
	return x
}
func (s *Stack) Push(x interface{}) {
	s.element = append(s.element, x)
}
func (s *Stack) Peek() interface{} {
	l := s.Len()
	if l > 0 {
		return s.element[l-1]
	}
	return nil
}

// Thread safe implements stack
type SafeStack struct {
	element []interface{}
	lock    *sync.RWMutex
}

func NewSafeStack() *SafeStack {
	return &SafeStack{make([]interface{}, 0), new(sync.RWMutex)}
}
func (s *SafeStack) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.element)
}
func (s *SafeStack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.element) < 1 {
		return nil
	}
	old := s.element
	n := len(old)
	x := old[n-1]
	s.element = old[0 : n-1]
	return x
}
func (s *SafeStack) Push(x interface{}) {
	s.lock.Lock()
	s.element = append(s.element, x)
	s.lock.Unlock()
}
func (s *SafeStack) Peek() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	l := len(s.element)
	if l > 0 {
		return s.element[l-1]
	}
	return nil
}
