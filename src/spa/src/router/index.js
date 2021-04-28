import Vue from 'vue'
import Router from 'vue-router'
import Clientes from '../components/Clientes'
import Ordenes from '../components/Ordenes'

Vue.use(Router)

export default new Router({
    routes: [
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

