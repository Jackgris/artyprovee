package main

import (
	models "artyprovee/models"
	_ "artyprovee/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "user:Passwd@/ayp?charset=utf8")
	orm.RegisterModel(new(models.Articulos), new(models.Categoria),
		new(models.Proveedor), new(models.Usuarios), new(models.Empresa))

}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	//activado para ver log en la consola
	//orm.Debug = true

	//Crea las tablas
	/*
		name := "default"
		force := false
		verbose := true
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			beego.Info(err)
		}
	*/

	beego.EnableAdmin = true
	beego.AdminHttpAddr = "localhost"
	beego.AdminHttpPort = 8088
	beego.Run()
}
