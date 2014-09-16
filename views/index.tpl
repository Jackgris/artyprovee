<div class="container">
    <div class="hero-text">
      <div>
        <div class=" text-center form-group col-xs-12 col-md-12">
          <img src="static/images/logo_body.jpg" >
        </div>
      </div>
      <div>
        <div class="page-header" class="  form-group col-xs-12 col-md-12">    
            <h1 class="text-center"><small>¡¡ Organiza tu informacion de Articulos y Porveedores !!</small></h1>
        </div>
      </div>   
    </div>
</div> 

<div class="container">
  <div class="  form-group col-xs-12 col-md-12">    
    <h3 class="text-center">Inicio de Sesion</h3>
    <h4 class="text-center">{{.mensaje}}</h4>
  </div>
</div>
  
<div class="container" >
  <div class="col-xs-3" ></div>
  <div class="col-xs-6" >
    <div class="container">
      
      <form  class="form-inline" role="form" id="login" method="POST" >
        <div class="form-group">
            <label class="sr-only" for="usuario">Usuario</label>  
            <p><input name="usuario" type="text" min="1" max="100" autocomplete="off" class="form-control" placeholder="Usuario"  required tabindex="1" autofocus /></p>
        </div>
    
        <div class="form-group" >
            <label class="sr-only" for="password">Contraseña</label>  
            <p><input name="password" type="password"  class="form-control" placeholder="Contraseña" required tabindex="2" /></p>
        </div>
          
        <div >
        {{create_captcha}}
        </div>
        <div>
            <p><input name="captcha" type="text" class="form-control" placeholder="Texto imagen" required tabindex="3"></p>
        </div>
        
        <div >
          <p><input name="login" type="submit"  class="btn btn-default" tabindex="4" {{.valor}}/></p>&nbsp;      
        </div>  
      </form>
    </div>          
  </div>
  <div class="col-xs-3" ><p></p></div>
</div>

