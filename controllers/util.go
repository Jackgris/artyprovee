package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/nfnt/resize"
	"image/jpeg"
	//"image/png"
	"log"
	"os"
	"strconv"
)

var n_rgs, pgs, r_pgs, cta_pgs int

type Util struct {
	beego.Controller
}

//funcion resize imagen
func (u Util) ImgRZ(s string) {

	//cargamos archivo de imageb jpg
	file, err := os.Open("static/images/fotos_ori/" + s)
	if err != nil {
		log.Fatal("en este 1", err)
	}

	// decode jpeg  image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal("en este 2", err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(216, 162, img, resize.Lanczos3)

	out, err := os.Create("static/images/fotos_rz/" + s)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	return

}

//Devuelve un map de informacion para pasarla a los data list
func (u Util) ListCatProv(nlist string) []orm.Params {
	o := orm.NewOrm()
	o.Using("default")

	var mapslist []orm.Params
	_, err := o.QueryTable(nlist).Values(&mapslist)
	if err == nil {
		return mapslist
	}
	return mapslist
}

//Verifica las cokies y la variable Sess
func (u Util) VeriCk(ck string) bool {
	o := orm.NewOrm()
	o.Using("default")

	//verificamos si existe sesion en la base, comparando
	//con la cokie del navegador
	//recuperamos cokie
	//ck_nav, _ := u.Ctx.Request.Cookie("a&p")
	//recuperamos el valor
	//ck := ck_nav.Value
	var maps []orm.Params
	num, err := o.Raw("SELECT session_key FROM session WHERE session_key=?", ck).Values(&maps)

	if num == 0 || err != nil {
		//si no hay cokie

		return false
	} else {
		//si existe, verificamos con la variable de sesion
		if Sess != 1 {
			return false
		} else {
			//si es 1, verificamos si la cookie de base y navegador son iguales
			ck_serv := maps[0]["session_key"]
			//si son iguales mostramos pagina
			if ck == ck_serv {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

//carga las tpl-no me esta funcionando bn
func LoadTpl(this *beego.Controller, s string, b bool) {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.Data["IsArtProv"] = b
	this.Data["place"] = "Busqueda por Id"
	this.Data["nombotton"] = "Buscar Id Prov"
	this.TplNames = s
}

//Funcion error de usuario admon
func (u *Util) Err_admon() {
	u.Layout = "basic-layout.tpl"
	u.LayoutSections = make(map[string]string)
	u.LayoutSections["Header"] = "header_ini.tpl" //header diferenet
	u.LayoutSections["Footer"] = "footer.tpl"
	u.TplNames = "err_admon.tpl"

	//recuperar nombre de usuario del la cokie segura
	s, _ := u.GetSecureCookie(Secreto, "a&p_p")

	u.Data["usuario"] = s
}

func Paginacion(this *beego.Controller, tbl string) {
	r_pgs = 20
	o := orm.NewOrm()
	o.Using("default")

	if cta_pgs == 0 {
		//verificamos si la sonsulta arroja resultados
		cnt, _ := o.QueryTable(tbl).Count()
		if cnt == 0 {
			//si no hay registros
			this.Data["noticia"] = "Lo siento, no existen registros!"
			return
		} else {
			//llamamos  a la funcion para recoger resultados consulta
			//ademas de dar datos para paginar, se van a dar 20
			//resultados por pagina y presentamos los primeros registros
			n_rgs, this.Data["records"] = PagSql(tbl, "0", strconv.Itoa(r_pgs))

			//colocamos valor de paginas
			pgs = n_rgs / r_pgs
			if n_rgs%r_pgs > 0 {
				//aproximamos vlr total paginas
				pgs++
				//aumentamos contador paginass
				cta_pgs++
				if cta_pgs == 1 && pgs == 1 { //>= pgs {
					//cta_pgs = 0
					this.Data["valor_sgte"] = "hidden"
					this.Data["valor_ante"] = "hidden"
				} else if cta_pgs == 1 {
					this.Data["valor_ante"] = "hidden"
				}
				//informacion sobre las paginas
				this.Data["pag"] = cta_pgs
				this.Data["npags"] = pgs
				this.Data["nreg"] = n_rgs
			}
		}
	} else if cta_pgs <= pgs {
		//revisamos boton pulsado
		if v := this.Ctx.Request.FormValue("pagsgte"); v == "+20" {

			//llamamos la consulta con la pagina correspondiente
			//teniendo en cuenta el LIMIT cta_pgs*r_pgs,r_pgs
			_, this.Data["records"] = PagSql(tbl, strconv.Itoa(cta_pgs*r_pgs), strconv.Itoa(r_pgs))

			cta_pgs++
			if cta_pgs == pgs {
				this.Data["valor_sgte"] = "hidden"
			}

		} else if v := this.Ctx.Request.FormValue("pagante"); v == "-20" {

			_, this.Data["records"] = PagSql(tbl, strconv.Itoa((cta_pgs-2)*r_pgs), strconv.Itoa(r_pgs))

			cta_pgs--
			if cta_pgs == 1 { //>= pgs {
				//cta_pgs = 0
				this.Data["valor_ante"] = "hidden"
			}
		}

		//informacion sobre las paginas
		this.Data["pag"] = cta_pgs
		this.Data["npags"] = pgs
		this.Data["nreg"] = n_rgs
	}
}

//realiza una consulta  en la base y devuelve el numero de filas
//para las paginaciones
func PagSql(tbl, pi, regxp string) (int, []orm.Params) {
	o := orm.NewOrm()
	o.Using("default")
	var maps []orm.Params
	var num int

	if tbl == "proveedor" {
		_, err := o.Raw(`SELECT SQL_CALC_FOUND_ROWS * 
    FROM proveedor LIMIT ` + pi + `,` + regxp).Values(&maps)
		o.Raw(`SELECT FOUND_ROWS()`).QueryRow(&num)
		if err != nil {
			beego.Info(err)
		}
	} else if tbl == "articulos" {
		//usamos una consulta join para poner datos en pagina
		_, err := o.Raw(`SELECT articulos.id,articulos.id_proveedor,
    articulos.ref_propia,articulos.descripcion,categoria.categoria 
    FROM articulos JOIN categoria WHERE articulos.id_categoria=
    categoria.id ORDER BY articulos.id  LIMIT ` + pi + `,` + regxp).Values(&maps)
		o.Raw(`SELECT FOUND_ROWS()`).QueryRow(&num)
		if err != nil {
			beego.Info(err)
		}
	}
	return num, maps
}
