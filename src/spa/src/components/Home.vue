<template>
    <div id="contenedor">
        <div id="info-cliente">
            Work in Progress
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
            ordenes: {}
        }
    },
    methods: {
        obtenerDatos() {
            axios
                .get("http://localhost:10040/clientes")
                .then(response => {
                    const dataClientes = response.data.data
                    this.clientes = dataClientes
                })
        }
    }
}
</script>

<style>
#contenedor {
    position: absolute;
    top: 50%;
    left: 50%;
}
</style>