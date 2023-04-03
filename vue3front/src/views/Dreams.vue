
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

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;


.maindata{
  color: clr.$clr-route-header;
  padding: 20px;
  background-color:clr.$bg-route-header;
  font-size: clr.$route-header-fontsize;
}


.wrapper{
  background-color: rgb(253, 253, 253);
   

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
