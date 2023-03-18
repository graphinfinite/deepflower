import type { RouteRecordRaw } from "vue-router"
import Home from "@/views/Home.vue"
import About from "@/views/About.vue"
import Dreams from "@/views/Dreams.vue"
import Tasks from "@/views/Tasks.vue"
import Settings from "@/views/Settings.vue"



const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: Home,
  },
  {
    path: '/about',
    component: About,
  },
  {
    path: '/dreams',
    component: Dreams,
  },
  {
    path: '/tasks',
    component: Tasks,
  },
  {
    path: '/settings',
    component: Settings,
  },
]

export default routes