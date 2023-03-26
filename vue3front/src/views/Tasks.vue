
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

<style scoped lang="css">


.maindata{
  color: blueviolet;
  padding: 20px;
  background-color: #15042B;
  margin-bottom: 30px;
}

.graphmodeler{
  
}

.wrapper{
  padding: 20px;
  margin-top: 30px;
}
</style>


