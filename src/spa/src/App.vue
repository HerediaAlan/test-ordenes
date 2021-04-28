<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title><h2>Menu</h2></v-toolbar-title>
      <router-link :to="{ name: 'Clientes' }">Clientes</router-link>
      <router-link :to="{ name: 'Ordenes' }">Ordenes</router-link>
    </v-toolbar>
    <v-content>
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>
  import Clientes from './components/Clientes'
  import Ordenes from './components/Ordenes'

  export default {
    name: 'App',
    components: Clientes, Ordenes
  }
</script>

<style>
  * {
    font-family: Avenir, Helvetica, Arial, sans-serif;
  }
  
  #app {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }

  v-toolbar-title {
    margin-bottom: 20px;
  }
</style>
