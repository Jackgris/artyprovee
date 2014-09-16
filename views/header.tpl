
<!DOCTYPE html>
<html lang="es">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title></title>
    

 
 <link rel="stylesheet" href="/static/css/bootstrap.min.css">
 <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
 <link rel="stylesheet" href="/static/css/starter-template.css" >
<style>
  body { padding-top: 70px; }
  .miclase{
    width: 118px;
    padding: 0 15px;
  }
 </style>
 
    

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  

  <body>
  <nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
    <div  class="container">
      <div class="navbar-header">
        <a class="navbar-brand miclase" href="#">  
              <img class="img-responsive" src="/static/images/logo_header.png" alt="No Hay Foto">
        </a>
        <button button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbar-collapse-1">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        
      </div>
      <div class="collapse navbar-collapse" id="navbar-collapse-1">
        <ul class="nav navbar-nav navbar-left">
          <li><a href=" {{urlfor "Apcontroller.Home"}}"><span class="glyphicon glyphicon-home"></span>Home</a></li>
          <li class="dropdown">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">Manejo de Registros <b class="caret"></b></a>
            <ul class="dropdown-menu" >
              <li class="dropdown-header"><span class="glyphicon glyphicon-list"></span> <strong> Acciones</strong></li>
             
              <li><a href="/adda">Nuevo articulo</a></li>
              <li><a href="/addc">Nueva categoria</a></li>
              <li><a href="/addp">Nuevo proveedor</a></li>
              <li class="divider"></li>

              <li><a href="/cons">Consultas</a></li>
            </ul>
          </li>
        </ul>
       {{if .IsArtProv}}{{template  "/nav_art.tpl".}}{{end}}
        {{if .IsCons}}{{template  "/nav_cons.tpl".}}{{end}}
        <ul class="nav navbar-nav navbar-right">
         <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">
          <span class="glyphicon glyphicon-cog"></span> <b class="caret"></b>
          </a>
           <ul class="dropdown-menu">
            <li><a href="/">Cerrar Sesi&oacute;n</a></li>
            <li class="divider"></li>
            <li><a href="/empresa" >Datos Empresa</a></li>
            <li><a href="/user">Usuarios</a></li>
            <li><a href="#">Contacto</a></li>
           </ul>
         </li>
        </ul>
      </div><!--/.nav-collapse -->
    </div>

  </nav>
