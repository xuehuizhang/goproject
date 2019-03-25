package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/reg",&controllers.MainController{})
    //当路由中指定了自定义的方法，将不再访问默认的Get/Post方法
	beego.Router("/",&controllers.MainController{},"get:ShowLogin;post:HandleLogin")
    beego.Router("/index",&controllers.MainController{},"get:Index;post:HandleArtiType")
    beego.Router("/addarti",&controllers.MainController{},"get:ShowAdd;post:HandleAdd")
    beego.Router("/content",&controllers.MainController{},"get:ShowContent")
    beego.Router("/update",&controllers.MainController{},"get:ShowUpdate;post:HandleUpdate")
    beego.Router("/delete",&controllers.MainController{},"get:DeleteArti")
    beego.Router("/addartitype",&controllers.MainController{},"get:ShowAddArtiType;post:HandleAddArtiType")
    beego.Router("/delArtiType",&controllers.MainController{},"get:DelArtiType")
    beego.Router("/logout",&controllers.MainController{},"get:LogOut")
    beego.Router("/select",&controllers.MainController{},"get:Query")
}
