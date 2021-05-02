<template>
    <div id="contenedor">
        <div id="info-cliente">
            <h2>Información sobre clientes</h2>

            <span class="font-weight-bold">Cantidad de clientes existentes:</span>
            <p>{{ this.clientes.length }}</p>

            <span class="font-weight-bold">Top 10 clientes con mayor cantidad de ordenes de soporte:</span>
            <b-table :items="ordenarClientes" :fields="clientesColumnas" small></b-table>

            <span class="font-weight-bold">Cantidad de clientes por entidad federativa:</span>
            <b-table :items="filtrarPorEntidad" :fields="clientesColumnasEntidad" small></b-table>
        </div>
        <hr style="height: 2px; background-color: black;"/>
        <div id="info-ordenes">
            <h2>Información sobre ordenes</h2>

            <span class="font-weight-bold">Cantidad de ordenes existentes:</span>
            <p>{{ this.ordenes.length }}</p>

            <span class="font-weight-bold">Promedio de tiempo de atención:</span>
            <p>{{ this.calcularPromedioAtencion }} horas</p>
        </div>
    </div>
</template>

<script> 
const axios = require("axios");
const moment = require("moment");

export default {
    name: "home",
    data() {
        return {
            clientes: {},
            clientesColumnas: ["nombre", "primerApellido", "segundoApellido", "ciudad", "cantidad"],
            clientesColumnasEntidad: ["entidadFederativa", "cantidad"],
            ordenes: {}
        }
    },
    computed: {
        ordenarClientes() {
            let datos = this.clientes
            let orden = datos.sort((a, b) => b.Ordenes.length - a.Ordenes.length).slice(0, 10)
            return orden
        },
        filtrarPorEntidad() {
            let datos = this.clientes
            // Este snippet fué obtenido de: https://stackoverflow.com/a/46794337
            const result = [...datos.reduce((r, o) => {
                const key = o.entidadFederativa;
                
                const item = r.get(key) || Object.assign({}, o, {
                    entidadFederativa: "",
                    cantidad: 0
                });
                
                item.entidadFederativa = o.entidadFederativa;
                item.cantidad += o.entidadFederativa === item.entidadFederativa ? 1 : 0;

                return r.set(key, item);
                }, new Map).values()];
            
            return result.sort((a, b) => b.cantidad - a.cantidad)
        },
        calcularPromedioAtencion() {
            // Este método retorna el promedio en horas del tiempo de 
            // respuesta entre orden y último comentario de soporte.
            // Función obtenida de: https://stackoverflow.com/a/45609104
            var sum = 0;
            this.ordenes.forEach((orden) => {
                if (orden.OrdenComentarios.length > 0) {
                    var fechaOrden = moment(orden.FechaCreacion);

                    var comentario = orden.OrdenComentarios.slice(-1)[0];
                    const fechaComentario = moment(new Date(comentario.FechaCreacion).toUTCString());

                    var diff = moment.duration(fechaComentario.diff(fechaOrden));
                    sum += diff.as('hours');
                }
            });

            var avg = sum / (this.ordenes.length - 1);

            return avg.toFixed(4);
        }
    },
    beforeMount() {
        this.obtenerDatos();
        this.obtenerOrdenes();
    },
    methods: {
        obtenerDatos() {
            axios
                .get("http://localhost:10040/clientes")
                .then(response => {
                    const dataClientes = response.data.data
                    dataClientes.forEach(element => {
                        element.cantidad = element.Ordenes.length
                    });
                    this.clientes = dataClientes
                })
                .catch((error) => console.log(error.response.data));
        },
        obtenerOrdenes() {
            axios
                .get("http://localhost:10040/ordenes")
                .then((response) => {
                    const dataOrdenes = response.data.data;
                    dataOrdenes.forEach((element) => {
                        element.FechaCreacion = new Date(element.FechaCreacion).toUTCString()
                    })
                    this.ordenes = dataOrdenes;
                })
                .catch((error) => console.log(error.response.data));
        },
    }
}
</script>

<style>
</style>