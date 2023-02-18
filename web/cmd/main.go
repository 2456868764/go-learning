package main

import (
	v2 "github.com/2456868764/go-learning/web/api/v2"
	"github.com/2456868764/go-learning/web/pkg/engine"
)

func main() {
	engine := engine.New()
	engine.GET("/headers", v2.GetHeaders)
	engine.GET("/ip", v2.GetIP)
	engine.GET("/user-agent", v2.GetUserAgent)
	engine.Run(":8080")
}
