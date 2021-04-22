package controllers

import (
	. "LOLDataServer/datadriver"
	"fmt"
	"github.com/astaxie/beego"
)


type HeroController struct {
	beego.Controller
}

func (c *HeroController) Post() {
	i := c.GetString("id")
	ind,er :=ForrangeID(HeroLi,i)
	//index = ind
	if !er{
		fmt.Println("检索失败，无此英雄ID",ind)
		return
	}
	c.Data["Hero"] = HeroLi[ind].Hero
	c.Data["Skins"] = HeroLi[ind].Skins
	c.Data["Spells"] = HeroLi[ind].Spells
	c.TplName = "hero.html"
}
func (c *HeroController) Get() {
	i := c.GetString("id")
	ind,er :=ForrangeID(HeroLi,i)
	//index = ind
	if !er{
		fmt.Println("检索失败，无此英雄ID",ind)
		return
	}
	c.Data["Hero"] = HeroLi[ind].Hero
	c.Data["Skins"] = HeroLi[ind].Skins
	c.Data["Spells"] = HeroLi[ind].Spells
	c.TplName = "hero.html"
}
