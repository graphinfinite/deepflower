
<script setup lang="ts">
import GraphModeler from '@/components/GraphModeler.vue';
</script>

<template>
    <div class="wrapper">
      <div class="maindata">
        <h1>Tasks</h1> 
      </div>

      <!-- WERY IMPORT ANT -->
      <div class="graphmodeler">
        <GraphModeler />
      </div>
    </div>


</template>

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;


.maindata{
  color: clr.$clr-route-header;
  padding: 20px;
  background-color: clr.$bg-route-header;
}

.graphmodeler{

  
}

.wrapper{
  border-left: 2px solid #ffffff;

}
</style>


