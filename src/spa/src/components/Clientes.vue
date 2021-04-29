<template>
    <div class="component-div">
        <h2 class="header-text">Clientes</h2>
        <div id="clientes-header">
            <input 
                v-model="search" 
                id="barra-busqueda" 
                placeholder="Ingrese el nombre a buscar">
            <button @click="showModal" class="btnAcciones btnCrear">Crear cliente</button>
        </div>
        <CrearCliente 
            v-bind:cliente="seleccion"
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
                <tr v-for="cliente in filtrarClientes" :key="cliente.ID">
                    <td>{{ cliente.nombre }}</td>
                    <td>{{ cliente.primerApellido }}</td>
                    <td>{{ cliente.segundoApellido }}</td>
                    <td>{{ cliente.domicilio }}</td>
                    <td>{{ cliente.ciudad }}</td>
                    <td>{{ cliente.entidadFederativa }}</td>
                    <td>{{ cliente.telefono }}</td>
                    <td>{{ cliente.email }}</td>
                    <td colspan="2">
                        <button v-on:click='editarCliente(cliente.ID)' class="btnAcciones btnEditar">Editar</button>
                        <button v-on:click='eliminarCliente(cliente.ID)' class="btnAcciones btnEliminar">Eliminar</button>
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
    import CrearCliente from './clientes/CrearCliente'
    const axios = require('axios')

    export default {
        name: 'Clientes',
        title: 'Clientes',
        data() {
            return {
                info: {},
                search: "",
                seleccion: "",
                isModalVisible: false
            }
        },
        components: {
            CrearCliente,
        },
        mounted() {
            this.obtenerClientes()
        },
        computed: {
            filtrarClientes() {
                // Basado en https://stackoverflow.com/questions/62711939/filtering-table-on-the-frontend-vue-js
                let clientesFiltrados = this.info
                if (this.search.length != 0) { 
                    clientesFiltrados = clientesFiltrados.filter( cliente => {
                        return cliente.nombre.search(this.search) != -1
                    })
                }

                return clientesFiltrados
            }
        },
        methods: {
            showModal() {
                this.isModalVisible = true
            },
            closeModal() {
                this.isModalVisible = false
            },
            obtenerClientes() {
                axios
                    .get('http://localhost:10040/clientes')
                    .then(response => this.info = response.data.data)
                    .catch(error => console.log(error.response.data))
            },
            editarCliente: function (ID) {
                this.seleccion = ID.toString()
                this.isModalVisible = true
            },
            eliminarCliente: function (ID) {
                var r = confirm("¿Estás seguro de eliminar este registro?")
                if (r == true) {
                    axios
                        .delete('http://localhost:10040/clientes/' + ID)
                        .then(() => {
                            // Borrar el cliente localmente para refrescar
                            // https://stackoverflow.com/questions/53142220/auto-reload-list-items-after-deletion-vue-js
                            const cliente = this.info.data.data.findIndex(p => p.id === ID)
                            this.info.data.data.splice(cliente, 1)
                        })
                        .catch(error => console.log(error.response.data))
                }
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
        margin-right: 20px;
        width: 300px;
    }

    #clientes-tabla {
        grid-area: clientes-tabla;
        width: 100%;
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

    .btnAcciones {
        padding: 7px;
        font-weight: bold;
        color: white;
    }

    .btnCrear {
        background: lightblue;
    }

    .btnEditar {
        background: lightgreen;
    }

    .btnEliminar {
        background: lightcoral;
    }

</style>