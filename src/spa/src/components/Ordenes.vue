<template>
    <div class="component-div">
        <!--https://stackoverflow.com/questions/61934264/how-to-disable-modal-ok-slot-of-b-modal-in-vue-bootstrap-->
        <b-modal v-model="isModalVisible" title="Añadir orden" hide-footer>
             <CrearOrden
                v-bind:orden="seleccion"
                @close="closeModal"
            />
        </b-modal>
        <h2 class="header-text">Ordenes</h2>
        <div id="ordenes-header">
            <input
                v-model="search"
                id="barra-busqueda"
                placeholder="Ingrese orden a buscar"
            />
            <b-button 
                @click="showModal" 
                variant="primary" 
                size="sm"
            >Añadir orden</b-button>
        </div>
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
            striped
            hover
            small
        >
        </b-table>
    </div>
</template>

<script>
import CrearOrden from "./ordenes/CrearOrden";
const axios = require("axios");

export default {
    name: "Ordenes",
    data() {
        return {
            currentPage: 1,
            perPage: 10,
            info: {},
            columnas: ["FechaCreacion", "ClienteID", "Asunto", "Descripcion"],
            search: "",
            seleccion: "",
            isModalVisible: false,
        };
    },
    components: {
        CrearOrden,
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
        rows() {
            return this.info.length;
        },
    },
    methods: {
        showModal() {
            this.isModalVisible = true;
        },
        closeModal() {
            this.isModalVisible = false;
        },
        obtenerOrdenes() {
            axios
                .get("http://localhost:10040/ordenes")
                .then((response) => (this.info = response.data.data))
                .catch((error) => console.log(error.response.data));
        },
        editarOrden: function (ID) {
            this.seleccion = ID.toString();
            this.isModalVisible = true;
        },
        eliminarOrden: function (ID) {
            var r = confirm("¿Estás seguro de eliminar este registro?");
            if (r == true) {
                axios
                    .delete("http://localhost:10040/ordenes/" + ID)
                    .then(() => {
                        // Borrar la orden localmente para refrescar
                        // https://stackoverflow.com/questions/53142220/auto-reload-list-items-after-deletion-vue-js
                        const orden = this.info.findIndex((p) => p.id === ID);
                        this.info.splice(orden, 1);
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
}

#ordenes-tabla {
    width: 100%;
}
</style>