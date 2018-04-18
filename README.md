# gomq 简单好用的go消息队列

# 介绍
    # go语言实现的单向链表内置了入队列和出队列方法，与go自带的容器不同，可以在请求并发或者来不及发送(接收)的时候入此队列 最后在依次取出

# 使用方法
    # 执行  go get github.com/bugfan/gomq 
    # 新建 test:=gomq.New()
    # 入队列 test.In(Data)
    # 长度 test.Len()
    # 出队列 test.Out()
        
    # 测试
      cd .../gomq
      执行 go test
