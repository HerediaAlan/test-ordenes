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
        <b-modal v-model="isModalVisible" title="Añadir cliente" hide-footer>
            <CrearCliente
                v-bind:cliente="seleccion"
                @close="closeModal"
            />
        </b-modal>
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
            striped hover small selectable
            @row-selected="onRowSelected"
        >
            <!--https://bootstrap-vue.org/docs/components/table#row-select-support-->
            <template #cell(acciones)="row">
                <b-button size="sm" @click="row.toggleDetails" class="mr-2">
                    {{ row.detailsShowing ? 'Cerrar' : 'Ver'}} Detalles
                </b-button>
            </template>

            <template #row-details="">
                <b-card>
                    <b-row class="mb-2">
                        <b-button @click="editarCliente" variant="info">Editar</b-button>
                        <b-button @click="eliminarCliente" variant="danger">Eliminar</b-button>
                    </b-row>
                </b-card>
            </template>
        </b-table>
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
            columnas: ["nombre", "primerApellido", "segundoApellido", "domicilio", "ciudad", "entidadFederativa", "telefono", "email", "acciones"],
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
        idSeleccion: function() {
            return this.seleccion.toString()
        },
        rows() {
            return this.info.length
        }
    },
    methods: {
        onRowSelected(items) {
            this.seleccion = items[0].ID
        },
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
        editarCliente: function () {
            this.seleccion = this.seleccion.toString()
            this.isModalVisible = true;
        },
        eliminarCliente: function () {
            this.boxTwo = ''
            this.$bvModal.msgBoxConfirm("¿Deseas eliminar a este cliente?", {
                title: "Confirmar eliminar dato",
                size: "sm",
                buttonSize: "sm",
                okVariant: "danger",
                okTitle: "Si",
                cancelTitle: "No",
                footerClass: "p-2",
                centered: true
            })
            .then(value => {
                if (value){
                    axios
                        .delete("http://localhost:10040/clientes/" + this.idSeleccion)
                        .then(() => {
                            // Borrar el cliente localmente para refrescar
                            // https://stackoverflow.com/questions/53142220/auto-reload-list-items-after-deletion-vue-js
                            const cliente = this.info.findIndex(
                                (p) => p.id === this.idSeleccion
                            );
                            this.info.splice(cliente, 1);
                        })
                        .catch((error) => console.log(error.response.data));
                }
            })
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