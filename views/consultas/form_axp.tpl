<form class="form-horizontal col-sm-8" role="form" method="POST">
  <div >
    <div class="form-group col-xs-6 col-md-3">
    	<label for="proveedor"class="control-label">Proveedor</label>
    </div>
    <div class="form-group col-xs-6 col-md-6">
      <input name="IdProveedor" type="text" class="form-control" id="id_proveedor" placeholder=""  list="opciones1">
      <datalist id="opciones1">
              {{range $prov := .proveedor}}
              <option>{{$prov.Id}} {{$prov.Empresa}}</option>
              {{end}}
      </datalist>
    </div>
 	<div class="form-group col-xs-6 col-md-3">
     <button type="submit" class="btn btn-default">Enviar</button>
    </div>
  </div>
</form>