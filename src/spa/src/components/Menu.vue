<template>
    <div id="navBar">
        <h2>API Ordenes</h2>
        <b-navbar-nav>
            <b-button-group vertical>
                <b-button to="/" variant="primary-outline" class="d-flex align-items-center">
                    <b-icon icon="house-door"></b-icon>
                    <router-link :to="{ name:'Home' }" exact>Inicio</router-link>
                </b-button>
                <b-button to="/clientes" variant="primary-outline" class="d-flex align-items-center">
                    <b-icon icon="person-fill"></b-icon>
                    <router-link :to="{ name:'Clientes' }">Clientes</router-link>
                </b-button>
                <b-button to="/ordenes" variant="primary-outline" class="d-flex align-items-center">
                    <b-icon icon="inbox"></b-icon>
                    <router-link :to="{ name:'Ordenes' }">Ordenes</router-link>
                </b-button>
            </b-button-group>
        </b-navbar-nav>
    </div>
</template>

<script>
// Basado fuertemente en https://github.com/ratracegrad/meal-prep/blob/master/src/components/AppNavigation.vue
export default {
    name: "Menu",
    data() {
        return {
            appTitle: "API Ordenes",
            myToggle: false,
            items: [
                { title: "Menu", url: "/menu" },
                { title: "Clientes", url: "/clientes" },
                { title: "Ordenes", url: "/ordenes" },
            ],
        };
    },
};
</script>

<style scoped>
#menuBar {
    width: 120px;
    background: #212121;
}
</style>