<div class="container">
  <div class="row">
    <div class="hero-text">        
     <h2>Detalles Proveedor: {{.Id}} </h2>
  
      <div>
        <h4>{{.noticia}}</h4>
      </div>
       
    </div>
  </div>
 </div>


<div class="container">
  <div >    
      <form  class="col-md-8" id="proveedor" method="POST">
        
        <div >
          <div  class="form-group col-xs-10 col-md-4" >
            <label class="sr-only" for="nit">Nit </label>  
            <input name="Nit" type="text" class="form-control" value="{{.records.Nit}}"  tabindex="1" />
          </div>
    
          <div  class="form-group col-xs-10 col-md-8">
            <label class="sr-only"  for="empresa">Empresa</label>  
            <input name="Empresa" type="text"  class="form-control" value="{{.records.Empresa}}" required tabindex="2" />
          </div>
        </div>

        <div>
          <div  class="form-group col-xs-10 col-md-6">
            <label class="sr-only"  for="contacto">Contacto</label>  
            <input name="Contacto" type="text"  class="form-control" value="{{.records.Contacto}}"  tabindex="3" />
          </div>
          
          <div  class="form-group col-xs-10 col-md-6">
            <label class="sr-only"  for="email">Email</label>  
            <input name="Email" type="email"  class="form-control" value="{{.records.Email}}"  tabindex="4" />
          </div>
        </div>

        <div>
          <div class="form-group col-xs-10 col-md-4">
            <label class="sr-only"  for="pais">Pais</label>  
            <input name="Pais" type="text"  class="form-control" value="{{.records.Pais}}" required  tabindex="7" />
          </div>
          <div class="form-group col-xs-10 col-md-4">
            <label class="sr-only"  for="ciudad">Ciudad</label>  
            <input name="Ciudad" type="text"  class="form-control" value="{{.records.Ciudad}}"  required tabindex="8" />
          </div>
          <div class="form-group col-xs-10 col-md-4">
            <label class="sr-only"  for="Tel">Tel</label>  
            <input name="Tel" type="text"  class="form-control" value="{{.records.Tel}}"  tabindex="6" />
          </div>
        </div>

        <div>
          <div class="form-group col-xs-10 col-md-5">
            <label class="sr-only"  for="url">Url</label>  
            <input name="Url" type="text"  class="form-control" value="{{.records.Url}}"  tabindex="5" />
          </div>
          <div class="form-group col-xs-10 col-md-7">
            <label class="sr-only"  for="direccion">Direccion</label>  
            <input name="Direccion" type="text"  class="form-control" value="{{.records.Direccion}}"  tabindex="9" />
          </div>
        </div>

        <div class="form-group col-xs-12 col-md-12">
          <label class="sr-only"  for="nota">Nota</label>  
          <input name="Nota" type="text"  class="form-control" value="{{.records.Nota}}"  tabindex="10" />
        </div>
        
        <p></p>
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
<div>
  <div class="container">
      <div><a href=" {{urlfor "Provcontroller.Prov"}}"><strong>Regresar a Provedores</strong></a></div>
  </div>
</div>
