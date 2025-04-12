# GoLearning
Golang学习记录

---

## 链接
[GO八小时入门](https://www.bilibili.com/video/BV1gf4y1r79E?spm_id_from=333.788.player.switch&vd_source=452811d53d64d58829c7c9b100c1115c&p=51)



---

### 切片
```go
type slice struct {
    array unsafe.Pointer  // 指向底层数组的指针
    len   int             // 当前长度
    cap   int             // 容量
}
```
>切片当作参数传入时确实可以类似于引用传递，但是**当切片扩容时**，系统会分配一个新的更大的底层数组，这是内部维护的指针会指向新的数组，这时对切片的修改将不会影响传入的切片


### 协程
为什么要引入协程
> 为了解决多进程多线程的高内存高消耗问题

**模式**： 用户态协程：内核态线程 -> `1:1`,`N:1`,`M:N`
**调度器设计策略： ![[截屏2025-04-10 16.21.40.png]]
- 复用线程
	- 偷取(work stealing)
	- 分离(hand off)
- 利用并行
- 抢占
- 全局Goroutine队列

## tips
**循环闭包内协程的变量快照问题**
```go
for _, f := range filenames {
    go func() {
        thumbnail.ImageFile(f) // NOTE: incorrect!
        // ...
    }()
}
```
> 上面这个单独的变量f是被所有的匿名函数值所共享，且会被连续的循环迭代所更新的。当新的goroutine开始执行字面函数时，for循环可能已经更新了f并且开始了另一轮的迭代或者（更有可能的）已经结束了整个循环,导致数据异常

解决方法
> 匿名函数传参临时保存变量
