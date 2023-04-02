package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	testBroker()
}

type handleFunc func(c *Consumer, msg string)

type Consumer struct {
	ch      chan string
	name    string
	handler handleFunc
}

func (c *Consumer) Start() {
	go func() {
		fmt.Printf("consumer:%s start\n", c.name)
		for {
			select {
			case msg, ok := <-c.ch:
				if !ok {
					fmt.Printf("consumer:%s return\n", c.name)
					return
				}
				c.handler(c, msg)

			}
		}

		fmt.Printf("consumer:%s end\n", c.name)
	}()
}

func (c *Consumer) close() {
	close(c.ch)
}

type Broker struct {
	consumers []*Consumer
}

func (b *Broker) Produce(msg string) {
	for _, consumer := range b.consumers {
		consumer.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumers = append(b.consumers, c)
}

func (b *Broker) Close() {
	for _, c := range b.consumers {
		c.close()
	}
}

func testBroker() {
	handler := func(c *Consumer, msg string) {
		fmt.Printf("consumer:%s get msg:%s\n", c.name, msg)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	b := &Broker{
		consumers: make([]*Consumer, 0, 4),
	}
	c1 := &Consumer{
		ch:      make(chan string, 1),
		name:    "A",
		handler: handler,
	}

	c2 := &Consumer{
		ch:      make(chan string, 1),
		name:    "B",
		handler: handler,
	}

	b.Subscribe(c1)
	b.Subscribe(c2)
	c1.Start()
	c2.Start()
	b.Produce("hello")
	b.Produce("world")
	<-ctx.Done()
	b.Close()
	time.Sleep(5 * time.Second)

}
