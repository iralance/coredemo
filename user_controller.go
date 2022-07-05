package main

import (
	"github.com/iralance/coredemo/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	//超时
	time.Sleep(3 * time.Second)
	foo, _ := c.QueryString("foo", "def")
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}
