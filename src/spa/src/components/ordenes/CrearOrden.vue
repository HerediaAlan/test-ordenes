<template>
    <b-form id="form-orden" @submit="onSubmit">
        <div class="basicInputs">
            <b-form-group id="clienteGroup" label="ID del cliente" label-for="clienteInput">
                <b-form-input
                    id="clienteInput"
                    v-model="form.clienteID"
                    placeholder="Ej: 3"
                    :state="esClienteValido"
                    aria-describedby="input-live-feedback"
                    trim
                ></b-form-input>
                <b-form-invalid-feedback id="input-live-feedback">
                    Es necesario registrar una clave de cliente
                </b-form-invalid-feedback>
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
                    v-model="form.descripcion"
                    placeholder="Ej: Se realizó una orden por medio de FB"
                ></b-form-input>
            </b-form-group>

            <b-form-group
                v-if="this.orden == ''"
                id="descripcionSeguimientoGroup"
                label="Descripcion de Seguimiento"
                label-for="descripcionSeguimientoInput"
            >
                <b-form-input
                    id="descripcionSeguimientoInput"
                    v-model="form.descripcion_seguimiento"
                    placeholder="Ej: Verificar compra"
                ></b-form-input>
            </b-form-group>
        </div>

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
                clienteID: "",
                asunto: "",
                descripcion: "",
                descripcion_seguimiento: ""
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
            return this.orden.toString()
        },
        esClienteValido() {
            if (this.form.clienteID) {
                const re = /^\d+$/
                return re.test(this.form.clienteID)
            } else {
                return false
            }
        }
    },
    beforeMount() {
        this.obtenerDatos()
    },
    methods: {
        obtenerDatos() {
            axios
                .get("http://localhost:10040/ordenes/" + this.idOrden)
                .then((response) => {
                    const respuesta = response.data.data;
                    this.form.fecha_creacion = respuesta.FechaCreacion;
                    this.form.clienteID = respuesta.ClienteID;
                    this.form.asunto = respuesta.Asunto;
                    this.form.descripcion = respuesta.Descripcion;
                })
                .catch(function (response) {
                    console.log("error", response);
                })
        },
        onSubmit() {
            if (!this.esClienteValido) {
                return
            }
            // https://stackoverflow.com/questions/49162345/prevent-form-from-submitting-with-vue-js-and-axios
            var formDataOrden = new FormData();
            if (this.orden === "") {
                formDataOrden.append(
                    "fecha_creacion",
                    new Date().toUTCString()
                );
                formDataOrden.append("clienteID", this.form.clienteID);
                formDataOrden.append("asunto", this.form.asunto);
                formDataOrden.append("descripcion", this.form.descripcion);
                axios({
                    method: "post",
                    url: "http://localhost:10040/ordenes",
                    data: formDataOrden,
                    config: {
                        headers: { "Content-Type": "multipart/form-data" },
                    },
                })
                    .then((response) => {
                        var formDataComentario = new FormData();
                        formDataComentario.append(
                            "fecha_creacion",
                            new Date().toUTCString()
                        );
                        formDataComentario.append(
                            "ordenID",
                            response.data.result.ID
                        );
                        formDataComentario.append(
                            "descripcion_seguimiento",
                            this.form.descripcion_seguimiento
                        );
                        axios({
                            method: "post",
                            url:
                                "http://localhost:10040/ordenes/" +
                                response.data.result.ID +
                                "/comentarios",
                            data: formDataComentario,
                            config: {
                                headers: {
                                    "Content-Type": "multipart/form-data ",
                                },
                            },
                        });
                    })
                    .then(this.close())
                    .catch(function (response) {
                        console.log("error", response);
                    });
            } else {
                formDataOrden.append("clienteID", this.form.clienteID);
                formDataOrden.append("asunto", this.form.asunto);
                formDataOrden.append("descripcion", this.form.descripcion);
                axios({
                    method: "put",
                    url: "http://localhost:10040/ordenes/" + this.idOrden,
                    data: formDataOrden,
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
#form-orden {
    display: grid;
}

.buttons {
    display: grid;
    justify-items: stretch;
    row-gap: 8px;
    width: 100%;
}
</style>