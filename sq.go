package gomq

import (
	"sync"
)

//new varible mq
func New()Mq{
	return Mq{}
}
type Mq struct{
	Lock sync.Mutex
	// Previous *Mq
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
	s.Lock.Lock()
	r.Data=v
	r.Next=&Mq{}
	s.Lock.Unlock()
	return v
}
func (s * Mq)Out()interface{}{
	s.Lock.Lock()
	r:=s.Data
	if r==nil{
		s.Lock.Unlock()				
		return nil
	}
	l:=s.Next.Lock
	s.Lock.Unlock()			
	*s=*s.Next
	s.Lock=l
	return r
}
func (s * Mq)Last()*Mq{
	var m *Mq=s	
	if m.Data==nil{
		return m
	}
	s.Lock.Lock()
	for{
		if m.Next==nil{
			s.Lock.Unlock()			
			return m
		}	
		m=m.Next
	}
}
func (s * Mq)Clear(){
	for{
		if nil==s.Out(){
			return
		}
	}	
}
