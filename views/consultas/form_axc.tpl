<form class="form-horizontal col-sm-8" role="form" method="POST">
  <div >
    <div class="form-group col-xs-6 col-md-3">
    	  <label for="categoria"class="control-label">Categoria</label>
    </div>
    <div class="form-group col-xs-6 col-md-6">
      <input name="IdCategoria" type="text" class="form-control" id="id_categoria" placeholder=""  list="opciones2"> 
      <datalist id="opciones2">
              {{range $categ := .categoria}}
              <option>{{$categ.CodCategoria}} {{$categ.Categoria}}</option>
              {{end}}
      </datalist>
    </div>
 	<div class="form-group col-xs-6 col-md-3">
     <button type="submit" class="btn btn-default">Enviar</button>
    </div>
  </div>
</form>