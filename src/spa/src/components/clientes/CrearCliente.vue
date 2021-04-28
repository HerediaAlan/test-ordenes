<template>
    <transition name="modal-fade">
        <div class="modal-backdrop">
            <div class="modal">
                <form 
                    id="form-cliente"
                    v-on:submit.prevent="onSubmit" 
                    method="POST">
                    <p>
                        <label for="nombre">Nombre</label>
                        <input type="text" v.model="nombre" name="nombre" id="nombre">
                    </p>
                    <p>
                        <label for="primerApellido">Primer Apellido</label>
                        <input type="text" v.model="primer_apellido" name="primer_apellido" id="primerApellido">
                    </p>
                    <p>
                        <label for="segundoApellido">Segundo Apellido</label>
                        <input type="text" v.model="segundo_apellido" name="segundo_apellido" id="segundoApellido">
                    </p>
                    <p>
                        <label for="domicilio">Domicilio</label>
                        <input type="text" v.model="domicilio" name="domicilio" id="domicilio">
                    </p>
                    <p>
                        <label for="ciudad">Ciudad</label>
                        <input type="text" v.model="ciudad" name="ciudad" id="ciudad">
                    </p>
                    <p>
                        <label for="entidadFederativa">Entidad Federativa</label>
                        <input type="text" v.model="entidad_federativa" name="entidad_federativa" id="entidadFederativa">
                    </p>
                    <p>
                        <label for="telefono">Telefono</label>
                        <input type="text" v.model="telefono" name="telefono" id="telefono">
                    </p>
                    <p>
                        <label for="email">Email</label>
                        <input type="text" v.model="email" name="email" id="email">
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
    name: 'CrearCliente',
    props: {
        cliente: {
            type: String,
            default: "",
        },
    },
    data() {
        return {
            nombre: "",
            primerApellido: "",
            segundoApellido: "",
            domicilio: "",
            ciudad: "",
            entidadFederativa: "",
            telefono: "",
            email: ""
        }        
    },
    watch: {
        cliente: function() {
            this.obtenerDatos()
        }
    },
    computed: {
        idCliente: function() {
            return this.cliente.toString()
        }
    },
    beforeUpdate() {
        this.obtenerDatos()
    },
    methods: {
        obtenerDatos() {
            console.log(this.idCliente)
            axios
                .get('http://localhost:10040/clientes/' + this.idCliente)
                .then(response => {
                    const respuesta = response.data.data[0]
                    document.getElementById("nombre").value = respuesta.nombre
                    this.nombre = respuesta.nombre
                    document.getElementById("primerApellido").value = respuesta.primerApellido
                    this.primerApellido = respuesta.primerApellido
                    document.getElementById("segundoApellido").value = respuesta.segundoApellido
                    this.segundoApellido = respuesta.segundoApellido
                    document.getElementById("domicilio").value = respuesta.domicilio
                    this.domicilio = respuesta.domicilio
                    document.getElementById("ciudad").value = respuesta.ciudad
                    this.ciudad = respuesta.ciudad
                    document.getElementById("entidadFederativa").value = respuesta.entidadFederativa
                    this.entidadFederativa = respuesta.entidadFederativa
                    document.getElementById("telefono").value = respuesta.telefono
                    this.telefono = respuesta.telefono
                    document.getElementById("email").value = respuesta.email
                    this.email = respuesta.email
                    })
        },
        onSubmit() {
            // https://stackoverflow.com/questions/49162345/prevent-form-from-submitting-with-vue-js-and-axios
            if (this.cliente === "") {
                axios({
                    method: 'post',
                    url: 'http://localhost:10040/clientes',
                    data: new FormData(document.getElementById('form-cliente')),
                    config: {
                        headers: {'Content-Type': 'multipart/form-data'}
                }
                })
                .then(this.Close())
                .catch(function(response) { console.log('error', response); });
            } else {
                axios({
                    method: 'put',
                    url: 'http://localhost:10040/clientes/' + this.idCliente,
                    data: new FormData(document.getElementById('form-cliente')),
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

    #form-cliente {
        padding: 15px;
        display: grid;
        grid-template-columns: auto auto;
        grid-template-rows: auto auto;
        column-gap: 5px;
        justify-items: center;
        height: auto;
    }

    #form-cliente p {
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