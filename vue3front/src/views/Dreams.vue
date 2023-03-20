
<script lang="ts">
import TableLite from "@/components/TableLite.vue"
import FormNewDream from "@/components/FormNewDream.vue"
</script>

<template>
    <div class="wrapper">
        <div class="maindata">
          <h1>Dreams</h1> 
        </div>
        <div class="tablelite">
          <TableLite />
        </div>

        <div class="dreamform">
          <FormNewDream />
        </div>
    </div>
</template>

<style scoped lang="css">


.maindata{
  color: azure;
  padding: 20px;
  background-color: #192819;
  margin-bottom: 30px;
}

.dreamform {


  padding:50px;
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
