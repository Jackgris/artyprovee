package controllers

import (
	"artyprovee/models"
	. "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
	"strings"
)

const (
	IMG_DIR_ORI = "static/images/fotos_ori/"
	IMG_DIR_RZ  = "static/images/fotos_rz/"
)

type Artcontroller struct {
	beego.Controller
	Util
}

//mostrar articulos
func (this *Artcontroller) Art() {

	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		//cargamos plantillas
		LoadTpl(&this.Controller, "art.tpl", true)

		o := orm.NewOrm()
		o.Using("default")

		//llama funcion
		Paginacion(&this.Controller, "articulos")
		//PagArt(this, "articulos")

		//Para la busqueda por referencia
		if this.Ctx.Input.Method() == "POST" {
			buscaRef := this.Ctx.Request.FormValue("buscame")

			//buscamos el id correspondiente
			var articulos models.Articulos
			err := o.QueryTable("articulos").Filter("RefPropia", buscaRef).One(&articulos, "Id")
			if err == nil {
				this.Redirect("/upda/"+strconv.Itoa(articulos.Id), 302)

			}
		}
	}
}

//adicionar articulos
func (this *Artcontroller) Adda() {
	//reseteamos contador paginas
	cta_pgs = 0
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		//cargamos elementos en la pagina
		LoadTpl(&this.Controller, "addart.tpl", false)

		//mensajes
		flash := beego.ReadFromRequest(&this.Controller)
		if ok := flash.Data["error"]; ok != "" {
			// Display error messages
			this.Data["flash"] = ok
		}

		//cargamos lista de categorias y proveedores *****
		//llamando a la funcion ListCatProv desde util
		this.Data["categoria"] = this.ListCatProv("categoria")
		this.Data["proveedor"] = this.ListCatProv("proveedor")

		o := orm.NewOrm()
		o.Using("default")

		articulos := models.Articulos{}

		//tomamos los datos Id de los datalist, Categoria y proveedor
		//ademas se buscan si existen
		existP, existC := DatosIdDatal(this, &articulos)

		//recojemos los otros datos del Form
		err := this.ParseForm(&articulos)
		if err != nil {
			beego.Error("No se pudo obtener datos: ", err)
		} else {

			//validamos el resto
			valid := validation.Validation{}
			isValid, _ := valid.Valid(articulos)

			if this.Ctx.Input.Method() == "POST" {
				//validamos proveedor y categoria
				if !existP || !existC {
					msg := Sprintf("Se debe crear ")
					if !existP {
						msg += "el Proveedor, " + strconv.Itoa(articulos.IdProveedor) + " !!"
					} else {
						msg += "La Categoria, " + strconv.Itoa(articulos.IdCategoria) + " !!"
					}
					beego.Debug(msg)
					flash.Error(msg)
					return
				}

				if !isValid {
					this.Data["Errors"] = valid.ErrorsMap
				} else {
					//verificamos si ya esta el articulo, por RefPropia

					searchArticle := models.Articulos{RefPropia: articulos.RefPropia}
					err := o.Read(&searchArticle, "RefPropia")
					if err == orm.ErrNoRows {
						beego.Debug("Articulo no encontrado: ", articulos)

						//funcion para cargar archivo-foto al directorio
						LoadFoto(this, &articulos)

						id, err := o.Insert(&articulos)
						flash.Store(&this.Controller)
						if err == nil {
							msg := Sprintf("Articulo insertado id: %d", id)
							beego.Debug(msg)
							flash.Notice(msg)
							flash.Store(&this.Controller)
						} else {
							msg := Sprintf("No se pudo crear nuevo articulo. Por: ", err)
							beego.Debug(msg)
							flash.Error(msg)
							flash.Store(&this.Controller)
						}
					} else {
						msg := "Este articulo:" + articulos.RefPropia + " ya existe!"
						beego.Debug(msg)
						this.Data["refexiste"] = msg
					}

				}

			}
		}
	}
}

