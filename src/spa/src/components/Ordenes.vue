<template>
    <div class="component-div">
        <h2 class="header-text">Ordenes</h2>
        <div id="ordenes-header">
            <input
                v-model="search"
                id="barra-busqueda"
                placeholder="Ingrese orden a buscar"
            />
            <b-button @click="crearOrden" variant="primary" size="sm"
                >Añadir orden</b-button
            >
        </div>
        <!--https://stackoverflow.com/questions/61934264/how-to-disable-modal-ok-slot-of-b-modal-in-vue-bootstrap-->
        <b-modal v-model="isModalVisible" title="Añadir orden" hide-footer>
            <CrearOrden v-bind:orden="seleccion" @close="closeModal" />
        </b-modal>
        <b-modal v-model="isComentarioModalVisible" title="Añadir comentario" hide-footer>
            <AgregarComentario v-bind:ordenID="seleccion" @close="closeComentarioModal" />
        </b-modal>
        <b-pagination
            v-model="currentPage"
            :total-rows="rows"
            :per-page="perPage"
            aria-controls="ordenes-tabla"
        ></b-pagination>
        <b-table
            id="ordenes-tabla"
            :items="filtrarOrdenes"
            :per-page="perPage"
            :current-page="currentPage"
            :fields="columnas"
            @row-clicked="onRowClicked"
            striped hover small selectable>

            <template #cell(show_details)="">
            </template>

            <template #row-details="row">
                <b-card class="mb-2">
                    <b-table-lite 
                        :fields="comentarioColumnas"
                        :items="row.item.OrdenComentarios"></b-table-lite>
                    <b-row class="mb-2">
                        <b-button @click="agregarComentario(row)" variant="primary">Agregar comentario</b-button>
                        <b-button @click="editarOrden(row)" variant="info">Editar</b-button>
                        <b-button @click="eliminarOrden(row)" variant="danger">Eliminar</b-button>
                    </b-row>
                </b-card>
            </template>
        </b-table>
    </div>
</template>

<script>
import CrearOrden from "./ordenes/CrearOrden";
import AgregarComentario from './ordenes/AgregarComentario';
const axios = require("axios");

export default {
    name: "ordenes",
    data() {
        return {
            currentPage: 1,
            perPage: 10,
            info: {},
            columnas: ["FechaCreacion", "ClienteID", "Asunto", "Descripcion"],
            comentarioColumnas: ["FechaCreacion", "DescripcionSeguimiento"],
            search: "",
            seleccion: "",
            isModalVisible: false,
            isComentarioModalVisible: false
        };
    },
    components: {
        CrearOrden, AgregarComentario
    },
    mounted() {
        this.obtenerOrdenes();
    },
    computed: {
        filtrarOrdenes() {
            // Basado en https://stackoverflow.com/questions/62711939/filtering-table-on-the-frontend-vue-js
            let ordenesFiltradas = this.info;
            if (this.search.length != 0) {
                ordenesFiltradas = ordenesFiltradas.filter((orden) => {
                    return orden.FechaCreacion.search(this.search) != -1;
                });
            }

            return ordenesFiltradas;
        },
        idSeleccion: function() {
            return this.seleccion.toString()
        },
        rows() {
            return this.info.length;
        },
    },
    methods: {
        onRowClicked(item) {
            // https://stackoverflow.com/questions/51836186/get-row-element-form-row-clicked-event
            item._showDetails = !item._showDetails
        },
        showModal() {
            this.isModalVisible = true;
        },
        closeModal() {
            this.isModalVisible = false;
        },
        closeComentarioModal() {
            this.isComentarioModalVisible = false;
        },
        crearOrden() {
            this.seleccion = "",
            this.showModal()
        },
        obtenerOrdenes() {
            axios
                .get("http://localhost:10040/ordenes")
                .then((response) => {
                    const data = response.data.data
                    // https://stackoverflow.com/questions/38922998/add-property-to-an-array-of-objects
                    data.forEach(element => {
                        element.FechaCreacion = new Date(element.FechaCreacion)
                        if (element.OrdenComentarios.length > 0) {
                            element.OrdenComentarios.forEach((comentario) => {
                                comentario.FechaCreacion = new Date(comentario.FechaCreacion)
                            })
                        }
                        element._showDetails = false
                    });
                    this.info = data
                })
                .catch((error) => console.log(error.response.data));
        },
        agregarComentario: function (item) {
            this.seleccion = item.item.ID
            this.isComentarioModalVisible = true
        },
        editarOrden: function (item) {
            this.seleccion = item.item.ID
            this.isModalVisible = true
        },
        eliminarOrden: function (item) {
            this.seleccion = item.item.ID
            this.boxTwo = ''
            this.$bvModal.msgBoxConfirm("¿Deseas eliminar esta orden?", {
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
                        .delete("http://localhost:10040/ordenes/" + this.idSeleccion)
                        .then(() => {
                            // Borrar el cliente localmente para refrescar
                            // https://stackoverflow.com/questions/53142220/auto-reload-list-items-after-deletion-vue-js
                            const ordenIndex = this.info.findIndex(
                                (p) => p.id === this.idSeleccion
                            );
                            this.info.splice(ordenIndex, 1);
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
}

#ordenes-tabla {
    width: 100%;
}
</style>