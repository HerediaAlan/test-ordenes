<template>
    <div>
        <h2>Clientes</h2>
        <table>
            <thead>
                <tr>
                    <th>Nombre</th>
                    <th>Primer Apellido</th>
                    <th>Segundo Apellido</th>
                    <th>Domicilio</th>
                    <th>Ciudad</th>
                    <th>Entidad Federativa</th>
                    <th>Telefono</th>
                    <th>Email</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in info.data.data" :key="item.ID">
                    <td>{{ item.nombre }}</td>
                    <td>{{ item.primerApellido }}</td>
                    <td>{{ item.segundoApellido }}</td>
                    <td>{{ item.domicilio }}</td>
                    <td>{{ item.ciudad }}</td>
                    <td>{{ item.entidadFederativa }}</td>
                    <td>{{ item.telefono }}</td>
                    <td>{{ item.email }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
    // https://router.vuejs.org/guide/advanced/data-fetching.html#fetching-after-navigation
    // https://github.com/axios/axios
    // https://vuejs.org/v2/cookbook/using-axios-to-consume-apis.html
    import { mdiAccount } from '@mdi/js';
    const axios = require('axios')

    export default {
        name: 'Clientes',
        title: 'Clientes',
        icon: mdiAccount,
        data() {
            return {
                info: {}
            }
        },
        mounted() {
            axios
                .get('http://localhost:10040/clientes')
                .then(response => {
                    this.info = response
                })
        }
    }
</script>

<style scoped>
    body {
        background: #212121;
    }
</style>