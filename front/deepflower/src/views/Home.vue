nes (7 sloc)  136 Bytes

<script setup lang="ts">


  import { api } from "@/modules/api"

 const userdata = api.get("")
</script>

<template>
  <div class="wrapper">

    {{ userdata }}
    



  </div>
</template>

<style scoped lang="scss">
</style>


<script setup>
import { computed, ref } from "vue";
setup() {

  const userData = ref({});
  async function load() {
        const url = "user";
        let result = await api.get(url);
        user = result.data
    }
    
    return { userData };
}




}</script>