<template>
  
<div id="app">

  <div v-if="notAuth"><Login/></div>
  <h1 v-else>Main</h1>

</div>
</template>

<script>
import Login from './components/Login.vue'



export default {
  name: 'App',
  components: {
    Login
  },
  data() {
    return {
      notAuth: true
    }
  }
}
</script>

<style>
</style>
