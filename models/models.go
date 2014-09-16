package models

type Articulos struct {
	Id           int `form:"-"`
	IdProveedor  int
	IdCategoria  int
	RefProveedor string `valid:"MaxSize(20)"`
	RefPropia    string `orm:"index";valid:"MaxSize(20)"`
	Descripcion  string `valid:"MaxSize(20)"`
	Color        string `valid:"MaxSize(20)"`
	Material     string `valid:"MaxSize(30)"`
	Peso         float64
	Empaque      string
	Volempaque   float64
	Foto1        string
	Foto2        string
	Foto3        string
	Foto4        string
}

type Categoria struct {
	Id           int `form:"-"`
	CodCategoria int
	Categoria    string `valid:"MaxSize(20)"`
}

type Proveedor struct {
	Id        int    `form:"-"`
	Nit       string `valid:"MaxSize(20)"`
	Empresa   string `valid:"MinSize(3);MaxSize(20)"`
	Contacto  string `valid:"MaxSize(20)"`
	Email     string
	Url       string
	Tel       string
	Pais      string `valid:"MaxSize(20)"`
	Ciudad    string `valid:"MaxSize(20)"`
	Direccion string `valid:"MaxSize(30)"`
	Nota      string `valid:"MaxSize(250)"`
}

//grupos:Administrador, usuario,cliente
type Usuarios struct {
	Id      int `form:"-"`
	Usuario string
	Grupo   string
	Passwd  string
}

type Empresa struct {
	Id            int `form:"-"`
	Empresa       string
	NitDni        string
	Representante string
	Direccion     string
	Telefono      string
	Ciudad        string
	Pais          string
}
