package controllers

import (
	. "LOLDataServer/datadriver"
	"github.com/astaxie/beego"
)

type BackController struct {
	beego.Controller
}

func (c *BackController) Get() {
	c.Data["data"] = HeroLi
	c.TplName = "herolist.html"
}

