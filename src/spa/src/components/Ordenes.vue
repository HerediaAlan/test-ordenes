<template>
  <div class="component-div">
    <h2 class="title">Ordenes</h2>
    <div id="ordenes-header">
      <input
        v-model="search"
        id="barra-busqueda"
        placeholder="Ingrese orden a buscar"
      />
      <button @click="showModal" class="btnAcciones btnCrear">
        Crear orden
      </button>
    </div>
    <CrearOrden
      v-show="isModalVisible"
      @close="closeModal"
    />
    <table id="ordenes-tabla">
      <thead>
        <tr>
          <th>Fecha de Creación</th>
          <th>Cliente</th>
          <th>Asunto</th>
          <th>Descripción</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="orden in info.data.data" :key="orden.ID">
          <td>{{ orden.FechaCreacion }}</td>
          <td>{{ orden.ClienteID }}</td>
          <td>{{ orden.Asunto }}</td>
          <td>{{ orden.Descripcion }}</td>
          <td colspan="2">
            <button
              v-on:click="editarOrden(orden.ID)"
              class="btnAcciones btnEditar"
            >
              Editar
            </button>
            <button
              v-on:click="eliminarOrden(orden.ID)"
              class="btnAcciones btnEliminar"
            >
              Eliminar
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import CrearOrden from './ordenes/CrearOrden'
const axios = require("axios");

export default {
  name: "Ordenes",
  data() {
    return {
      info: {},
      isModalVisible: false
    };
  },
  components: {
    CrearOrden
  },
  mounted() {
    this.obtenerOrdenes();
  },
  methods: {
    showModal() {
      this.isModalVisible = true
    },
    closeModal() {
      this.isModalVisible = false
    },
    obtenerOrdenes() {
      axios
        .get("http://localhost:10040/ordenes")
        .then((response) => (this.info = response))
        .catch((error) => console.log(error.response.data));
    },
    editarOrden() {
      alert("Editar")
    },
    eliminarOrden() {
      alert("Eliminar")
    }
  },
};
</script>

<style scoped>
    .componentView {
        display: grid;
        grid-template-rows: 60px 60px auto;
        grid-template-areas:
            "header-text"
            "barra-busqueda"
            "ordenes-tabla";    
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
    }

    #ordenes-tabla {
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