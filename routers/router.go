package routers

import (
	"LOLDataServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/lol", &controllers.BodyController{})
	beego.Router("/heroSpells", &controllers.SpellsController{})
	beego.Router("/herodate", &controllers.HeroController{})
	beego.Router("/back", &controllers.BackController{})

}
