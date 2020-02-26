# 简易协程池

# 使用方法 
```go
p := pool.New(16, count) //需要总数为count的，大小限制为16的协程池
p.Run(func(){}) //开始任务


p.Wait() //等待所有协程执行完毕
```