package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type Conscontroller struct {
	beego.Controller
	Util //cargamos para acceder a los metodos
}

func (this *Conscontroller) Cons() {
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		LoadTplcon(this, "cons.tpl", true)
		//colocamos parametros en las opciones de la lista
		this.Data["cons_a"] = "ap"
		this.Data["cons_b"] = "ac"
		this.Data["cons_c"] = "axp"
		this.Data["cons_d"] = "axc"

		o := orm.NewOrm()
		o.Using("default")

		var maps []orm.Params

		para := this.Ctx.Input.Param(":cons")

		//titulos fijos
		this.Data["cabeza_tres"] = "Cod. Art"
		this.Data["cabeza_cuatro"] = "Ref. Propia"
		this.Data["cabeza_cinco"] = "Descripcion"

		switch para {
		case "ap":
			this.Data["titulocons"] = "Articulos por Proveedor"
			this.Data["cabeza_uno"] = "Cod. Prov."
			this.Data["cabeza_dos"] = "Prov."

			//usamos una consulta join para poner datos en pagina
			_, err := o.Raw(`SELECT articulos.id_proveedor, proveedor.empresa, 
        articulos.id, ref_propia,articulos.descripcion
        FROM articulos JOIN proveedor 
        WHERE articulos.id_proveedor = proveedor.id
        GROUP BY articulos.id_proveedor, proveedor.empresa, articulos.id, ref_propia,
        articulos.descripcion`).Values(&maps)

			if err == nil {
				this.Data["records"] = maps
			}

		case "ac":
			this.Data["titulocons"] = "Articulos por Categoria"
			this.Data["cabeza_uno"] = "Cod. Cate."
			this.Data["cabeza_dos"] = "Categoria."

			//usamos una consulta join para poner datos en pagina
			_, err := o.Raw(`SELECT categoria.cod_categoria, categoria.categoria, 
        articulos.id, ref_propia,articulos.descripcion
        FROM articulos JOIN categoria 
        WHERE articulos.id_categoria = categoria.id
        GROUP BY categoria.cod_categoria, categoria.categoria, 
        articulos.id, ref_propia,articulos.descripcion`).Values(&maps)

			if err == nil {
				this.Data["records"] = maps
			}
		case "axp":
			this.Data["titulocons"] = "Articulos por un Proveedor"

			//habilitamos form de envio
			this.Data["form_axp"] = true

			//cargamos datalist desde util_controllers
			this.Data["proveedor"] = this.ListCatProv("proveedor")

			//recojemos valor del data ListCatProv()
			prov := this.Ctx.Request.FormValue("IdProveedor")
			provee := (strings.Split(prov, " ")[0])

			//ponemos valor id en el campo input
			this.Ctx.Request.Form.Set("IdProveedor", provee)

			if this.Ctx.Input.Method() == "POST" {
				//empresa
				emp := (strings.Split(prov, " ")[1])
				//pasamos emp al titulo1
				this.Data["titulocons1"] = "Proveedor :" + emp

				//recuperamos parametro
				idprov := this.Ctx.Request.FormValue("IdProveedor")

				//hacemos consulta para llamar datos
				i, err := o.Raw(`SELECT articulos.id, ref_propia,articulos.descripcion
            FROM articulos  
            WHERE articulos.id_proveedor =? 
            ORDER BY articulos.id`, idprov).Values(&maps)
				if err == nil {
					beego.Debug("metodo post ok")
					s := " :" + strconv.Itoa(int(i)) + " Registros"
					this.Data["num"] = s
					this.Data["records"] = maps
				}
			}
		case "axc":
			this.Data["titulocons"] = "Articulos por una Categoria"
			//habilitamos form de envio
			this.Data["form_axc"] = true

			//cargamos datalist desde util_controllers
			this.Data["categoria"] = this.ListCatProv("categoria")

			//recojemos valor del data ListCatProv()
			cat := this.Ctx.Request.FormValue("IdCategoria")
			trun_id_cat := (strings.Split(cat, " ")[0])

			//ponemos valor id en el campo input
			this.Ctx.Request.Form.Set("IdCategoria", trun_id_cat)

			if this.Ctx.Input.Method() == "POST" {
				//empresa
				trun_categ_cat := (strings.Split(cat, " ")[1])
				//pasamos emp al titulo1
				this.Data["titulocons1"] = "Categoria :" + trun_categ_cat

				//recuperamos parametro
				idcat := this.Ctx.Request.FormValue("IdCategoria")

				//hacemos consulta para llamar datos
				i, err := o.Raw(`SELECT articulos.id, ref_propia,articulos.descripcion
            FROM articulos  
            WHERE articulos.id_categoria =? 
            ORDER BY articulos.id`, idcat).Values(&maps)
				if err == nil {
					beego.Debug("metodo post ok")
					s := " :" + strconv.Itoa(int(i)) + " Registros"
					this.Data["num"] = s
					this.Data["records"] = maps
				}
			}
		}
	}
}

//carga plantillas de las paginas
func LoadTplcon(this *Conscontroller, s string, b bool) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.Data["IsCons"] = b

	this.TplNames = s
}
