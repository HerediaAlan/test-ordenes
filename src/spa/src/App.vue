<template>
  <v-app>
    <div id="app">
      <div id="navBar">
        <h2>API Ordenes</h2>
        <router-link :to="{ name: 'Clientes' }" tag="a">Clientes</router-link>
        <router-link :to="{ name: 'Ordenes' }" tag="a">Ordenes</router-link>
      </div>
      <v-content>
        <router-view/>
      </v-content>
    </div>
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
    margin: 0;
    padding: 0;
    font-family: Avenir, Helvetica, Arial, sans-serif;
  }

  body, html {
    height: 100%;
  }

  #app {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    margin: 0;

    display: grid;
    height: 100%;
    grid-template-columns: 140px auto;
  }

  #navBar {
    display: grid;
    grid-template-rows: 40px 40px;
    align-items: start;
    padding-top: 20px;
    margin-right: 15px;
    background: darkblue;
  }

  #navBar h2 {
    font-size: 16px;
    font-weight: bold;
    text-align: center;
    color: white;
  }

  #navBar a {
    padding: 10px 10px 10px 10px;
    align-self: start;
    text-decoration: none;
    text-align: center;
    color: white;
    padding-left: 10px;
  }

  #navBar a:hover {
    background: #000021;
  }

  #navBar a.router-link-active {
    color: lightblue;
  }

  v-toolbar-title {
    margin-bottom: 20px;
  }
</style>
