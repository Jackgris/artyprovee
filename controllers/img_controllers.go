package controllers

import (
	"github.com/astaxie/beego"
)

type Imgcontroller struct {
	beego.Controller
	Util //cargamos para acceder a los metodos
}

func (this *Imgcontroller) Img() {

	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplNames = "l_img.tpl"

	//para la prueba de inamegens
	datos := this.Util.ListCatProv("articulos")
	this.Data["prueba"] = datos

}
