import type { RouteRecordRaw } from "vue-router"
import Home from "@/views/Home.vue"
import About from "@/views/About.vue"



const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/about',
    component: About,
  },
]

export default routes