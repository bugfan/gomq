package gomq

import (
	"errors"
	"sync"
)

func NewStack() Stack {
	return Stack{}
}

type Stack struct {
	sync.Mutex
	Data []interface{}
}

func (s *Stack) Push(v interface{}) error {
	if v == nil {
		return errors.New("push data is nil!")
	}
	s.Lock()
	s.Data = append(s.Data, v)
	s.Unlock()
	return nil
}
func (s *Stack) Pop() interface{} {
	if len(s.Data) < 1 {
		return nil
	} else {
		s.Lock()
		data := s.Data[len(s.Data)-1]
		s.Data = append(s.Data[:len(s.Data)-1], s.Data[len(s.Data):]...)
		s.Unlock()
		return data
	}
}
func (s *Stack) Size() int {
	s.Lock()
	size := len(s.Data)
	s.Unlock()
	return size
}
func (s *Stack) Clear() {
	for {
		if s.Pop() == nil {
			return
		}
	}
}
