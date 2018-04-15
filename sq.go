package goq

import (
	"log"
	"sync"
)

//new varible mq
func New()Mq{
	return Mq{}
}
type Mq struct{
	Lock sync.Mutex
	Previous *Mq
	Data interface{}
	Next *Mq
}
func (s * Mq)Len()int64{
	var i int64
	var m *Mq=s
	for{
		if m.Next==nil{
			return i
		}
		i++
		m=m.Next
	}
}
func (s * Mq)In(v interface{})interface{}{
	if v==nil{
		return nil
	}
	r:=s.Last()
	n:=&Mq{
		Previous:r,
		Data:v,
		Next:nil,
	}
	s.Lock.Lock()
	r=n
	s.Lock.Unlock()
	return v
}
func (s * Mq)Out()interface{}{
	s.Lock.Lock()
	defer s.Lock.Unlock()
	log.Println("b:",s.Next.Data,s.Previous.Data)		
	if s.Data==nil{
		return nil
	}
	v:=s.Data
	s=s.Next
	return v
}
func (s * Mq)Last()*Mq{
	s.Lock.Lock()
	var m *Mq=s
	for{
		if m.Next==nil{
			s.Lock.Unlock()			
			return m
		}	
		m=m.Next
	}
}