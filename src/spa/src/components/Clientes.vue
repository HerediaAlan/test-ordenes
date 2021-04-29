<template>
    <div class="component-div">
        <h2 class="header-text">Clientes</h2>
        <div id="clientes-header">
            <input
                v-model="search"
                id="barra-busqueda"
                placeholder="Ingrese el nombre a buscar"
            />
            <b-button
                @click="showModal"
                variant="primary"
                size="sm"
            >Añadir cliente</b-button>
        </div>
        <CrearCliente
            v-bind:cliente="seleccion"
            v-show="isModalVisible"
            @close="closeModal"
        />
        <b-pagination
            v-model="currentPage"
            :total-rows="rows"
            :per-page="perPage"
            aria-controls="clientes-tabla"
        ></b-pagination>
        <b-table 
            id="clientes-tabla"
            :items="filtrarClientes" 
            :per-page="perPage"
            :current-page="currentPage"
            :fields="columnas"
            striped hover small
        >
        </b-table>
         <!-- 
             <button
                            v-on:click="editarCliente(cliente.ID)"
                            class="btnAcciones btnEditar"
                        >
                            Editar
                        </button>
                        <button
                            v-on:click="eliminarCliente(cliente.ID)"
                            class="btnAcciones btnEliminar"
                        >
                            Eliminar
                        </button>
         -->
    </div>
</template>

<script>
// https://router.vuejs.org/guide/advanced/data-fetching.html#fetching-after-navigation
// https://github.com/axios/axios
// https://vuejs.org/v2/cookbook/using-axios-to-consume-apis.html
import CrearCliente from "./clientes/CrearCliente";
const axios = require("axios");

export default {
    name: "Clientes",
    title: "Clientes",
    data() {
        return {
            currentPage: 1,
            perPage: 10,
            info: {},
            columnas: ["nombre", "primerApellido", "segundoApellido", "domicilio", "ciudad", "entidadFederativa", "telefono", "email"],
            search: "",
            seleccion: "",
            isModalVisible: false,
        };
    },
    components: {
        CrearCliente,
    },
    mounted() {
        this.obtenerClientes();
    },
    computed: {
        filtrarClientes() {
            // Basado en https://stackoverflow.com/questions/62711939/filtering-table-on-the-frontend-vue-js
            let clientesFiltrados = this.info;
            if (this.search.length != 0) {
                clientesFiltrados = clientesFiltrados.filter((cliente) => {
                    return cliente.nombre.search(this.search) != -1;
                });
            }

            return clientesFiltrados;
        },
        rows() {
            return this.info.length
        }
    },
    methods: {
        showModal() {
            this.isModalVisible = true;
        },
        closeModal() {
            this.isModalVisible = false;
        },
        obtenerClientes() {
            axios
                .get("http://localhost:10040/clientes")
                .then((response) => (this.info = response.data.data))
                .catch((error) => console.log(error.response.data));
        },
        editarCliente: function (ID) {
            this.seleccion      = ID.toString();
            this.isModalVisible = true;
        },
        eliminarCliente: function (ID) {
            var r = confirm("¿Estás seguro de eliminar este registro?");
            if (r == true) {
                axios
                    .delete("http://localhost:10040/clientes/" + ID)
                    .then(() => {
                        // Borrar el cliente localmente para refrescar
                        // https://stackoverflow.com/questions/53142220/auto-reload-list-items-after-deletion-vue-js
                        const cliente = this.info.data.data.findIndex(
                            (p) => p.id === ID
                        );
                        this.info.data.data.splice(cliente, 1);
                    })
                    .catch((error) => console.log(error.response.data));
            }
        },
    },
};
</script>

<style scoped>
    .header-text {
        margin-top: 15px;
        margin-bottom: 15px;
    }

    #barra-busqueda {
        padding: 4px 12px;
        color: rgba(0, 0, 0, 0.7);
        border: 1px solid rgba(0, 0, 0, 0.12);
        background: white;
        margin-bottom: 15px;
        margin-right: 20px;
        width: 300px;
    }

    #clientes-tabla {
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

    .btnEditar {
        background: lightgreen;
    }

    .btnEliminar {
        background: lightcoral;
    }
</style>