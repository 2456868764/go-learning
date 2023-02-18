package v2

import (
	"fmt"
	"github.com/2456868764/go-learning/web/pkg/engine"
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
