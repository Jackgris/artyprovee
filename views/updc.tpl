<div class="container">
  <div class="row">
    <p>
    <div class="hero-text">        
     <h2>Detalles Categoria: {{.Id}} </h2>
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
  <div class="row">       
   <form id="categoria" method="POST">
    <div class="form-group col-xs-6 col-md-8" >
      <div >
          <label for="cod_categoria">Cod. Categoria</label>  
          <input name="CodCategoria" type="text"  class="form-control" value="{{.records.CodCategoria}}"  tabindex="5" />
      </div>
      <div >
          <label for="cod_categoria">Categoria</label>  
          <input name="Categoria" type="text"  class="form-control" value="{{.records.Categoria}}"  tabindex="5" />
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