<!-->
-. Proveedor y sus articulos
-. Articulos Agrupados por Categoria
-. Articulos Agrupados por proveedor
-. Listado Proveedores x orden alfabetico
-. Listado Articulos por orden alfabetico
-. Articulos por una categoria
-. Articulos por un proveedor
</!-->
<div class="container">
  <div class="row">
    <div class="hero-text">   	
     <h3>{{.titulocons}} <small>{{.num}}</small></h3>
  </div>
</div>
</div>

<div class="container" >	
	{{if .form_axp}}{{template  "consultas/form_axp.tpl".}}{{end}}
	{{if .form_axc}}{{template  "consultas/form_axc.tpl".}}{{end}}
</div>
<div class="container" >
	<h4>{{.titulocons1}}</h4>
</div>
<div class="container">
  <div class="row table-responsive">
    <table class="table table-striped table-hover table-condensed">
      <thead>
        <tr>
          <th>{{.cabeza_uno}}</th><!-->id proveedor o categoria</!-->
          <th>{{.cabeza_dos}}</th><!-->proveedor o categoria</!-->
          <th>{{.cabeza_tres}}</th><!-->codigo del articulo</!-->
          <th>{{.cabeza_cuatro}}</th><!-->referencia propia del articulo</!-->
          <th>{{.cabeza_cinco}}</th><!-->descripcion del articulo</!-->
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>
			<a  href="/updp/{{$record.id_proveedor}}" title="Mas   Datos - Proveedor"
           > {{$record.id_proveedor}}{{$record.cod_categoria}}</td>  
          <td>{{$record.empresa}}{{$record.categoria}}</td>
          <td>
            <a class="id_art" href="/upda/{{$record.id}}" title="Mas   Datos - Articulo"
           > {{$record.id}}
            </a>
          </td>
          <td>{{$record.ref_propia}}</td>
          <td>{{$record.descripcion}}</td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="8"><a href="{{urlfor "Apcontroller.Home"}}" title="Ingresar Articulo">Regresar </a></td>
          <td colspan="3"></td>

        </tr>
      </tfoot>
    </table>
  </div>
</div>