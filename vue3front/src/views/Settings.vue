<template>
    <div class="wrapper">
        <div class="maindata">
        <h1>Settings</h1> 
        </div>
    </div>


</template>

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;

.maindata{
  color: clr.$clr-route-header;
  padding: 20px;
  background-color: clr.$bg-route-header;
  font-size: clr.$route-header-fontsize;
}

.wrapper{
  background-color: white;

}
</style>


<script setup lang="ts">


</script>