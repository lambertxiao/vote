package routers

import (
	"vote/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/campaign", &controllers.CampaignController{})
}
