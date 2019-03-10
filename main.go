package main

import (
	_ "beego-project/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/layui", "/static/layui")
	beego.SetStaticPath("/txt", "static")//静态文件路由注册 可以直接访问static文件下的123.txt
	beego.BConfig.WebConfig.TemplateLeft = "[[["
	beego.BConfig.WebConfig.TemplateRight = "]]]"
	beego.Run()
}
