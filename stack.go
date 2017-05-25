package main

// This should probably conform to FILO http://en.cppreference.com/w/cpp/container/stack

import (
	"sync"
	"errors"
)

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s []GameState
}

func NewStack() *stack {
	return &stack {sync.Mutex{}, make([]GameState,0), }
}

func (s stack) Peek() GameState {
	if s.Len() == 0 {
		return nil
	}

	return s.s[s.Len()-1]
}

func (s stack) Len() int {
	return len(s.s)
}

func (s *stack) Push(v GameState) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() (GameState, error) {
	s.lock.Lock()
	defer s.lock.Unlock()


	l := len(s.s)
	if l == 0 {
		return nil, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}