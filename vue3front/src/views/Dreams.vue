
<script lang="ts">
import TableLiteDream from "@/components/TableLiteDream.vue"
</script>

<template scoped>
    <div class="wrapper">
        <div class="maindata">
          <h1>Dreams</h1> 
        </div>
        <div class="tablelite">
          <TableLiteDream /> 
        </div>
    </div>
</template>

<style scoped lang="css">


.maindata{
  color: whitesmoke;
  padding: 20px;
  background-color: #172025;
}



.wrapper{
  padding: 20px;
  margin-top: 30px;
}
</style>


<script setup lang="ts">

// import { reactive } from "vue"
// const state = reactive({
//   DreamForCreate: {
//     Name :  "",
//     Info  : "",
//     Location: ""
//   }
// })

</script>
