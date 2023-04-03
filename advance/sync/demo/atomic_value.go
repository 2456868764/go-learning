package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	// 从数据库或者文件系统中读取配置信息，然后以map的形式存放在内存里

	config := make(map[string]string)
	config["rand"] = fmt.Sprintf("%d", rand.Intn(100))
	fmt.Printf("load config: %+v\n", config)
	return config
}

func main() {
	// config变量用来存放该服务的配置信息
	var config atomic.Value
	// 初始化时从别的地方加载配置文件，并存到config变量里
	config.Store(loadConfig())
	time.Sleep(2 * time.Second)

	go func() {
		// 每10秒钟定时的拉取最新的配置信息，并且更新到config变量里
		for {
			// 对应于赋值操作 config = loadConfig()
			config.Store(loadConfig())
			time.Sleep(1 * time.Second)
		}
	}()

	ch := make(chan int)
	// 创建工作线程，每个工作线程都会根据它所读取到的最新的配置信息来处理请求
	for i := 0; i < 2; i++ {
		go func() {
			for r := range ch {
				// 对应于取值操作 c := config
				// 由于Load()返回的是一个interface{}类型，所以我们要先强制转换一下
				c := config.Load().(map[string]string)
				fmt.Printf("get config:%+v\n", c)
				time.Sleep(1 * time.Second)
				// 这里是根据配置信息处理请求的逻辑...
				_, _ = r, c
			}
		}()
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}
	time.Sleep(10 * time.Second)
}
