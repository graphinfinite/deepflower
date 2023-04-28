<script lang="ts">
import { RouterView } from 'vue-router'
import SidebarMenu from '@/components/SidebarMenu.vue';
import Login from '@/components/Login.vue'
</script>

<template>
  <div id="root">
    <SidebarMenu class="sidebar"/>

    <div v-if="isAuthorazed">
      <RouterView class="router-view"/>
      
    </div>
    <div v-else><Login/></div>

  </div>
</template>

<script setup lang="ts">
import {ref} from "vue"
const isAuthorazed = ref(Boolean(localStorage.getItem("tokenAccess")))
</script>
<style scoped lang="scss">
#root {
   display: grid;
   grid-template-areas: 'sidebar router-view';
   grid-template-columns: auto 1fr;
}

 .sidebar {
   grid-area: 'sidebar';
 }
 .router-view {
    grid-area: 'router-view';
 }

</style>
