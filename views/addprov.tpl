<div class="container">
  <div class="row">
   <div class="hero-text">        
    <h1 class="col-xs-4">Nuevo Proveedor</h1> 
     {{if .flash.error}}
     <blockquote class="col-xs-8">{{.flash.error}}</blockquote>
      {{end}}

      {{if .flash.notice}}
      <blockquote class="col-xs-8">{{.flash.notice}}</blockquote>
      {{end}}
    </div>
  </div>
 </div>


<div class="container">
  <div >
       
      <form  class="col-md-8" id="proveedor" method="POST">
        
        <div >
          <div  class="form-group col-xs-10 col-md-4" >
            <label class="sr-only" for="nit">Nit </label>  
            <input name="Nit" type="text" class="form-control" placeholder="Introduce Nit/CC/ID   {{if .Errors.Nit}}(***!!!{{.Errors.Nit}}!!!***){{end}}"  tabindex="1" />
          </div>
    
          <div  class="form-group col-xs-10 col-md-8">
            <label class="sr-only"  for="empresa">Empresa</label>  
            <input name="Empresa" type="text"  class="form-control" placeholder="Nombre Empresa  {{if .Errors.Empresa}}(***!!!{{.Errors.Empresa}}!!!***){{end}}" required tabindex="2" />
          </div>
        </div>

        <div>
          <div  class="form-group col-xs-10 col-md-6">
            <label class="sr-only"  for="contacto">Contacto</label>  
            <input name="Contacto" type="text"  class="form-control" placeholder="Contacto "  tabindex="3" />
          </div>
          
          <div  class="form-group col-xs-10 col-md-6">
            <label class="sr-only"  for="email">Email</label>  
            <input name="Email" type="email"  class="form-control" placeholder="Email "  tabindex="4" />
          </div>
        </div>

        <div>
          <div class="form-group col-xs-10 col-md-4">
            <label class="sr-only"  for="pais">Pais</label>  
            <input name="Pais" type="text"  class="form-control" placeholder="Pais " required  tabindex="7" />
          </div>
          <div class="form-group col-xs-10 col-md-4">
            <label class="sr-only"  for="ciudad">Ciudad</label>  
            <input name="Ciudad" type="text"  class="form-control" placeholder="Ciudad "  required tabindex="8" />
          </div>
          <div class="form-group col-xs-10 col-md-4">
            <label class="sr-only"  for="Tel">Tel</label>  
            <input name="Tel" type="text"  class="form-control" placeholder="Tel "  tabindex="6" />
          </div>
        </div>

        <div>
          <div class="form-group col-xs-10 col-md-5">
            <label class="sr-only"  for="url">Url</label>  
            <input name="Url" type="text"  class="form-control" placeholder="Url "  tabindex="5" />
          </div>
          <div class="form-group col-xs-10 col-md-7">
            <label class="sr-only"  for="direccion">Direccion</label>  
            <input name="Direccion" type="text"  class="form-control" placeholder="Direccion "  tabindex="9" />
          </div>
        </div>

        <div class="form-group col-xs-12 col-md-12">
          <label class="sr-only"  for="nota">Nota</label>  
          <input name="Nota" type="text"  class="form-control" placeholder="Nota "  tabindex="10" />
        </div>
        
        <input type="submit" form="proveedor"value="Crear Proveedor" class="btn btn-default" tabindex="11" /> &nbsp;
        <input type="reset" form="proveedor"value="Cancelar" class="btn btn-default" tabindex="12" />
           
      </form>

  </div>
</div>

 <div class="container">
      <div><a href="/prov"><strong>Regresar a Provedores</strong></a></div>
  </div>
