import Vue from 'vue'
import Router from 'vue-router'
import Clientes from '../components/Clientes'
import Ordenes from '../components/Ordenes'
import Home from '../components/Home'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home
        },
        {
            path: '/clientes',
            name: 'Clientes',
            component: Clientes
        },
        {
            path: '/ordenes',
            name: 'Ordenes',
            component: Ordenes
        }
    ]
})

