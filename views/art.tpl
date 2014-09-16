<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Articulos</h1>
         {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    {{.noticia}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row table-responsive">
    <table class="table table-striped table-hover table-condensed">
      <thead>
        <tr>
          <th>C&oacute;digo</th>
          <th>Proveedor</th>
          <th>Categoria</th>
          <th>RefPropia</th>
          <th>Descripci&oacute;n</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>
            <a class="id_art" href="upda/{{$record.id}}" title="Mas   Datos - Articulo"
           > {{$record.id}}
            </a>
          </td>
          <td>
            <a  href="updp/{{$record.id_proveedor}}" title="Mas   Datos - Proveedor"
           > {{$record.id_proveedor}}</a>
          </td>
          <td>{{$record.categoria}}</td>
          <td>{{$record.ref_propia}}</td>
          <td>{{$record.descripcion}}</td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="8"><a href="{{urlfor "Artcontroller.Adda"}}" title="Ingresar Articulo">Ingresar Articulo </a></td>
          <td colspan="3"></td>

        </tr>
      </tfoot>
    </table>
  </div>
</div>

<div class="container">
 <div>
   <form action="" method="post">
    <div class="col-xs-5 col-md-2">
      <button name="pagante" type="submit" value="-20" {{.valor_ante}}>Pagina ante </button>
    </div>
    <div class="col-xs-5 col-md-2">
      <button name="pagsgte" type="submit" value="+20" {{.valor_sgte}}>Pagina sgte</button>
    </div>
    <div class="col-xs-8 col-md-6">
      <h4>Total de Registros: <strong>{{.nreg }}. </strong>Numero de paginas <span class="label label-default">{{.pag}} de {{.npags}}</span></h4>
    </div>
  </form> 
 </div>
</div>
<hr />

