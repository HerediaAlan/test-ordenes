<template>
    <div id="app">
        <Menu></Menu>
        <router-view></router-view>
    </div>
</template>

<script>
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import Menu from "./components/Menu";

export default {
    name: "App",
    components: {
        Menu
    }
};
</script>

<style>
* {
    margin: 0;
    padding: 0;
    font-family: Avenir, Helvetica, Arial, sans-serif;
}

body,
html {
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
