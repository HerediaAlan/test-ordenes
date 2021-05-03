<template>
    <b-form @submit="onSubmit">
        <b-form-group
            id="descripcionGroup"
            label="Comentario de Seguimiento"
            label-for="descripcionInput"
        >
            <b-form-input
                id="descripcionInput"
                v-model.trim="form.descripcion_seguimiento"
                placeholder="Ej: La orden cambió de prioridad"
                :state="esDescripcionValida"
                aria-describedby="input-live-feedback"
                trim
            ></b-form-input>
            <b-form-invalid-feedback id="input-live-feedback">
                Es necesario registrar un comentario
            </b-form-invalid-feedback>
        </b-form-group>

        <div class="buttons">
            <b-button type="submit" variant="primary">Guardar</b-button>
            <b-button @click="close" variant="danger">Cancelar</b-button>
        </div>
    </b-form>
</template>

<script>
const axios = require("axios");

export default {
    name: "AgregarComentario",
    props: {
        ordenID: {
            type: String,
            default: ""
        },
    },
    data() {
        return {
            form: {
                fecha_creacion: "",
                ordenID: "",
                descripcion_seguimiento: ""
            }
        }
    },
    computed: {
        idOrden: function() {
            return this.ordenID.toString()
        },
        esDescripcionValida() {
            if (this.form.descripcion_seguimiento) {
                return this.esValido(this.form.descripcion_seguimiento)
            } else {
                return false
            }
        }
    },
    methods:{
        esValido(dato) {
            return dato.length > 0 ? true : false
        },
        onSubmit() {
            if (!this.esDescripcionValida) {
                return
            } else {
                var formData = new FormData()
                formData.append("fecha_creacion", new Date().toUTCString())
                formData.append("ordenID", this.idOrden)
                formData.append("descripcion_seguimiento", this.form.descripcion_seguimiento)
                axios({
                    method: "post",
                    url: "http://localhost:10040/ordenes/" + this.idOrden + "/comentarios",
                    data: formData,
                    config: {
                        headers: { "Content-Type": "multipart/form-data"},
                    },
                })
                .then(this.close())
                .catch(function (response) {
                    console.log("error", response)
                })
            }
        },
        close() {
            this.$emit("close");
        },
    }
}
</script>

<style>
</style>