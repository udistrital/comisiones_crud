# comisiones_crud
:heavy_check_mark: Check: Repositorio API CRUD para el sistema de gestión de comisiones.


## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang 1.25](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo 1.12.3](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)


### Variables de Entorno
```shell
# Parámetros de la API
COMISIONES_CRUD_HTTPPORT=[Puerto de exposición de la API]
COMISIONES_CRUD_RUNMODE=[Modo de ejecución: dev o prod]

# Parámetros de base de datos
COMISIONES_CRUD_DB_USER=[Usuario de la base de datos]
COMISIONES_CRUD_DB_PASS=[Contraseña del usuario de la base de datos]
COMISIONES_CRUD_DB_URL=[URL, dominio o endpoint de la base de datos]
COMISIONES_CRUD_DB_PORT=[Puerto de la base de datos]
COMISIONES_CRUD_DB_NAME=[Nombre de la base de datos]
COMISIONES_CRUD_DB_SCHEMA=[Nombre del esquema de la base de datos]

```

**NOTA:** Las variables se pueden consultar en el archivo `conf/app.conf` y están identificadas principalmente con el prefijo `COMISIONES_CRUD_`.


### Ejecución del Proyecto
```shell
# 1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/comisiones_crud.git

# 2. Moverse a la carpeta del repositorio
cd comisiones_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. Configurar las variables de entorno
export COMISIONES_CRUD_HTTPPORT=8080

# 5. Ejecutar el proyecto
bee run
```

### Ejecución Dockerfile
```shell
# El Dockerfile está implementado para el despliegue mediante
# el sistema de integración continua (CI).

# 1. Compilar el proyecto para Linux
GOOS=linux GOARCH=amd64 go build -o main

# 2. Construir la imagen
docker build -t comisiones_crud .

# 3. Ejecutar el contenedor
docker run --name comisiones_crud \
  --env-file custom.env \
  -p 8080:8080 \
  comisiones_crud

# 4. Comprobar que el contenedor esté en ejecución
docker ps
```


### Ejecución docker-compose
```shell
# No implementado actualmente.
```

## Estado CI

| Develop | release/0.0.1 | Master | Sonar |
| -- | -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/comisiones_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/comisiones_crud)| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/comisiones_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/comisiones_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/comisiones_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/comisiones_crud) | [![Quality Gate Status](https://sonarqube.portaloas.udistrital.edu.co/api/project_badges/measure?project=comisiones_crud&metric=alert_status)](https://sonar.portaloas.udistrital.edu.co/dashboard?id=comisiones_crud) |


## Modelo de datos

[Modelo de datos](database/modelos.svg)



## Licencia

This file is part of comisiones_crud.

comisiones_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

comisiones_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with comisiones_crud. If not, see https://www.gnu.org/licenses/.
