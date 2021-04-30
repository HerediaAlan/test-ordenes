<template>
    <b-form id="form-cliente" @submit="onSubmit">
        <b-form-group id="nombreGroup" label="Nombre" label-for="nombreInput">
            <b-form-input
                id="nombreInput"
                v-model.trim="form.nombre"
                placeholder="Ej: José"
                :state="esNombreValido"
                aria-describedby="input-live-feedback"
                trim
            ></b-form-input>
            <b-form-invalid-feedback id="input-live-feedback">
                Es necesario registrar un nombre
            </b-form-invalid-feedback>
        </b-form-group>

        <b-form-group id="´primerApellidoGroup" label="Primer Apellido" label-for="primerApellidoInput">
            <b-form-input
                id="primerApellidoInput"
                v-model="form.primer_apellido"
                placeholder="Ej: García"
                :state="esPrimerApellidoValido"
                aria-describedby="input-live-feedback"
            ></b-form-input>
            <b-form-invalid-feedback id="input-live-feedback">
                Es necesario registrar este apellido
            </b-form-invalid-feedback>
        </b-form-group>

        <b-form-group id="segundoApellidoGroup" label="Segundo Apellido" label-for="segundoApellidoInput">
            <b-form-input
                id="segundoApellidoInput"
                v-model="form.segundo_apellido"
                placeholder="Ej: Gonzales"
            ></b-form-input>
        </b-form-group>

        <b-form-group
            id="domicilioGroup"
            label="Domicilio"
            label-for="domicilioInput"
        >
            <b-form-input
                id="domicilioInput"
                v-model="form.domicilio"
                placeholder="Ej: Av. Tlaxcala y 23"
            ></b-form-input>
        </b-form-group>

        <b-form-group id="ciudadGroup" label="Ciudad" label-for="ciudadInput">
            <b-form-input
                id="ciudadInput"
                v-model="form.ciudad"
                placeholder="Ej: San Luís Río Colorado"
            ></b-form-input>
        </b-form-group>

        <b-form-group
            id="entidadFederativaGroup"
            label="Entidad Federativa"
            label-for="entidadFederativaInput"
        >
            <b-form-input
                id="entidadFederativaInput"
                v-model="form.entidad_federativa"
                placeholder="Ej: Sonora"
                :state="esEntidadFederativaValida"
                aria-describedby="input-live-feedback"
            ></b-form-input>
            <b-form-invalid-feedback id="input-live-feedback">
                Es necesario especificar el estado
            </b-form-invalid-feedback>
        </b-form-group>

        <b-form-group
            id="telefonoGroup"
            label="Telefono"
            label-for="telefonoInput"
        >
            <b-form-input
                id="telefonoInput"
                v-model="form.telefono"
                placeholder="Ej: 6531000500"
                :state="esTelefonoValido"
                aria-describedby="input-live-feedback"
            ></b-form-input>
            <b-form-invalid-feedback id="input-live-feedback">
                Es necesario registrar un teléfono
            </b-form-invalid-feedback>
        </b-form-group>

        <b-form-group id="emailGroup" label="Email" label-for="emailInput">
            <b-form-input
                id="emailInput"
                v-model="form.email"
                placeholder="Ej: nombre@gmail.com"
                :state="esEmailValido"
                aria-describedby="input-live-feedback"
            ></b-form-input>
            <b-form-invalid-feedback id="input-live-feedback">
                Es necesario registrar un email
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
    name: "CrearCliente",
    props: {
        cliente: {
            type: String,
            default: "",
        },
    },
    data() {
        return {
            form: {
                nombre: "",
                primer_apellido: "",
                segundo_apellido: "",
                domicilio: "",
                ciudad: "",
                entidad_federativa: "",
                telefono: "",
                email: "",
            },
        };
    },
    watch: {
        cliente: function () {
            this.obtenerDatos();
        },
    },
    computed: {
        idCliente: function () {
            return this.cliente.toString();
        },
        esNombreValido() {
            if (this.form.nombre) {
                return this.esValido(this.form.nombre)
            } else {
                return false
            }
        },
        esPrimerApellidoValido() {
            if (this.form.primer_apellido) {
                return this.esValido(this.form.primer_apellido)
            } else {
                return false
            }
        },
        esEntidadFederativaValida() {
            if (this.form.entidad_federativa) {
                return this.esValido(this.form.entidad_federativa)
            } else {
                return false
            }
        },
        esTelefonoValido() {
            if (this.form.telefono) {
                return this.esValido(this.form.telefono)
            } else {
                return false
            }
        },
        esEmailValido() {
            if (this.form.email) {
                const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
                return re.test(this.form.email)
            } else {
                return false
            }
        }
    },
    beforeMount() {
        this.obtenerDatos();
    },
    methods: {
        esValido(dato) {
            return dato.length > 0 ? true : false
        },
        obtenerDatos() {
            axios
                .get(
                    "http://localhost:10040/clientes/" + this.cliente.toString()
                )
                .then((response) => {
                    const respuesta = response.data.data[0];
                    this.form.nombre = respuesta.nombre;
                    this.form.primer_apellido = respuesta.primerApellido;
                    this.form.segundo_apellido = respuesta.segundoApellido;
                    this.form.domicilio = respuesta.domicilio;
                    this.form.ciudad = respuesta.ciudad;
                    this.form.entidad_federativa = respuesta.entidadFederativa;
                    this.form.telefono = respuesta.telefono;
                    this.form.email = respuesta.email;
                })
                .catch(function (response) {
                    console.log("error", response);
                });
        },
        onSubmit() {
            if (!this.esNombreValido || !this.esPrimerApellidoValido || !this.esEntidadFederativaValida
                || !this.esTelefonoValido || !this.esEmailValido) {
                return
            }
            var formData = new FormData();
            // https://stackoverflow.com/questions/49162345/prevent-form-from-submitting-with-vue-js-and-axios
            if (!this.cliente) {
                alert("POST")
                formData.append("nombre", this.form.nombre);
                formData.append("primer_apellido", this.form.primer_apellido);
                formData.append("segundo_apellido", this.form.segundo_apellido);
                formData.append("domicilio", this.form.domicilio);
                formData.append("ciudad", this.form.ciudad);
                formData.append(
                    "entidad_federativa",
                    this.form.entidad_federativa
                );
                formData.append("telefono", this.form.telefono);
                formData.append("email", this.form.email);
                axios({
                    method: "post",
                    url: "http://localhost:10040/clientes",
                    data: formData,
                    config: {
                        headers: { "Content-Type": "multipart/form-data" },
                    },
                })
                    .then(this.close())
                    .catch(function (response) {
                        console.log("error", response);
                    });
            } else {
                alert("PUT")
                formData.append("nombre", this.form.nombre);
                formData.append("primer_apellido", this.form.primer_apellido);
                formData.append("segundo_apellido", this.form.segundo_apellido);
                formData.append("domicilio", this.form.domicilio);
                formData.append("ciudad", this.form.ciudad);
                formData.append(
                    "entidad_federativa",
                    this.form.entidad_federativa
                );
                formData.append("telefono", this.form.telefono);
                formData.append("email", this.form.email);
                axios({
                    method: "put",
                    url: "http://localhost:10040/clientes/" + this.idCliente,
                    data: formData,
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
#form-cliente {
    display: grid;
    grid-template-columns: auto auto;
    column-gap: 10px;
}

.buttons {
    grid-column: 1 / 3;
    display: grid;
    justify-items: stretch;
    row-gap: 8px;
    width: 100%;
}
</style>