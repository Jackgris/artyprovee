<div class="container">
  <div class="row">
    <p>
    <div class="hero-text">        
     <h2>Detalles Articulo: {{.Id}} </h2>
    </div>
    </p>
    <p>
     <div>
       <h4>{{.noticia}}</h4>
     </div>
    </p>
  </div>
 </div>

<div class="container">
  <div>
      <form id="articulo" method="POST">
        <div>
            <div class="col-sm-6 col-md-6">
              <a href="#" class="thumbnail">
                <img onclick="javascript:this.width=450;this.height=338" ondblclick="javascript:this.width=80;this.height=20" src="/{{.records.Foto1}}" width="25%" alt="No Hay Foto" /> 
              </a>
            </div>
            <div class="col-sm-6 col-md-6">
              <a href="#" class="thumbnail">
                <img onclick="javascript:this.width=450;this.height=338" ondblclick="javascript:this.width=80;this.height=20" src="/{{.records.Foto2}}" width="25%" alt="No Hay Foto" /> 
              </a>
            </div>
        </div> 
        <div>
            <div class="col-sm-6 col-md-6">
              <a href="#" class="thumbnail">
                <img onclick="javascript:this.width=450;this.height=338" ondblclick="javascript:this.width=80;this.height=20" src="/{{.records.Foto3}}" width="25%" alt="No Hay Foto" /> 
              </a>
            </div>
            <div class="col-sm-6 col-md-6">
              <a href="#" class="thumbnail">
                <img onclick="javascript:this.width=450;this.height=338" ondblclick="javascript:this.width=80;this.height=20" src="/{{.records.Foto4}}" width="25%" alt="No Hay Foto" /> 
              </a>
            </div>
        </div>
   
        <div>
          <div class="form-group col-xs-6 col-md-3">
           <label for="idproveedor">CodProv</label>  
            <input name="IdProveedor" type="text" class="form-control" value="{{.records.IdProveedor}}" list="opciones1" />
              <datalist id="opciones1">
              {{range $proveedor := .proveed}}
              <option>{{$proveedor.Id}} {{$proveedor.Empresa}}</option>
              {{end}}
              </datalist>
          </div>
          <div class="form-group col-xs-6 col-md-3">
            <label for="idcategoria">Categoria</label>  
            <input name="IdCategoria" type="text" class="form-control" value="{{.records.IdCategoria}}" list="opciones2" />
            <datalist id="opciones2">
              {{range $categ := .categoria}}
            <option>{{$categ.Id}} {{$categ.Categoria}}</option>
              {{end}}
            </datalist>
          </div>
          <div class="form-group col-xs-6 col-md-3">
             <label for="refproveedor">RefProv</label>  
             <input name="RefProveedor" type="text"  class="form-control" value="{{.records.RefProveedor}} {{if .Errors.RefProveedor}}(***!!!{{.Errors.RefProveedor}}!!!***){{end}}"  tabindex="2" />
          </div>

          <div class="form-group col-xs-6 col-md-3">
              <label for="refpropia">RefPro</label>  
              <input name="RefPropia" type="text"  class="form-control" value="{{.records.RefPropia}} {{if .Errors.RefPropia}}(***!!!{{.Errors.RefPropia}}!!!***){{end}}" required tabindex="2" />
          </div>
        </div>

        <div>
        <div class="form-group col-xs-6 col-md-4" >
          <label for="descripcion">Descrip</label>  
          <input name="Descripcion" type="text"  class="form-control" value="{{.record.Descripcion}} {{if .Errors.Descripcion}}(***!!!{{.Errors.Descripcion}}!!!***){{end}}" required tabindex="3" />
        </div>
        
      
        <div class="form-group col-xs-6 col-md-4">
          <label for="color">Color</label>  
          <input name="Color" type="text"  class="form-control" value="{{.records.Color}} {{if .Errors.Color}}(***!!!{{.Errors.Color}}!!!***){{end}}"  tabindex="4" />
        </div>

        <div class="form-group col-xs-6 col-md-4">
          <label for="material">Material</label>  
          <input name="Material" type="text"  class="form-control" value="{{.records.Material}}"  tabindex="5" />
        </div>
        </div>

        <div>
           
         <div class="form-group col-xs-4">
              <label for="peso">Peso</label>  
              <input name="Peso" type="text"  class="form-control" value="{{.records.Peso}}"  tabindex="6" />
            </div>
               
            <div class="form-group col-xs-4">
              <label for="empaque">Empaq</label>  
              <input name="Empaque" type="text"  class="form-control" value="{{.records.Empaque}}" tabindex="6" />
            </div>

            <div class="form-group col-xs-4">
              <label for="volempaque">Vol</label>  
              <input name="Volempaque" type="text"  class="form-control" value="{{.records.Volempaque}}"  tabindex="7" />
            </div>
         </div>
          
          <div class="form-group col-xs-12 col-md-12">
       <div class="alert alert alert-danger">
         <strong>¡Tenga presente!</strong> Los cambios son IRREVERSIBLES!!.
       </div>
       <div class="has-warning" >
         <label for="nombre" class="control-label" title="!!LOS CAMBIOS SON IRREVERSIBLES!!">ACCIONES SOBRE EL REGISTRO </label>
          <ul>
           <li>
             <input type="radio" name="accion"  title="!!LOS CAMBIOS SON IRREVERSIBLES!!" value="up">Editarlo
           </li>
           <li>
             <input type="radio" name="accion" title="!!LOS CAMBIOS SON IRREVERSIBLES!!" value="del">ELiminarlo
           </li>
          </ul>  
        </div >
        <div >
         <input type="submit" value="enviar" title="!!LOS CAMBIOS SON IRREVERSIBLES!!" class="btn btn-default btn-lg" > 
         <input type="reset" value="cancelar" class="btn btn-default btn-lg"  />
       </div>
      </div> 

      </form> 
  </div>
</div>

  <div class="container">
      <div><a href="/art"><strong>Regresar a Articulos</strong></a></div>
  </div>

<p>...</p>