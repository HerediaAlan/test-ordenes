<template>
    <div class="component-div">
        <h2 class="header-text">Clientes</h2>
        <div id="clientes-header">
            <input type="text" v-model="search" id="barra-busqueda" placeholder="Ingrese nombre a buscar">
            <button @click="showModal">Crear cliente</button>
        </div>
        <CrearCliente 
            v-show="isModalVisible"
            @close="closeModal"/>
        <table id="clientes-tabla">
            <thead>
                <tr>
                    <th>Nombre</th>
                    <th>Primer Apellido</th>
                    <th>Segundo Apellido</th>
                    <th>Domicilio</th>
                    <th>Ciudad</th>
                    <th>Entidad Federativa</th>
                    <th>Telefono</th>
                    <th>Email</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in info.data.data" :key="item.ID">
                    <td>{{ item.nombre }}</td>
                    <td>{{ item.primerApellido }}</td>
                    <td>{{ item.segundoApellido }}</td>
                    <td>{{ item.domicilio }}</td>
                    <td>{{ item.ciudad }}</td>
                    <td>{{ item.entidadFederativa }}</td>
                    <td>{{ item.telefono }}</td>
                    <td>{{ item.email }}</td>
                    <td colspan="2">
                        <button v-on:click='editarCliente(item.ID)' class="btnEditar">Editar</button>
                        <button v-on:click='eliminarCliente(item.ID)' class="btnEliminar">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    // https://router.vuejs.org/guide/advanced/data-fetching.html#fetching-after-navigation
    // https://github.com/axios/axios
    // https://vuejs.org/v2/cookbook/using-axios-to-consume-apis.html
    import { mdiAccount } from '@mdi/js';
    import CrearCliente from './clientes/CrearCliente'
    const axios = require('axios')

    export default {
        name: 'Clientes',
        title: 'Clientes',
        icon: mdiAccount,
        data() {
            return {
                info: {},
                isModalVisible: false
            }
        },
        components: {
            CrearCliente,
        },
        mounted() {
            axios
                .get('http://localhost:10040/clientes')
                .then(response => {
                    this.info = response
                })
        },
        methods: {
            showModal() {
                this.isModalVisible = true
            },
            closeModal() {
                this.isModalVisible = false
            },
            editarCliente: function (ID) {
                alert(ID)
            },
            eliminarCliente: function (ID) {
                alert(ID)
            }
        }
    }
</script>

<style scoped>
    .componentView {
        display: grid;
        grid-template-rows: 60px 60px auto;
        grid-template-areas:
            "header-text"
            "barra-busqueda"
            "clientes-tabla";    
    }

    .header-text {
        grid-area: header-text;
        margin-top: 15px;
        margin-bottom: 15px;
    }

    #barra-busqueda {
        grid-area: barra-busqueda;
        padding: 4px 12px;
        color: rgba(0, 0, 0, .70);
        border: 1px solid rgba(0, 0, 0, .12);
        background: white;
        margin-bottom: 15px;
    }

    #clientes-tabla {
        grid-area: clientes-tabla;
    }

    th {
        text-align: left;
        padding-top: 7px;
        padding-bottom: 7px;
        padding-right: 5px;
    }

    td {
        text-align: left;
        font-size: 14px;
        padding-top: 5px;
        padding-bottom: 5px;
    }

    .btnEditar {
        background: lightgreen;
    }

    .btnEliminar {
        background: lightcoral;
    }

</style>