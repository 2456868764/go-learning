package main

import (
	v3 "github.com/2456868764/go-learning/web/api/v3"
	"github.com/2456868764/go-learning/web/pkg/engine"
)

func main() {
	engine := engine.New()
	engine.GET("/headers", v3.GetHeaders)
	engine.GET("/ip", v3.GetIP)
	engine.GET("/user-agent", v3.GetUserAgent)
	engine.GET("/user/:userId/profile", v3.GetUserProfile)
	engine.Run(":8080")
}
