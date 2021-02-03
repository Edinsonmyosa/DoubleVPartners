# Tickets

_API que permite crear,eliminar,editar y recuperar tickets por ID o todos_

### Instalación

_En caso de estar corriendo el docker será necesario detener el mismo_

```
sudo docker stop prueba
```

_Luego eliminarlo_

```
sudo docker rm prueba
```

_Borrar la imagen_

```
sudo docker image rm prueba
```

_Una vez se tengan los campos requeridos se puede construir la imagen_

```
sudo docker build . -t prueba
```

_Finalmente ejecutar la misma, en este caso asignando una zona Horaria y haciendo el puerto accesible_

```
sudo docker run -d --net="host" --restart=always -e TZ=America/Bogota --name prueba prueba
```
