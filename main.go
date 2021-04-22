package main

import (
	. "LOLDataServer/datadriver"
	_ "LOLDataServer/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	err := GetLoldate()
	if !err{
		fmt.Println("英雄信息获取错误")
		return
	}
	//2、静态资源路径设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	fmt.Println("Every is ok ! You can do that !")
   beego.Run()
}

