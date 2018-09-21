package controllers

import (
	"github.com/astaxie/beego"
    "vote/models"
)

type CampaignController struct {
    beego.Controller
}

func (this CampaignController) Get() {
    var campaign models.Campaign
    c := campaign.GetById("123")
    this.Data["json"] = c
    this.ServeJSON()
}
