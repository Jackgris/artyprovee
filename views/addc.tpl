<div class="container">
  <div class="row">
   <p>
    <div class="hero-text">        
     <h1 >Nueva Categoria</h1>
      {{if .flash.error}}
      <blockquote class="col-xs-8">{{.flash.error}}</blockquote>
      {{end}}

      {{if .flash.notice}}
      <blockquote >{{.flash.notice}}</blockquote>
      {{end}}
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
    
    
    <p>          
      <form class="col-xs-8" id="categoria" method="POST">
           
        <div class="form-group has-feedback">
          <label for="cod_categoria">Cod. Categoria </label>  
          <input name="CodCategoria" type="number" min="1" max="100" autocomplete="off" class="form-control" placeholder="Introduce Cod_Categoria   {{if .Errors.CodCategoria}}(***!!!{{.Errors.CodCategoria}}!!!***){{end}}"  tabindex="1" />
        </div>
  
        <div class="form-group">
          <label for="empresa">Categoria</label>  
          <input name="Categoria" type="text"  class="form-control" placeholder="Nombre Categoria  {{if .Errors.Categoria}}(***!!!{{.Errors.Categoria}}!!!***){{end}}" required tabindex="2" />
        </div>
              
        <input type="submit" form="categoria"value="Crear Categoria" class="btn btn-default" tabindex="11" /> &nbsp;
        <input type="reset" form="categoria"value="Cancelar" class="btn btn-default" tabindex="12" />
           
      </form>
    </p>
  </div>
</div>