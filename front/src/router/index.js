import Vue from 'vue'
import Router from 'vue-router'
import Recetas from '@/components/Recetas'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Recetas',
      component: Recetas
    }
  ]
})
