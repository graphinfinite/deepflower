<script setup lang="ts">

import { RouterView } from 'vue-router'
import SidebarMenu from '@/components/SidebarMenu.vue';

import Login from '@/components/Login.vue'

</script>

<template>
  <div id="root">
    <div v-if="notAuth"><Login/></div>
    <SidebarMenu class="sidebar" />
    <RouterView class="router-view" />
  </div>
</template>


<script lang="ts">
export default {
  data() {
    return {
      notAuth: false
    }

}
}
</script>

<style scoped lang="scss">
#root {
   display: grid;
   grid-template-areas: 'sidebar router-view';
   grid-template-columns: auto 1fr;
    font-family: sans-serif;
}
 .sidebar {
   grid-area: 'sidebar';
 }
 .router-view {
   grid-area: 'router-view';
   padding: 1rem;
 }
</style>
