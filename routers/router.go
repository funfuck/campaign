// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"campaign/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/campaign",
		beego.NSNamespace("/fgf",
			beego.NSRouter("/getpoint", &controllers.FGFController{}, "post:GetPoint"),
			beego.NSRouter("/addfgf", &controllers.FGFController{}, "post:AddFgf"),
		),
		beego.NSNamespace("/fp",
			beego.NSRouter("/getprizes", &controllers.FPController{}, "post:GetPrizes"),
			beego.NSRouter("/addfp", &controllers.FPController{}, "post:AddFp"),
			beego.NSRouter("/getfp", &controllers.FPController{}, "get:GetFp"),
		),
		beego.NSRouter("/expression", &controllers.FGFController{}, "get:Expression"),
	)
	beego.AddNamespace(ns)
}
