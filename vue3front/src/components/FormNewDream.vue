<template>
      <div id="dreaminput">

        <h1>Create new dream!</h1>
        <form @submit.prevent="doSend">
          <label for="dreamname">Dream name</label>
          <input type="text" id="dreamname" v-model="dreamname" placeholder="..." autocomplete="off">
          <label for="dreaminfo">Dream info</label>&nbsp;
          <textarea id="dreaminfo" v-model="dreaminfo" placeholder="..."></textarea>
          <label for="location">Location</label>&nbsp;
          <input  id="location" v-model="location" placeholder="...">
          <button type="submit">->...</button> 
          <div class="form-group">
          <div v-if="messageErr" class="alert alert-danger" role="alert">
            {{ messageErr }}
          </div>
        </div>
        </form>
      </div>

</template>

<script lang="ts" setup>
  import { computed, ref, reactive } from "vue";

  import API from "@/modules/api"
  const messageErr = ref("")
  const dreamname = ref("");
  const dreaminfo = ref("");
  const location = ref("");


  const newdream = reactive({
    Name: dreamname,
    Info: dreaminfo,
    Location:location
  })

var url = "/dreams"
const doSend = () => API.post(url, JSON.stringify(newdream))

</script>


<style scoped>
#dreaminput {
    font-family: Verdana, sans-serif;
    padding: 20px;
    border: 1px solid #ebe8f0;
}

#dreaminput h1 {
    margin-bottom: 30px;
    margin-top: 30px;
    color:blueviolet;
}

#dreaminput form {
    display: flex;
    width: 100%;
    flex-direction: column;
}

#dreaminput form label {
    margin-top: 20px;

}

#dreaminput form input, textarea {
background-color: #f8f7fa;
width: 100%;
padding: 20px 20px;
border: 1px solid #ded4e9;
border-radius: 4px;
}


#dreaminput form button {
  color: azure;
  background-color: #99caf8;
  cursor: pointer;
  border: none;
  padding: 10px;
  transition: background-color 1.5s ease-in-out;

  margin-top: 30px;
  width: 10%;
}
#dreaminput form button:hover {
  background-color: #ffffff;
}

</style>