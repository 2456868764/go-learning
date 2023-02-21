package v3

import (
	"fmt"
	"github.com/2456868764/go-learning/web/pkg/engine"
	"strconv"
)

func GetUserAgent(c *engine.Context) {
	ip := c.GetHeader("User-Agent")
	c.StringOk(fmt.Sprintf("User-Agent=%s\n", ip))
}

func GetIP(c *engine.Context) {
	ip := c.GetHeader("REMOTE-ADDR")
	c.StringOk(fmt.Sprintf("IP=%s\n", ip))
}

func GetHeaders(c *engine.Context) {
	content := ""
	for k, v := range c.R.Header {
		content += fmt.Sprintf("%s=%s\n", k, v)
	}
	c.StringOk(content)
}

func GetUserProfile(c *engine.Context) {
	userId, err := strconv.Atoi(c.PathParams["userId"])
	if err != nil {
		c.ServerErrorJson("can not find user id")
		return
	}
	user := UserProfile{
		Id: userId,
		UserName: "Jun",
		Age: 20,
	}
	c.OKJson(user)
}

type UserProfile struct {
	Id int `json:"user_id"`
	UserName string `json:"user_name"`
	Age int  `json:"age"`
}

