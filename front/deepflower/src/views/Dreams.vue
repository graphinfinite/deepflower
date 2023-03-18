
<script lang="ts">
import TableLite from "@/components/TableLite.vue"
</script>

<template>
    <div class="wrapper">

        <div class="maindata">
            <h1>Dreams</h1> 
        </div>
        <div class="tablelite">
            <TableLite />
        </div>

    </div>


</template>

<style scoped lang="css">


.maindata{
  color: azure;
  padding: 20px;
  background-color: #15042B;
  margin-bottom: 30px;
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
