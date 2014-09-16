
<div class="container">
    <div class="hero-text">
      <div>
        <div class=" text-center form-group col-xs-12 col-md-12">
          <img src="static/images/logo_body.jpg" >
          {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
        </div>
      </div>
      <div>
        <div class="page-header" class="  form-group col-xs-12 col-md-12">    
            <h3 class="text-left">Administracion de Usuarios</h3>
              <p>Las cuentas de Usuario, deben de pertenecer a un grupo; los grupos preestablecidos son: Administradores, Usuarios, Clientes. </p>
        </div>
      </div>   
    </div>
</div>


<div class="container">
  <div >    
    <h4 class="text-left">Usuarios Creados</h4>
              
  </div>
  <div  class="col-xs-12 col-md-7" class="row table-responsive">
    <table class="table table-striped table-hover table-condensed" >
      <thead>
        <tr>
          <th>Usuario</th>
          <th>Grupo</th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <form method="POST">
            <td><input name ="Usuario" class="sr-only" for="" value="{{$record.Usuario}}">{{$record.Usuario}}</td>
            <td>{{$record.Grupo}}</td>
            <td>
              <button name="Eliminar" type="submit" value="ok" >Eliminar</button>
            </td>
          </form>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="3"></td>
        </tr>
      </tfoot>
    </table>
  </div>
</div>

<hr />

<div class="container" >
<div >    
    <h4 class="text-left">Crear Usuario</h4>
              
  </div>
<form class="form-horizontal col-xs-12 col-md-6" role="form" id="user" method="POST">
  <div class="form-group">
    <label class="sr-only" for="Usuario" class="col-sm-2 control-label">Usuario</label>
    <div class="col-sm-10">
      <input name="Usuario" type="text" class="form-control"  placeholder="Usuario, entre 5 y 12 caracteres" pattern="[a-zA-Z0-9]{5,}" maxlength="12" required tabindex="1" autofocus>
    </div>
  </div>
  <div class="form-group">
    <label class="sr-only" for="Grupo" class="col-sm-2 control-label">Grupo</label>
    <div class="col-sm-10">
      <input name="Grupo" type="text" list="lista" class="form-control" placeholder="Grupo, utiliza el boton" required tabindex="1">
        <datalist id="lista">
          <option>Administrador</option>
          <option>Usuario</option>
          <option>Cliente</option>
        </datalist>
    </div>
  </div>
  <div class="form-group">
    <label class="sr-only" for="Passwd" class="col-sm-2 control-label">Password</label>
    <div class="col-sm-10">
      <input name="Passwd" type="Passwd" class="form-control" placeholder="Passwd, entre 5 y 12 caracteres" pattern="[a-zA-Z0-9]{5,}" maxlength="12" required tabindex="2">
    </div>
  </div>
 
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button name ="Adicionar" type="submit" class="btn btn-default" value="add"tabindex="4">Enviar</button>
    </div>
  </div>
</form>
</div>