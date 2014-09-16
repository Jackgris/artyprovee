<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Listado de Imagenes</h1>   
    </div>
  </div>
</div> 

<div class="container">
  
<div class="row">
  {{range $record := .prueba}}
  <div class="col-sm-6 col-md-3">
    <div class="thumbnail">
      <img src="{{$record.Foto1}}" alt="No Hay Foto" width="45%">
      <div class="caption">
        <h5></h5>
        <p><a href="upda/{{$record.Id}}" >{{$record.Descripcion}}</a> </p>
      </div>
    </div>
  </div>
  
 {{end}}

</div>
</div>