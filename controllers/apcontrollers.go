package controllers

import (
	"artyprovee/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
)

var (
	Sess    interface{}
	Secreto string
	cpt     *captcha.Captcha
	Grupo   string
)

type Apcontroller struct {
	beego.Controller
	Util
}

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

func (this *Apcontroller) Home() {
	//reseteamos contador paginas
	cta_pgs = 0
	//recuperamos securecookie
	s, _ := this.GetSecureCookie(Secreto, "a&p_p")

	this.Data["usuario"] = s
	//recuperamos cookie
	ck_nav, _ := this.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	ck := ck_nav.Value
	//verificamos el valor con la funcion
	if !this.VeriCk(ck) {
		this.Redirect("/", 302)
	} else {
		LoadTplap(this, "header.tpl", "home.tpl")
	}
}

func (this *Apcontroller) Index() {
	//cargamos plantillas pagina Index
	LoadTplap(this, "header_ini.tpl", "index.tpl")

	//para el capcha
	v_capcha := cpt.VerifyReq(this.Ctx.Request)
	this.DestroySession()

	//si el idusuario que hay en la cokie existe en la base,
	//revisamos la variable sess, si esta es nil, pasamos la pagina de login
	Sess = this.GetSession("a&p")

	//si sess es nil
	if Sess == nil {

		u := this.Ctx.Request.FormValue("usuario")
		pw := this.Ctx.Request.FormValue("password")

		o := orm.NewOrm()
		o.Using("default")
		var usuarios models.Usuarios
		qs := o.QueryTable("usuarios")
		res := qs.Filter("usuario", u).One(&usuarios, "usuario", "grupo", "passwd")
		Grupo = usuarios.Grupo

		if res == nil && this.Ctx.Input.Method() == "POST" {

			//comparamos datos de usuario y passwd
			if u == usuarios.Usuario && pw == usuarios.Passwd {
				if v_capcha == true {

					//damos valor 1
					this.SetSession("a&p", int(1))
					Sess = this.GetSession("a&p")

					//para la cokie segura
					Secreto = "a&" + this.Ctx.Request.FormValue("captcha") + "&p"
					//creamos cookie segura para guardar usuario para los privilegios
					//si el usuario es administrador guarda el usuario en la
					//cookie segura, de lo contrario guarda nada como usuario
					if Grupo == "Administrador" {
						this.SetSecureCookie(Secreto, "a&p_p", u)
					} else {
						//usamos el string y no la variable como parametro
						//para recibir false
						this.SetSecureCookie("Secreto", "a&p_p", u)
					}

					//lo manda al home
					this.Redirect("/home", 307)

				} else {
					this.Data["mensaje"] = "caracteres no coinciden!!"
				}
			} else {
				this.Data["mensaje"] = " Usuario o Contraseña errada!!"
			}
		}
	} else {
		//subimos valor a la sesion
		this.SetSession("a&p", Sess.(int)+1)
	}

}

func (this *Apcontroller) Empresa() {
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
		LoadTplap(this, "header.tpl", "empresa.tpl")

		o := orm.NewOrm()
		o.Using("default")

		//cargamos datos de la empresa
		var empresa models.Empresa
		_, err := o.QueryTable("empresa").All(&empresa)

		if err == nil {
			this.Data["records"] = &empresa
		} else {
			beego.Info(err)
			return
		}
		//para pasarle el dato al update
		id := empresa.Id

		//verificamos datos de entrada de opcion del form
		inputs := this.Input()
		val := inputs.Get("accion")

		if this.Ctx.Input.Method() == "POST" && val == "up" {
			empresa := models.Empresa{Id: id}
			if err = this.ParseForm(&empresa); err == nil {
				o.Update(&empresa)
				this.Redirect("/empresa", 302)
			}
		}

	} else {
		this.Redirect("/err_admon", 302)
	}
}

func LoadTplap(this *Apcontroller, h, s string) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = h
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplNames = s
}
