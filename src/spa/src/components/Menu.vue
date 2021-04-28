<template>
  <v-toolbar app id="menuBar">
      <v-toolbar-title>Menu</v-toolbar-title>
      <router-link :to="{ name: 'Clientes' }">Clientes</router-link>
  </v-toolbar>
</template>

<script>
    import Clientes from './Clientes'

    export default {
        name: 'Menu',
        components: Clientes
    }
</script>

<style scoped>
    #menuBar {
        width: 100px;
        background: #212121;
    }
</style>