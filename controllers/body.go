package controllers

import (
	. "LOLDataServer/datadriver"
	"github.com/astaxie/beego"
)

type BodyController struct {
	beego.Controller
}

func (c *BodyController) Get() {
	c.Data["data"] = HeroLi
	c.TplName = "herolist.html"
}
