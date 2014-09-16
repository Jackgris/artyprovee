<div class="container">
  <div class="row">
    <div class="hero-text">
      <p>
        <h1 >Nuevo Articulo</h1> 
      </p>
      <p>
        {{if .flash.error}}
       <blockquote >{{.flash.error}}</blockquote>
        {{end}}
        {{if .flash.notice}}
        <blockquote >{{.flash.notice}}</blockquote>
        {{end}}
      </p>
      <p>
        {{.errFile}}{{.refexiste}}
      </p>
     </div>
  </div>
 </div>


<div class="container">
  <div class="row">
        
    <p>          
      <form  id="articulos" enctype="multipart/form-data" method="POST">
               
        <div class="form-group col-xs-6 col-md-6">
          <label class="sr-only"  for="idproveedor">IdProvedor</label>  
          <input name="IdProveedor" type="text"  class="form-control" placeholder="Codigo Proveedor" required tabindex="1" list="opciones1"/>
            <datalist id="opciones1">
              {{range $prov := .proveedor}}
              <option>{{$prov.Id}} {{$prov.Empresa}}</option>
              {{end}}
            </datalist>
        </div>

        <div class="form-group col-xs-6 col-md-6">
          <label class="sr-only"  for="idcategoria">Categoria</label>  
          <input name="IdCategoria" type="text" class="form-control" placeholder="Codigo Categoria" required list="opciones2" />
            <datalist id="opciones2">
              {{range $categ := .categoria}}
              <option>{{$categ.CodCategoria}} {{$categ.Categoria}}</option>
              {{end}}
            </datalist>
        </div>

        <div>          
          <div class="form-group col-xs-6 col-md-6">
            <label class="sr-only"  for="refproveedor">RefProveedor</label>  
            <input name="RefProveedor" type="text"  class="form-control" placeholder="RefProveedor {{if .Errors.RefProveedor}}(***!!!{{.Errors.RefProveedor}}!!!***){{end}}"  tabindex="2" />
          </div>
          <div class="form-group col-xs-6 col-md-6">
            <label class="sr-only"  for="refpropia">RefPropia</label>  
            <input name="RefPropia" type="text"  class="form-control" placeholder="RefPropia {{if .Errors.RefPropia}}(***!!!{{.Errors.RefPropia}}!!!***){{end}}" required tabindex="2" />
          </div>
        </div>

        <div>
          <div class="form-group ">
            <label class="sr-only"  for="descripcion">Descripcion</label>  
            <input name="Descripcion" type="text"  class="form-control" placeholder="Descripcion {{if .Errors.Descripcion}}(***!!!{{.Errors.Descripcion}}!!!***){{end}}" required tabindex="3" />
          </div>
        </div>

        <div>
          <div class="form-group col-xs-6 col-md-6">
            <label class="sr-only"  for="color">Color</label>  
            <input name="Color" type="text"  class="form-control" placeholder="Color {{if .Errors.Color}}(***!!!{{.Errors.Color}}!!!***){{end}}"  tabindex="4" />
          </div>
          <div class="form-group col-xs-6 col-md-6">
            <label class="sr-only"  for="material">Material</label>  
            <input name="Material" type="text"  class="form-control" placeholder="Material"  tabindex="5" />
          </div>
        </div>
        
        <div>
          <div class="form-group col-xs-4 col-md-4">
            <label class="sr-only"  for="peso">Peso</label>  
            <input name="Peso" type="text"  class="form-control" placeholder="Peso"  tabindex="6" />
          </div>
          <div class="form-group col-xs-4 col-md-4">
            <label class="sr-only"  for="empaque">Empaque</label>  
            <input name="Empaque" type="text"  class="form-control" placeholder="Empaque"  tabindex="7" />
          </div>
          <div class="form-group col-xs-4 col-md-4">
            <label class="sr-only"  for="volempaque">Volempaque</label>  
            <input name="Volempaque" type="text"  class="form-control" placeholder="Volumen de empaque"  tabindex="8" />
          </div>
        </div>

        <p>
          <div class="form-group col-xs-6 col-md-6">
            <label   for="foto1">Foto1</label>  
            <input name="Foto1" type="file"  class="form-control" tabindex="9" value="{{.fotico}}" />
          </div>  
          <div class="form-group col-xs-6 col-md-6">
            <label   for="foto2">Foto2</label>  
            <input name="Foto2" type="file"  class="form-control" tabindex="10" />
          </div>
        </p>
        <div>
          <div class="form-group col-xs-6 col-md-6">
            <label   for="foto3">Foto3</label>  
            <input name="Foto3" type="file"  class="form-control" tabindex="11" />
          </div>
          <div class="form-group col-xs-6 col-md-6">
            <label   for="foto4">Foto4</label>  
            <input name="Foto4" type="file"  class="form-control" tabindex="12" />
          </div>
        </div>

        <input type="submit" form="articulos"value="Crear Articulo" class="btn btn-default" tabindex="8" /> &nbsp;
        <input type="reset" form="articulos"value="Cancelar" class="btn btn-default" tabindex="13" />
      </form>
    </p>
  </div>
</div> 
<div class="container">
      <div><a href="/art"><strong>Regresar a Articulos</strong></a></div>
  </div>