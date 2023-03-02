# Context

context 包的核心 API 有四个： 
- context.WithValue：设置键值对，并且返 回一个新的 context 实例 
- context.WithCancel 
- context.WithDeadline 
- context.WithTimeout： 三者都返回一个 可取消的 context 实例，和取消函数 

> 注意：
> context 实例是不可变的，每一次都是新创建


## Context接口
```go
type Context interface {
    // 返回 context 是否设置了超时时间以及超时的时间点
    // 如果没有设置超时，那么 ok 的值返回 false
	Deadline() (deadline time.Time, ok bool)
	
    // 每次调用都会返回相同的结果
    // 如果 context 被取消，这里会返回一个被关闭的 channel
    // 如果是一个不会被取消的 context，那么这里会返回 nil
   Done() <-chan struct{}

   // 返回 done() 的原因
   // 如果 Done() 对应的通道还没有关闭，这里返回 nil
   // 如果通道关闭了，这里会返回一个非 nil 的值：
   // - 若果是被取消掉的，那么这里返回 Canceled 错误
   // - 如果是超时了，那么这里返回 DeadlineExceeded 错误
   Err() error
   
   // 获取 context 中保存的 key 对应的 value，如果不存在则返回 nil
   Value(key any) any
}
```

- Deadline() 表示如果有截止时间的话，得返回对应 deadline 时间；如果没有，则 ok 的值为 false。

- Done() 表示关于 channel 的数据通信，而且它的数据类型是 struct{}，一个空结构体，因此在 Go 里都是直接通过 close channel 来进行通知的，不会涉及具体数据传输。

- Err() 返回的是一个错误 error，如果上面的 Done() 的 channel 没被 close，则 error 为 nil；如果 channel 已被 close，则 error 将会返回 close 的原因，比如超时或手动取消。

- Value() 则是用来存储具体数据的方法。


Context 接口并不需要我们实现，context 包里内置实现了 2 个， 代码中最开始都是以这两个内置的作为最顶层的 partent context，衍生出更多的子 Context。

- context.Background()
- context.TODO()


```go



var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	return background
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
	return todo
}

```

## Context With TimeOut

```go
func contextWithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	start := time.Now().Unix()
	<- ctx.Done()
	end := time.Now().Unix()
	fmt.Printf("total time = %d\n" , end -start)
	err := ctx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("context canceled")
	case context.DeadlineExceeded:
		fmt.Println("context timeout")
	default:
		fmt.Println("context done")
	}

}
```

## Context With Cancel

```go
func contextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(){
		fmt.Println("go routine started")
		<- ctx.Done()
		fmt.Println("go routine canceled")
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
	err := ctx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("context canceled")
	case context.DeadlineExceeded:
		fmt.Println("context timeout")
	default:
		fmt.Println("context done")
	}

}
```

## Context With Deadline

```go
func contextWithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
	defer cancel()
	start := time.Now().Unix()
	<- ctx.Done()
	end := time.Now().Unix()
	fmt.Printf("total time = %d\n" , end -start)
	err := ctx.Err()
	switch err {
	case context.Canceled:
		fmt.Println("context canceled")
	case context.DeadlineExceeded:
		fmt.Println("context timeout")
	default:
		fmt.Println("context done")
	}
}
```

## Context With Value

```go
func contextWithValue() {
	parentKey := "parent_key"
	parentValue := "parent_value"
	parentCtx := context.WithValue(context.Background(), parentKey, parentValue)
	sonKey := "son_key"
	sonValue := "son_value"

	sonCtx := context.WithValue(parentCtx, sonKey, sonValue)

	if parentCtx.Value(sonKey) == nil {
		fmt.Println("parent can not get son context")
	}

	if sonCtx.Value(parentKey) != nil {
		fmt.Println("son can get parent context")
	}
}
```

## Context使用样例

```go
func exampleContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2* time.Second)
	defer cancel()

	bsCh := make(chan struct{})
	go func() {
		processBusiness()
		bsCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("TimeoutExample context timeout")
	case <-bsCh:
		fmt.Println("TimeoutExample process business end")
	}
}

func processBusiness() {
	time.Sleep(1 * time.Second)
}
```

## Context超时控制

> 注意:
>- 子Context超时时间超过父Context无效

```go
func exampleContextTimeoutParentSon() {
	bg := context.Background()
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	//这里子Context超时时间超过父Context无效
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, 4*time.Second)
	defer cancel2()
	defer cancel1()
	go func() {
		// 一秒钟之后就会过期
		<-subCtx.Done()
		fmt.Printf("subctx timout")
	}()

	time.Sleep(2 * time.Second)
}
```

## 最佳实践
- context.Background 只应用在最高等级，作为所有派生 context 的根。
- context.TODO 应用在不确定要使用什么的地方，或者当前函数以后会更新以便使用 context。
- context 取消是建议性的，这些函数可能需要一些时间来清理和退出。
- context.Value 应该很少使用，它不应该被用来传递可选参数。这使得 API 隐式的并且可以引起错误。取而代之的是，这些值应该作为参数传递。
- 尽量不要将 context 存储在结构中，在函数中显式传递它们，最好是作为第一个参数。

