package controllers

import (
	"artyprovee/models"
	. "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type Provcontroller struct {
	beego.Controller
	Util
}

//var n_rgs, pgs, r_pgs, cta_pgs int
//declaradas en Util

func (this *Provcontroller) Prov() {
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value

	//recuperamos securecookie _= dato guardado
	_, b := this.GetSecureCookie(Secreto, "a&p_p")

	//verificamos cokies
	if !this.VeriCk(ck) {

		this.Redirect("/", 302)
	} else if b {
		LoadTpl(&this.Controller, "prov.tpl", true)
		//llama funcion
		Paginacion(&this.Controller, "proveedor")
	} else {

		this.Redirect("/err_admon", 302)
	}
}

//formulario para ingresar nuevos proveedores
func (this *Provcontroller) Addp() {
	//reseteamos contador paginas
	cta_pgs = 0
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
		LoadTpl(&this.Controller, "addprov.tpl", false)

		flash := beego.ReadFromRequest(&this.Controller)

		if ok := flash.Data["error"]; ok != "" {
			// Display error messages
			this.Data["flash"] = ok
		}

		o := orm.NewOrm()
		o.Using("default")

		proveedor := models.Proveedor{}

		if err := this.ParseForm(&proveedor); err != nil {
			beego.Error("No pudieron obtenerse los datos: ", err)
		} else {
			this.Data["Proveedor"] = proveedor
			valid := validation.Validation{}
			isValid, _ := valid.Valid(proveedor)

			if this.Ctx.Input.Method() == "POST" {
				if !isValid {
					this.Data["Errors"] = valid.ErrorsMap
					beego.Error("Datos no validos")
				} else {
					//Buscamos si ya existe el proveedor, usando nombre de la empresa
					//como no es por la cklave principal se debe especificar campo

					searchProv := models.Proveedor{Empresa: proveedor.Empresa}
					err := o.Read(&searchProv, "Empresa")
					if err == orm.ErrNoRows {
						id, err := o.Insert(&proveedor)
						if err == nil {
							msg := Sprintf("Proveedor creado %d id:", id)
							beego.Debug(msg)
							flash.Notice(msg)
							flash.Store(&this.Controller)
						} else {
							msg := Sprintf("No se pudo crear el proveedor. Reason: ", err)
							beego.Debug(msg)
							flash.Error(msg)
							flash.Store(&this.Controller)
						}
					} else {
						msg := Sprint("Proveedor ya existe")
						beego.Debug(msg)
						flash.Notice(msg)
						flash.Store(&this.Controller)
					}

				}

			}
		}
	} else {

		this.Redirect("/err_admon", 302)
	}
}

//formulario para editar o eliminarproveedores
func (this *Provcontroller) Updp() {
	//reseteamos contador paginas
	cta_pgs = 0
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
		o := orm.NewOrm()
		o.Using("default")

		//recojemos el id del parametro
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

		//cargamos plantilla
		LoadTpl(&this.Controller, "updp.tpl", false)

		//pasamos  id proveedor
		this.Data["Id"] = id

		//cargamos datos de proveedor
		proveedor := models.Proveedor{Id: id}
		err := o.Read(&proveedor)
		if err == nil {
			this.Data["records"] = &proveedor
		} else {
			return
		}

		//recojemos datos de accion con el form
		inputs := this.Input()
		val := inputs.Get("accion")

		if val == "up" {
			if this.Ctx.Input.Method() == "POST" {
				//si se han cambiado dtos y esta todo bn, actualizamos
				if err = this.ParseForm(&proveedor); err == nil {
					o.Update(&proveedor)
					this.Redirect("/prov", 302)
				}
			}
		} else if val == "del" {
			//verificamos si este proveedor tiene articulos
			exist := o.QueryTable("articulos").Filter("IdProveedor", id).Exist()
			if exist {
				msg := Sprint("  !! Este proveedor Tiene Articulos !!")
				beego.Debug(msg)
				this.Data["noticia"] = msg
				return
			} else {
				o.Delete(&models.Proveedor{Id: id})
				this.Redirect("/prov", 302)
			}
		}
	} else {

		this.Redirect("/err_admon", 302)
	}
}
