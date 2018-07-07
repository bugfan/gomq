package gomq

import (
	"log"
	"testing"
)

func TestSq(t *testing.T) {
	// 测试 mq.go
	log.Println("zxy")
	test := NewMQ()
	log.Println(test.In(12))
	log.Println(test.In(123))
	log.Println(test.Len())
	log.Println(test.Out())
	log.Println(test.Len())
	log.Println(test.Out())
	log.Println(test.In(121))
	log.Println(test.Out())
	log.Println(test.Len())
	log.Println(test.Out())
	log.Println(test.Out())
	log.Println(test.In("avvv"))
	log.Println(test.In("121"))
	test.Clear()
	log.Println(test.Out())
	log.Println(test.Out())
	log.Println(test.Out())
	log.Println(test.Out())

	// 测试 Stack
	log.Println("\n 测试 Stack")
	stack := NewStack()
	stack.Push(456)
	stack.Push("在西方")
	s1 := stack.Pop()
	log.Println("stack s1:", s1)
	s1 = stack.Pop()
	log.Println("stack s2", s1)
	log.Println("stack len:", stack.Size())
	s1 = stack.Pop()
	log.Println("stack s3", s1)

}
