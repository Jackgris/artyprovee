<div class="container">
    <div class="hero-text">
      <div>
        <div class=" text-center form-group col-xs-12 col-md-12">
          <img src="static/images/logo_body.jpg" >
        </div>
      </div>
      <div>
        <div class="page-header" class="  form-group col-xs-12 col-md-12">    
            <h3 class="text-left">Datos de la Empresa</h3>
              <p>Establesca los datos de su empresa, estos se utilizaran para algunos encabezados.</p>
        </div>
      </div>   
    </div>
</div>

<div class="container" >
      <form class="form-horizontal " role="form" id="empresa" method="POST" >
        <div class="form-group">
            <label class="col-sm-2 control-label" for="Empresa">Empresa</label>  
            <div class="col-xs-4 col-md-6">
            <p><input name="Empresa" type="text" min="1" max="100" autocomplete="off" class="form-control" value="{{.records.Empresa}}"  required tabindex="1" autofocus /></p>
          </div>    
        </div>

        <div class="form-group">
            <label class="col-sm-2 control-label" for="Representante">Representante</label>  
            <div class="col-xs-4 col-md-6">
            <p><input name="Representante" type="text" min="1" max="100" autocomplete="off" class="form-control" value="{{.records.Representante}}"  required tabindex="1" autofocus /></p>
          </div>    
        </div>

        <div class="form-group">
            <label class="col-sm-2 control-label" for="Nit/Dni">Nit/Dni</label>
            <div class="col-xs-4 col-md-6">  
            <p><input name="NitDni" type="text" min="1" max="100" autocomplete="off" class="form-control" value="{{.records.NitDni}}" required tabindex="1" autofocus /></p>
          </div>
        </div>
    
        <div class="form-group" >
            <label class="col-sm-2 control-label" for="Direccion">Direccion</label>
            <div class="col-xs-4 col-md-6">  
            <p><input name="Direccion" type="Direccion"  class="form-control" value="{{.records.Direccion}}" required tabindex="2" /></p>
          </div>
        </div>  
        
        <div class="form-group">
            <label class="col-sm-2 control-label" for="Telefono">Telefono</label>
            <div class="col-xs-4 col-md-6">  
            <p><input name="Telefono" type="text" min="1" max="100" autocomplete="off" class="form-control"  value="{{.records.Telefono}}" required tabindex="1" autofocus /></p>
          </div>
        </div>
        
         <div class="form-group">
            <label class="col-sm-2 control-label" for="Ciudad">Ciudad</label> 
            <div class="col-xs-4 col-md-6"> 
            <p><input name="Ciudad" type="text" min="1" max="100" autocomplete="off" class="form-control" value="{{.records.Ciudad}}" required tabindex="1" autofocus /></p>
          </div>
        </div>  

         <div class="form-group">
            <label class="col-sm-2 control-label" for="Pais">Pais</label>
            <div class="col-xs-4 col-md-6">  
            <p><input name="Pais" type="text" min="1" max="100" autocomplete="off" class="form-control"  value="{{.records.Pais}}" required tabindex="1" autofocus /></p>
          </div>
        </div>

        <p></p>
     <div class="form-group col-xs-6 col-md-6">
       <div class="alert alert alert-danger">
         <strong>Tenga Presente!</strong> Los cambios son IRREVERSIBLES!!.
       </div>
       <div class="form-group has-warning" >
         <label for="nombre" class="control-label" title="!!LOS CAMBIOS SON IRREVERSIBLES!!">ACCIONES SOBRE EL REGISTRO </label>
          <ul>
           <li>
             <input type="radio" name="accion"  title="!!LOS CAMBIOS SON IRREVERSIBLES!!" value="up">Editarlo
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
