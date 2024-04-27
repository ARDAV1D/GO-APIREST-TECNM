# GO-APIREST-TECNM ALVARADO DAVID
APIREST creada con Golang.
Se creo la base de datos en un contenedor docker postgreSQL, la API se creo con el lenguaje Go, se utilizo un archivo Dockerfile para crear la imagen de la API para poder asignar la imagen a un contenedor y poder comunicarse con el contenedor de la base de datos, se creo el archivo DockerCompose.yml para definir y ejecutar multi-contenedores y asignarlos a red docker para que se puedan comunicar entre si. 

En la base de datos se crearon 4 tablas
especialidades, licenciaturas, materia, reticula. 
En la tabla de licenciatura se crearon 2 columnas, nombre y una llave foranea con especialidad.
En la tabla de especialidades se crearon 3 columnas, nombre_Esp, licenciatura_id que es la llave foranea de licenciaturas y materias que es la llave foranea hacia la tabla de materia
En la tabla de materia se crearon 3 columnas, nombre_Ma, especialidad_id que es la llave foranea hacia especialidad y reticulas que es la llave foranea de MateriaID a la tabla de reticula.
y por ultimo en la tabla de reticula se crearon 11 columnas, especialidad_id que es la llave foranea de especialidad, semestre_1 hasta el 9 y materias_id que es la llave foranea de la tabla materia.

#COMANDOS PARA IMAGENES DOCKERHUB 
#APITECNM
docker pull ardav1d/apitecnm

#POSTGRES
docker pull ardav1d/postgres
