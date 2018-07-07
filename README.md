# gomq 简单好用的消息队列和栈

## 功能介绍
- 1.链表(消息队列) 
- 2.栈 
## 链表 使用方法
- `执行`  go get github.com/bugfan/gomq 
- `新建` test:=gomq.NewMQ()
- `入队列` test.In(Data)
- `长度` test.Len()
- `出队列` test.Out()
- `清空队列` test.Clear()
- go语言实现的单向链表内置了入队列和出队列等方法，与go自带的容器不同,可以在请求并发或者来不及发送(接收)的时候入此队列 最后依次取出,缓存聊天消息记录等

## 栈 使用方法
- `执行`  go get github.com/bugfan/gomq 
- `新建` stack:=gomq.NewStack()
- `压栈` stack.Push(Data)
- `出栈` stack.Pop()
- `长度` stack.Size()
- `清空栈` stack.Clear()


## 测试 cd .../gomq 执行 go test
