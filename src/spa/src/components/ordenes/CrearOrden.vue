<template>
    <transition name="modal-fade">
        <div class="modal-backdrop">
            <div class="modal">
                <form 
                    id="form-orden"
                    v-on:submit.prevent="onSubmit" 
                    method="POST">
                    <p>
                        <label for="fechaCreacion">Fecha de Creación</label>
                        <input type="text" v.model="fecha_creacion" name="fecha_creacion" id="fechaCreacion">
                    </p>
                    <p>
                        <label for="primerApellido">Cliente ID</label>
                        <input type="text" v.model="clienteID" name="clienteID" id="clienteID">
                    </p>
                    <p>
                        <label for="asunto">Asunto</label>
                        <input type="text" v.model="asunto" name="asunto" id="asunto">
                    </p>
                    <p>
                        <label for="domicilio">Descripcion</label>
                        <input type="text" v.model="descripcion" name="descripcion" id="descripcion">
                    </p>
                    <div class="form-buttons">
                        <p><input type="submit" value="Guardar" class="btn"></p>
                        <p><input type="button" name="cancel" value="Cancelar" class="btn" v-on:click="close"/></p>
                    </div>
                </form>
            </div>
        </div>
    </transition>
</template>

<script>
const axios = require('axios')

export default {
    name: 'CrearOrden',
    data() {
        return {
            fechaCreacion: "",
            clienteID: "",
            asunto: "",
            descripcion: "",
        }        
    },
    beforeUpdate() {
        this.obtenerDatos()
    },
    methods: {
        onSubmit() {
            // https://stackoverflow.com/questions/49162345/prevent-form-from-submitting-with-vue-js-and-axios
            if (this.clienteID === "") {
                axios({
                    method: 'post',
                    url: 'http://localhost:10040/ordenes',
                    data: new FormData(document.getElementById('form-orden')),
                    config: {
                        headers: {'Content-Type': 'multipart/form-data'}
                }
                })
                .then(this.Close())
                .catch(function(response) { console.log('error', response); });
            } else {
                axios({
                    method: 'put',
                    url: 'http://localhost:10040/ordenes/',
                    data: new FormData(document.getElementById('form-orden')),
                    config: {
                        headers: {'Content-Type': 'multipart/form-data'}
                }
                })
                .then(this.Close())
                .catch(function(response) { console.log('error', response); })
            }
        },
        close() {
            this.$emit('close')
        }
    }
}
</script>

<style scoped>
    .modal-backdrop {
        position: fixed;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background-color: rgba(0, 0, 0, 0.3);
        display: flex;
        justify-content: center;
        align-items: center;
    }

    #form-orden {
        padding: 15px;
        display: grid;
        grid-template-columns: auto auto;
        grid-template-rows: auto auto;
        column-gap: 5px;
        justify-items: center;
        height: auto;
    }

    #form-orden p {
        display: grid;
        justify-self: start;
        margin: 0px 10px 10px 0px;
    }

    .form-buttons {
        margin-top: 15px;
        grid-column: 1 / span 2;
        width: 100%;
    }

    .btn {
        padding: 7px;
        color: white;
        background: blue;
    }

    .modal {
        background: #FFFFFF;
        box-shadow: 2px 2px 20px 1px;
        overflow-x: auto;
        display: flex;
        flex-direction: column;
    }

    .modal-fade-enter,
    .modal-fade-leave-to {
        opacity: 0;
    }

    .modal-fade-enter-active,
    .modal-fade-leave-active {
        transition: opacity .5s ease;
  }
</style>