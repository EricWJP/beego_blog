package main

import (
	_ "b_blog/helpers"
	"b_blog/models"
	_ "b_blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
