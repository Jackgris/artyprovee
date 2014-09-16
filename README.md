artyprovee
==========

Esta es una Aplicacion Web realizada mediante el uso del Lenguaje de Programacion "GO". Para la integracion de esto al marco Web, se uso el Framework Beego.

El objetivo ha sido el "aprendizaje de go y el framework", esperando se entienda esto. **Invito a la comunidad** a que revise el codigo para corregir lo que crean conveniente, lo cual agradeceria, ya que permitiria avanzar en el aprendizaje de los mismos, ademas ser un modelo de referencia.

Agradezco a los desarrolladores de Beego por su excelente trabajo, ya que me ha sido muy comodo "empezar" a utilizar este gran lenguaje de programacion, y poder entender y tratar de hacer una aplicacion web con su ayuda. Tambien agradezco a Jan Schlicht, ya que he usado "github.com/nfnt/resize" para el manejo de las imagenes, al amigo @Jackgris por ayudarme a revisar la aplicacion y a @yavallejo por mostrarme el uso basico de css.


La aplicacion permite crear articulos, categorias y proveedores de los mismos. Los datos se guardan en Mysql.

Estan creadas algunas consultas sobre los datos al igual que el manejo de edicion y borrado de registros.

He intentado manejar la seguridad usando el registro de los usuarios mediante: usuario, contraseña y el capcha.  Segun el nivel que se le de al usuario "Administrador, usuario o cliente", se puede restringir el acceso a algunos modulos.

En cuanto a las imagenes, permite subir imagenes en formato .jpg haciendoles un "resize" para anexarlas al registro del articulo.

Hay varias cosas pendientes las cuales no he podido realizar, algunas por no saber aun como hacerlas (impresion, exportar informes a un formato de archivo, bloqueo de reenvio del formulario, entre otras..).

