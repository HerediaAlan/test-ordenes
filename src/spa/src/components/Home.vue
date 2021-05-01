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
        <div id="info-ordenes"></div>
    </div>
</template>

<script> 
const axios = require("axios");

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
        }
    },
    beforeMount() {
        this.obtenerDatos();
        this.obtenerOrdenes();
    },
    methods: {
        groupBy(datos) {
            const entidades = datos.map(x => x.entidadFederativa);
            return entidades.unique();
        },
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
                    this.ordenes = dataOrdenes;
                })
                .catch((error) => console.log(error.response.data));
        }
    }
}
</script>

<style>
</style>