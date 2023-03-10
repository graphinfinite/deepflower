import { createRouter, createWebHistory,type Router, type RouterHistory } from 'vue-router'
import routes from './routes'

const router: Router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router