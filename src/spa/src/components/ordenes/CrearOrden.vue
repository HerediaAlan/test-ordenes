<template>
    <b-form id="form-orden" v-on:submit.prevent="onSubmit">
        <b-form-group
            id="clienteGroup"
            label="ID del cliente"
            label-for="clienteInput"
        >
            <b-form-input
                id="clienteInput"
                v-model="form.clienteID"
                placeholder="Ej: 3"
            ></b-form-input>
        </b-form-group>

        <b-form-group id="asuntoGroup" label="Asunto" label-for="asuntoInput">
            <b-form-input
                id="asuntoInput"
                v-model="form.asunto"
                placeholder="Ej: Compra express"
            ></b-form-input>
        </b-form-group>

        <b-form-group
            id="descripcionGroup"
            label="Descripcion"
            label-for="descripcionInput"
        >
            <b-form-input
                id="descripcionInput"
                v-model="form.descripcion_seguimiento"
                placeholder="Ej: Se realizó una orden por medio de FB"
            ></b-form-input>
        </b-form-group>

        <b-form-group
            id="comentarioGroup"
            label="Comentario"
            label-for="comentarioInput"
        >
            <b-form-input
                id="comentarioInput"
                v-model="form.comentario"
                placeholder="Ej: Añadir paquete extra"
            ></b-form-input>
        </b-form-group>
        <div class="buttons">
            <b-button @click="onSubmit" variant="primary">Guardar</b-button>
            <b-button @click="close" variant="danger">Cancelar</b-button>
        </div>
    </b-form>
</template>

<script>
const axios = require("axios");

export default {
    name: "CrearOrden",
    props: {
        orden: {
            type: String,
            default: "",
        },
    },
    data() {
        return {
            form: {
                fecha_creacion: "",
                clienteID: "",
                asunto: "",
                descripcion_seguimiento: "",
                comentario: "",
            },
        };
    },
    watch: {
        orden: function () {
            this.obtenerDatos();
        },
    },
    computed: {
        idOrden: function () {
            return this.orden.toString();
        },
    },
    beforeUpdate() {
        this.obtenerDatos();
    },
    methods: {
        obtenerDatos() {
            axios
                .get("http://localhost:10040/ordenes/" + this.idOrden)
                .then((response) => {
                    const respuesta = response.data.data;
                    document.getElementById("fechaCreacion").value =
                        respuesta.FechaCreacion;
                    this.fechaCreacion = respuesta.FechaCreacion;
                    document.getElementById("clienteID").value =
                        respuesta.ClienteID;
                    this.clienteID = respuesta.ClienteID;
                    document.getElementById("asunto").value = respuesta.Asunto;
                    this.asunto = respuesta.Asunto;
                    document.getElementById("descripcion").value =
                        respuesta.Descripcion;
                    this.descripcion = respuesta.Descripcion;
                })
                .catch(function (response) {
                    console.log("error", response);
                });
        },
        onSubmit() {
            // https://stackoverflow.com/questions/49162345/prevent-form-from-submitting-with-vue-js-and-axios
            var d = new Date()
            if (this.orden === "") {
                var formDataOrden = new FormData();
                formDataOrden.append(
                    "fecha_creacion",
                    `${d.toLocaleString("default", { weekday: "short" })}, ${d.getDate()} ${d.toLocaleString("default", { month: "short" })} ${d.getFullYear()} ${('0' + d.getHours()).slice(-2)}:${d.getMinutes()}:${d.getSeconds()} MST`
                );
                formDataOrden.append(
                    "clienteID",
                    this.form.clienteID
                );
                formDataOrden.append(
                    "asunto",
                    this.form.asunto
                );
                formDataOrden.append(
                    "descripcion",
                    this.form.descripcion
                )
                axios({
                    method: "post",
                    url: "http://localhost:10040/ordenes",
                    data: formDataOrden,
                    config: {
                        headers: { "Content-Type": "multipart/form-data" },
                    },
                })
                    .then((response) => {
                        var formDataComentario = new FormData()
                        var d = new Date();
                        formDataComentario.append(
                            "fecha_creacion",
                            `${d.toLocaleString("default", { weekday: "short" })}, ${d.getDate()} ${d.toLocaleString("default", { month: "short" })} ${d.getFullYear()} ${('0' + d.getHours()).slice(-2)}:${d.getMinutes()}:${d.getSeconds()} MST`
                        )
                        formDataComentario.append(
                            "ordenID",
                            response.data.result.ID
                        )
                        formDataComentario.append(
                            "descripcion_seguimiento",
                            this.form.descripcion_seguimiento
                        )
                        axios({
                            method: "post",
                            url: "http://localhost:10040/ordenes/" + response.data.result.ID + "/comentarios",
                            data: formDataComentario,
                            config: {
                                headers: {
                                    "Content-Type": "multipart/form-data ",
                                },
                            },
                        })
                    })
                    .then(this.close())
                    .catch(function (response) {
                        console.log("error", response);
                    });
            } else {
                axios({
                    method: "put",
                    url: "http://localhost:10040/ordenes/" + this.idOrden,
                    data: new FormData(document.getElementById("form-orden")),
                    config: {
                        headers: { "Content-Type": "multipart/form-data" },
                    },
                })
                    .then(this.close())
                    .catch(function (response) {
                        console.log("error", response);
                    });
            }
        },
        close() {
            this.$emit("close");
        },
    },
};
</script>

<style scoped>
.buttons {
    display: grid;
    justify-items: stretch;
    row-gap: 8px;
    width: 100%;
}
</style>