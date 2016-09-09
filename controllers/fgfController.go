package controllers

import (
	"github.com/astaxie/beego"
	r "campaign/databases/redis"
	"encoding/json"
	"github.com/zdebeer99/goexpression"
	"strconv"
	"campaign/models"
	"campaign/vars"
	"time"
)

type FGFController struct {
	beego.Controller
	res vars.Response
}

func (c *FGFController) GetPoint() {
	c.res.Path = c.Ctx.Request.RequestURI
	c.res.TimeStamp = time.Now()

	// get active fgf in mysql
	fgf := models.GetActiveFGF()
	t, _ := time.Parse(time.RFC3339, c.Ctx.Input.Header("now"))

	if fgf == nil {
		// not found campaign
		c.res.Success = false
		c.res.ResultCode = "001"
		c.res.ErrorDescription = "not found campaign"
	} else if models.IsInCampaignDuration(&t, fgf) {
		// match point
		friendNo, _ := strconv.Atoi(c.Ctx.Input.Header("friendNo"))
		fgf.MatchedUpline = models.MatchedPoint(fgf.Uplines, friendNo)

		c.res.Success = true
		c.res.ResultCode = "000"
		c.res.Result = fgf
	} else {
		// not in campaign duration
		c.res.Success = false
		c.res.ResultCode = "002"
		c.res.ErrorDescription = "out of campaign duration"
	}

	c.Data["json"] = c.res
	c.ServeJSON()
}

func (c *FGFController) AddFgf() {
	var fgf models.TFriendGetFriend
	json.Unmarshal(c.Ctx.Input.RequestBody, &fgf)

	fgfStr, _ := json.Marshal(fgf)

	session := r.Conn()
	defer session.Close()
	session.Do("SET", "dooadsCampaignFgf", fgfStr)
}

func (c *FGFController) Expression() {
	context := map[string]interface{}{"x": 10.25, "y": 50, "z": 3.0,}
	ans := goexpression.Eval("1 + x  *( 50-47)/z", context)
	c.Ctx.WriteString(strconv.FormatFloat(ans, 'f', 6, 64))
}