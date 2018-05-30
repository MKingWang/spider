package main

import (
	"fmt"

	"mk/spider_a/server/queue"

	"github.com/astaxie/beego"

	_ "mk/spider_a/server/router"
)

func main() {
	config()
	go queue.Run()
	beego.Run()
	fmt.Println("ls")
}

func config() {
	beego.BConfig.CopyRequestBody = true
}
