package controllers

import (
	"github.com/astaxie/beego"
	"campaign/models"
	"encoding/json"
	r "campaign/databases/redis"
	"fmt"
	"campaign/vars"
	"time"
)

type FPController struct {
	beego.Controller
	res vars.Response
}

func (c *FPController) GetPrizes() {
	c.res.Path = c.Ctx.Request.RequestURI
	c.res.TimeStamp = time.Now()

	// get body
	var param *vars.ParamGetPrizeBody
	json.Unmarshal(c.Ctx.Input.RequestBody, &param)

	// get active flash point from redis (zero or more) & match campaign
	c.res.Result = models.GetMatchedCampaign(param)

	c.Data["json"] = c.res
	c.ServeJSON()
}

func (c *FPController) AddFp() {
	var fp models.TFlashPoint
	json.Unmarshal(c.Ctx.Input.RequestBody, &fp)

	fpStr, _ := json.Marshal(fp)

	session := r.Conn()
	defer session.Close()
	session.Do("SET", "dooadsCampaignFp_" + fmt.Sprint(fp.ID), fpStr)
}

func (c *FPController) GetFp() {
	c.Data["json"] = &models.TFlashPoint{
		Prizes:[]*models.TPrize{{}},
		Adss:[]*models.TFpAds{{}},
		Operators:[]*models.TOperator{{}},
		Genders:[]*models.Gender{{}},
	}
	//c.Data["json"] = &vars.RequestGetPrizesBody{MemberProfile:&vars.MemberProfile{}}
	c.ServeJSON()
}