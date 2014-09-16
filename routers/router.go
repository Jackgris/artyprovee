package routers

import (
	"artyprovee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Apcontroller{}, "*:Index")
	beego.Router("/home", &controllers.Apcontroller{}, "get,post:Home")
	beego.Router("/empresa", &controllers.Apcontroller{}, "get,post:Empresa")
	beego.Router("/user", &controllers.Usercontroller{}, "get,post:User")
	beego.Router("/prov", &controllers.Provcontroller{}, "get:Prov")
	beego.Router("/addp", &controllers.Provcontroller{}, "get,post:Addp")
	beego.Router("/updp/:id([0-9]+)", &controllers.Provcontroller{}, "get,post:Updp")
	beego.Router("/art", &controllers.Artcontroller{}, "get,post:Art")
	beego.Router("/adda", &controllers.Artcontroller{}, "get,post:Adda")
	beego.Router("/upda/:id([0-9]+)", &controllers.Artcontroller{}, "get,post:Upda")
	beego.Router("/cat", &controllers.Catcontroller{}, "get,post:Cat")
	beego.Router("/addc", &controllers.Catcontroller{}, "get,post:Addc")
	beego.Router("/updc/:id([0-9]+)", &controllers.Catcontroller{}, "get,post:Updc")
	beego.Router("/cons", &controllers.Conscontroller{}, "get,post:Cons")
	beego.Router("/cons/:cons(ap)", &controllers.Conscontroller{}, "get,post:Cons")
	beego.Router("/cons/:cons(ac)", &controllers.Conscontroller{}, "get,post:Cons")
	beego.Router("/cons/:cons(axp)", &controllers.Conscontroller{}, "get,post:Cons")
	beego.Router("/cons/:cons(axc)", &controllers.Conscontroller{}, "get,post:Cons")
	beego.Router("/img", &controllers.Imgcontroller{}, "get,post:Img")
	beego.Router("/err_admon", &controllers.Util{}, "get,post:Err_admon")
}
