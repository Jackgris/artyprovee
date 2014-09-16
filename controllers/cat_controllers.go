package controllers

import (
	"artyprovee/models"
	. "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type Catcontroller struct {
	beego.Controller
	Util
}

//muestra categorias en pantalla
func (this *Catcontroller) Cat() {
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {

		//cargamos elementos en la pagina
		LoadTpl(&this.Controller, "cat.tpl", false)

		o := orm.NewOrm()
		o.Using("default")

		var categoria []*models.Categoria
		num, err := o.QueryTable("categoria").All(&categoria)

		if err != orm.ErrNoRows && num > 0 {
			this.Data["records"] = categoria
		}
	}
}

//formulario para adicionar categorias
func (this *Catcontroller) Addc() {
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		//cargamos elementos en la pagina
		LoadTpl(&this.Controller, "addc.tpl", false)

		o := orm.NewOrm()
		o.Using("default")

		//mensajes
		flash := beego.ReadFromRequest(&this.Controller)
		if ok := flash.Data["error"]; ok != "" {
			// Display error messages
			this.Data["flash"] = ok
		}

		categoria := models.Categoria{}
		err := this.ParseForm(&categoria)
		if err != nil {
			beego.Error("No se pudo obtener datos: ", err)
		} else if this.Ctx.Input.Method() == "POST" {
			//verificamos si la nueva categoria existe
			codcat := this.Ctx.Request.FormValue("CodCategoria")
			exist := o.QueryTable("categoria").Filter("CodCategoria", codcat).Exist()

			if !exist {

				//validamos datos
				valid := validation.Validation{}
				isValid, _ := valid.Valid(categoria)
				if !isValid {
					msg := Sprint("  !! Los datos no son validos !!")
					this.Data["noticia"] = msg
				} else {
					//insertamos datos
					id, err := o.Insert(&categoria)
					flash.Store(&this.Controller)
					if err == nil {
						msg := Sprintf("categoria insertada id: %d", id)
						beego.Debug(msg)
						flash.Notice(msg)
						flash.Store(&this.Controller)
					} else {
						msg := Sprintf("No se pudo crear nuevo categoria. Por: ", err)
						beego.Debug(msg)
						flash.Error(msg)
						flash.Store(&this.Controller)
					}
				}
			} else {
				msg := Sprint("  !! La categoria YA existe !!")
				beego.Debug(msg)
				this.Data["noticia"] = msg
			}
		}
	}
}

//permite ver categorias borrarlas y editar
func (this *Catcontroller) Updc() {
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		o := orm.NewOrm()
		o.Using("default")

		//para la categoria tenemos en cuenta que el cod de categoria
		//no es el id real de la tabla
		id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

		//buscamos el id correspondiente al codcat
		var cat models.Categoria
		o.QueryTable("categoria").Filter("CodCategoria", id).One(&cat, "Id")

		//cargamos plantillas
		LoadTpl(&this.Controller, "updc.tpl", false)

		//pasamos valor real id para cargar datos
		//busca el registro correspondiente al Id=id y lo almacena en la variable
		//para pasarlo con la plantilla a la pagina
		this.Data["Id"] = cat.Id
		categoria := models.Categoria{Id: cat.Id}
		err := o.Read(&categoria)
		if err == nil {
			this.Data["records"] = &categoria
		} else {
			return
		}

		//recojemos datos de accion con el form
		inputs := this.Input()
		val := inputs.Get("accion")

		if val == "up" {
			if this.Ctx.Input.Method() == "POST" {

				//verificamos si cambio la categoria o el codcategoria
				i, _ := strconv.Atoi(this.Ctx.Request.FormValue("CodCategoria"))
				if categoria.CodCategoria != i {
					//verificamos si el codcategoria existe
					codcat := this.Ctx.Request.FormValue("CodCategoria")
					exist := o.QueryTable("categoria").Filter("CodCategoria", codcat).Exist()
					if exist {
						msg := Sprint("  !! El Cod. de categoria YA existe !!")
						beego.Debug(msg)
						this.Data["noticia"] = msg
						return
					}
				} else if categoria.Categoria != this.Ctx.Request.FormValue("Categoria") {
					//verificamos si la categoria existe
					cate := this.Ctx.Request.FormValue("Categoria")
					exist := o.QueryTable("categoria").Filter("Categoria", cate).Exist()
					if exist {
						msg := Sprint("  !! La categoria  YA existe !!")
						beego.Debug(msg)
						this.Data["noticia"] = msg
						return
					}
				}

				//verificamos si la categoria ya see sta usando
				codcat := this.Ctx.Request.FormValue("CodCategoria")
				exist := o.QueryTable("articulos").Filter("IdCategoria", codcat).Exist()
				if exist {
					msg := Sprint("  !! Categoria usandose No se puede Modificar ni Borrar !!")
					beego.Debug(msg)
					this.Data["noticia"] = msg
					return
					//si se han cambiado dtos y esta todo bn, actualizamos
				} else if err = this.ParseForm(&categoria); err == nil {
					o.Update(&categoria)
					this.Redirect("/cat", 302)
				}
			}
		} else if val == "del" {
			//verificamos si esta categoria tiene articulos
			exist := o.QueryTable("articulos").Filter("IdCategoria", cat.Id).Exist()
			if exist {
				msg := Sprint("  !! Esta Categoria Tiene Articulos !!")
				beego.Debug(msg)
				this.Data["noticia"] = msg
				return
			} else {
				o.Delete(&models.Proveedor{Id: cat.Id})
				this.Redirect("/cat", 302)
			}
		}
	}
}
