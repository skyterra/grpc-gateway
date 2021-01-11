package main

import (
	_ "xframework/web-server/boot"
	_ "xframework/web-server/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
