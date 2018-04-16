# gosq 简单好用的go消息队列

#介绍
    #1.go语言实现的单向链表内置了入队列和出队列方法，与go自带的容器不同，可以在请求并发或者来不及发送(接收)的时候入此队列 最后在依次取出

#使用方法
    #1. 方法一
        go get github.com/bugfan/gosq 
        新建 test:=gosq.New()
        入队列 test.In(Data)
        长度 test.Len()
        出队列 test.Out()
    #2. 方法二
        cd .../gosq
        执行 go test