//Detalles articulo, editar,borrar
func (this *Artcontroller) Upda() {
	//reseteamos contador paginas
	cta_pgs = 0
	//recuperamos cookies
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos cokies
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		o := orm.NewOrm()
		o.Using("default")

		id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
		if err == nil {
			LoadTpl(&this.Controller, "upart.tpl", false)

			this.Data["Id"] = id

			//busca el registro correspondiente al Id=id y lo almacena en la variable
			//para pasarlo con la plantilla a la pagina
			articulos := models.Articulos{Id: id}
			err := o.Read(&articulos)
			if err == nil {
				this.Data["records"] = articulos
			}
		}

		//cargamos lista de categorias y proveedores *****
		//llamando a la funcion ListCatProv desde util
		this.Data["categoria"] = this.ListCatProv("categoria")
		this.Data["proveed"] = this.ListCatProv("proveedor")

		//con Input se puede recoger datos de la pagina
		//los datos se recojen usando el  name de los elementos
		inputs := this.Input()
		val := inputs.Get("accion")

		//en caso de actualizar
		if val == "up" {

			//buscar IdProveedor en tabla
			prov := inputs.Get("IdProveedor")
			//consulta para ver si el proveedor existe
			exist := o.QueryTable("proveedor").Filter("Id", prov).Exist()

			//id categoria
			//categoria := models.Categoria{}
			cat := inputs.Get("IdCategoria")
			exist1 := o.QueryTable("categoria").Filter("Id", cat).Exist()

			if !exist {
				msg := Sprint("  !!El Cod. de Proveedor NO existe !!")
				beego.Debug(msg)
				this.Data["noticia"] = msg
			} else if !exist1 {
				msg := Sprint("  !!El Cod. de Categoria NO existe !!")
				beego.Debug(msg)
				this.Data["noticia"] = msg
			} else {
				articulos := models.Articulos{Id: id}
				//toma solo el numero del data list de proveedores
				//lo pasa a int y lo asigna al IdProveedor
				//lo mismo para categoria
				articulos.IdProveedor, _ = strconv.Atoi((strings.Split(prov, " ")[0]))
				articulos.IdCategoria, _ = strconv.Atoi((strings.Split(cat, " ")[0]))

				err = this.ParseForm(&articulos)
				o.Update(&articulos)
				this.Redirect("/art", 302)
			}
			//en caso de borrar
		} else if val == "del" {
			o.Delete(&models.Articulos{Id: id})
			this.Redirect("/art", 302)
		}
	}
}

//carga archivos de fotos en el dir y path en la base
func LoadFoto(this *Artcontroller, art *models.Articulos) {
	f := "Foto"

	for i := 1; i < 5; i++ {
		//revisamos si existe dato en el form
		fn := f + strconv.Itoa(i)
		_, ne, err := this.Ctx.Request.FormFile(fn)
		newfn_rz := IMG_DIR_RZ

		if err == nil {
			//verificamos si es una imagen
			esimg := ne.Header.Get("Content-Type")
			slesimg := strings.Split(esimg, "/")

			if slesimg[0] == "image" {

				//para extraer la extension. se pueden pedir otros datos a ne con
				slne := strings.Split(ne.Filename, ".")
				//variable para pasar a resize
				fname := this.Ctx.Request.FormValue("RefPropia") + "_" +
					fn + "." + slne[len(slne)-1]
				//creamos ruta con nuevo nombre de la imagen para grabar en dir foto_ori
				newfn := IMG_DIR_ORI + fname
				//creamos ruta que guarda la imagen resize
				newfn_rz += fname
				//grabamos archivo en la ruta indicada
				this.SaveToFile(fn, newfn)

				//llamamos a la funcion de Resize
				this.Util.ImgRZ(fname)

			} else {

				msg := Sprint("!!El Archivo, ", fn, " NO es imagen !!")
				beego.Debug(msg)
				this.Data["errFile"] = msg
			}

		} else {

			newfn_rz += "logo_sinimg.jpg"
		}

		//damos valor al campo articulos.foto _> ruat de la foto
		switch i {
		case 1:
			art.Foto1 = newfn_rz
		case 2:
			art.Foto2 = newfn_rz
		case 3:
			art.Foto3 = newfn_rz
		case 4:
			art.Foto4 = newfn_rz
		}
	}
}

//toma los datos Id de los datalist
//tambien verifica si existe el proveedor y la categoria
// si son introducidos directamente
func DatosIdDatal(this *Artcontroller, art *models.Articulos) (existP, existC bool) {
	prov := this.Ctx.Request.FormValue("IdProveedor")
	categ := this.Ctx.Request.FormValue("IdCategoria")

	//damos valor al campo IdProveedor y Idcategoria
	prov = (strings.Split(prov, " ")[0])
	categ = (strings.Split(categ, " ")[0])

	//ponemos estos valores en el Form, para que no se tome el valor
	//compuesto del data list de id y nombre
	this.Ctx.Request.Form.Set("IdProveedor", prov)
	this.Ctx.Request.Form.Set("IdCategoria", categ)

	//verificamos si existe prov y categ
	o := orm.NewOrm()
	o.Using("default")

	existP = o.QueryTable("proveedor").Filter("Id", prov).Exist()
	existC = o.QueryTable("categoria").Filter("Id", categ).Exist()
	return
}
