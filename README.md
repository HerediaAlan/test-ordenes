# Vue.js + Go REST API
Proyecto integrador de una API con backend en Go, y vistas usando Vue.js + Bootstrap-Vue

## Tecnologías utilizadas
* Vue.js (frontend)
* Go (backend)
* Bootstrap-Vue (framework de soporte para frontend)
* Gin (framework de backend para APIs)
* Gorm (ORM de GO)
* Axios.js (Cliente HTTP)
* Moment.js (Framework dinámico para fechas en JS)
* GoFakeIt (creador de datos falsos para propósitos de _testing_)

## Diagrama Entidad-Relación
![alt text](https://github.com/AlanHerediaG/test-ordenes/raw/main/docs/diagrama-basico-er.png "Diagrama")

## Rutas disponibles
### Clientes
* GET _/clientes_
* GET _/clientes/:ID_
* POST _/clientes_
* PUT _/clientes/:ID_
* DELETE _/clientes/:ID_

### Ordenes
* GET _/ordenes_
* GET _/ordenes/:ID_
* POST _/ordenes_
* PUT _/ordenes/:ID_
* DELETE _/ordenes/:ID_
* GET _/ordenes/:ID/comentarios_
* POST _/ordenes/:ID/comentarios_