
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

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;


.maindata{
  color: clr.$clr-route-header;
  padding: 20px;
  background-color: clr.$bg-route-header;
}

.wrapper{
  background-color: white;

}
</style>


<script setup lang="ts">
//

</script>