<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Categorias</h1>
         {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>C&oacute;digo</th>
          <th>Des. Categoria</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>
            <a class="id_cat" href="updc/{{$record.CodCategoria}}" title="Mas   Datos - Categoria"
           > {{$record.CodCategoria}}
            </a>
          </td>
          <td>{{$record.Categoria}}</td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="8"><a href="{{urlfor "Catcontroller.Addc"}}" title="Ingresar castegoria">Ingresar Categoria </a></td>
          <td colspan="3"></td>

        </tr>
      </tfoot>
    </table>
  </div>
</div>