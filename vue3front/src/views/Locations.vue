
<script lang="ts">
import TableLiteLocations from "@/components/TableLiteLocation.vue"
</script>

<template scoped>
    <div class="wrapper">
        <div class="maindata">
          <h1>Locations</h1> 
        </div>
        <div class="tablelite">
          <TableLiteLocations />
        </div>
    </div>
</template>

<style scoped lang="css">


.maindata{
  color: azure;
  padding: 20px;
  background-color: #172025;
}



.wrapper{
  padding: 20px;
  margin-top: 30px;
}
</style>


<script setup lang="ts">
//

</script>