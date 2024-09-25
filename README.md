# :globe_with_meridians: OIKOS API

API CRUD para la gestión de dependencias y espacios físicos dentro de la Universidad Distrital

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# Ejemplo que se debe actualizar acorde al proyecto
OIKOS_API__ENABLEDOCS = [descripción]
OIKOS_API__HTTPPORT = [descripción]
OIKOS_API__RUNMODE = [descripción]
OIKOS_API__PGSCHEMA = [descripción]
OIKOS_API__PGDB = [descripción]
OIKOS_API__PGPASS = [descripción]
OIKOS_API__PGURLS = [descripción]
OIKOS_API__PGUSER = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con OIKOS_API__...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/oikos_api

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/oikos_api

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
OIKOS_API__HTTPPORT=8080 OIKOS_API__PGURLS=127.0.0.1:27017 OIKOS_API__SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=oikos_api . --no-cache
# docker run -p 80:80 oikos_api
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/oikos_api

#2. Moverse a la carpeta del repositorio
cd oikos_api

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```
### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/oikos_api/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/oikos_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/oikos_api/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/oikos_api) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/oikos_api/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/oikos_api) |


## Modelo de Datos

[SVG](database/oikos.svg) -
[PGmodeler](database/oikos.dbm)

## Licencia

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with oikos_api. If not, see https://www.gnu.org/licenses/.

### UNIVERSIDAD DISTRITAL FRANCISCO JOSÉ DE CALDAS
### OFICINA ASESORA DE SISTEMAS
### 2019
