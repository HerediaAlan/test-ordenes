<template>
    <transition name="modal-fade">
        <div class="modal-backdrop">
            <div class="modal">
                <form action="http://localhost:10040/clientes" method="POST">
                    <p>
                        <label for="nombre">Nombre</label>
                        <input type="text" v.model="nombre" name="nombre" id="nombre">
                    </p>
                    <p>
                        <label for="primerApellido">Primer Apellido</label>
                        <input type="text" v.model="primerApellido" name="primer_apellido" id="primerApellido">
                    </p>
                    <p>
                        <label for="segundoApellido">Segundo Apellido</label>
                        <input type="text" v.model="segundoApellido" name="segundo_apellido" id="segundoApellido">
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
                        <input type="text" v.model="entidadFederativa" name="entidad_federativa" id="entidadFederativa">
                    </p>
                    <p>
                        <label for="telefono">Telefono</label>
                        <input type="text" v.model="telefono" name="telefono" id="telefono">
                    </p>
                    <p>
                        <label for="email">Email</label>
                        <input type="text" v.model="email" name="email" id="email">
                    </p>

                    <input type="submit" value="Crear">
                    <button @click="close">Cancelar</button>
                </form>
            </div>
        </div>
    </transition>
</template>

<script>
export default {
    name: 'CrearCliente',
    methods: {
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