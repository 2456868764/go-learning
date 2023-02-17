package main

import (
	v1 "github.com/2456868764/go-learning/web/api/v1"
	"github.com/2456868764/go-learning/web/pkg/engine"
)

func main() {
	engine := engine.New()
	engine.GET("/headers", v1.GetHeaders)
	engine.GET("/ip", v1.GetIP)
	engine.GET("/user-agent", v1.GetUserAgent)
	engine.Run(":8080")
}
