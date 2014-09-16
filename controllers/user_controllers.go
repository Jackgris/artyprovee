package controllers

import (
	"artyprovee/models"
	. "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type Usercontroller struct {
	beego.Controller
	Util
}

func (this *Usercontroller) User() {
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value

	//recuperamos securecookie _= dato guardado
	//b =bool
	_, b := this.GetSecureCookie(Secreto, "a&p_p")

	//verificamos cokies
	if !this.VeriCk(ck) {

		this.Redirect("/", 302)
	} else if b {
		//cargamos plantilla
		LoadTplu(this, "user.tpl")
		//ponemos datos de ta tabla
		this.Luser()

		a := this.Ctx.Request.Form
		s := this.Ctx.Request.FormValue("Usuario")

		if a.Get("Eliminar") == "ok" {
			this.Deluser(s)
			this.Luser()
		}
		if a.Get("Adicionar") == "add" {
			this.Add()
			this.Luser()
		}
	} else {

		this.Redirect("/err_admon", 302)
	}
}

func (this *Usercontroller) Add() {

	o := orm.NewOrm()
	o.Using("default")

	//mensajes
	flash := beego.ReadFromRequest(&this.Controller)

	usuarios := models.Usuarios{}
	err := this.ParseForm(&usuarios)
	if err != nil {
		beego.Error("No se pudo obtener datos: ", err)
	} else if this.Ctx.Input.Method() == "POST" {
		//verificamos si usuario existe
		user := this.Ctx.Request.FormValue("Usuario")
		exist := o.QueryTable("Usuarios").Filter("Usuario", user).Exist()

		if !exist {

			//insertamos datos
			id, err := o.Insert(&usuarios)
			flash.Store(&this.Controller)
			if err == nil {
				msg := Sprintf("Usuario insertado id: %d", id)
				beego.Debug(msg)
				flash.Notice(msg)

			} else {
				msg := Sprintf("No se pudo crear nuevo usuario. Por: ", err)
				beego.Debug(msg)
				flash.Error(msg)
			}

		} else {
			msg := Sprint("  !! El usuario YA existe !!")
			beego.Debug(msg)
			flash.Error(msg)
		}
	}

}

//carga pantillas
func LoadTplu(this *Usercontroller, s string) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplNames = s
}

//carga listado de usuarios
func (this *Usercontroller) Luser() {
	//revisar datos cargados y mostrarlos en la pantalla
	o := orm.NewOrm()
	o.Using("default")

	var mapslist []orm.Params
	_, err := o.QueryTable("usuarios").Values(&mapslist)
	if err != nil {
		beego.Info(err)
	} else {
		//mostramos datos
		this.Data["records"] = mapslist
	}
}

//borra elementos de la tabla usuarios
func (this *Usercontroller) Deluser(s string) {
	o := orm.NewOrm()
	o.Using("default")

	//recuperamos id usuario usuario
	usuario := models.Usuarios{Usuario: s}
	o.Read(&usuario, "Usuario")

	if this.Ctx.Input.Method() == "POST" {
		//borramos usuario
		o.Delete(&models.Usuarios{Id: usuario.Id})
	}
}
