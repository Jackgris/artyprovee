<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Proveedores</h1>
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

<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>Codigo</th>
          <th>Nit</th>
          <th>Empresa</th>
          <th>Contacto</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td><a  href="updp/{{$record.id}}" title="Mas   Datos - Proveedor"
           > {{$record.id}}</a>
          </td>
          <td>{{$record.nit}}</td>
          <td>{{$record.empresa}}</td>
          <td>{{$record.contacto}}</td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="4"><a href="{{urlfor "Provcontroller.Addp"}}" title="Ingresar Nuevo Proveedor">Ingresar Nuevo Proveedor</a></td>
    
        </tr>
      </tfoot>
    </table>
    <hr />
  </div>
</div>


<div class="container">
 <div>
   <form action="" >
    <div class="col-xs-5 col-md-2">
      <button name="pagante" type="submit" value="-20" {{.valor_ante}}>Pagina ante</button>
    </div>
    <div class="col-xs-5 col-md-2">
      <button name="pagsgte" type="submit" value="+20"  {{.valor_sgte}} >Pagina sgte</button>
    </div>
    <div class="col-xs-8 col-md-6">
      <h4>Total de Registros: <strong>{{.nreg }}. </strong>Numero de paginas <span class="label label-default">{{.pag}} de {{.npags}}</span></h4>
    </div>
  </form> 
 </div>
</div>
<hr />

