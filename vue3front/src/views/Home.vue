nes (7 sloc)  136 Bytes

<template>
  <div class="wrapper">
    <div class="data">

      <div class="mainuserdate">
        <ul>

        <li>Username: {{ state.userData.Username }}</li>
        <li>Energy: {{ state.userData.Energy }}</li>
      </ul>

      </div>
      <div class="userdate">
        <h1>User Data</h1><br>
        <ul >
          <li>ID: {{ state.userData.ID }}</li>
          <li>Status: {{ state.userData.Status }}</li>
          <li>Active: {{ state.userData.Active }}</li>
          <li>CreatedAt: {{ state.userData.CreatedAt }}</li>
          <li>UpdatedAt: {{ state.userData.UpdatedAt }}</li>
        </ul>
      <br>
        <h1>Telegram User Data</h1><br>
      <ul>
        <li>TgId: {{ state.userData.TgId }}</li>
        <li>TgUserName: {{ state.userData.TgUserName }}</li>
        <li>TgFirstLastName: {{ state.userData.TgFirstName }} {{ state.userData.TgLastName }}</li>
        <li>TgLanguageCode: {{ state.userData.TgLanguageCode }}</li>
      </ul>
      </div>
    </div>


    <div class="logout">
      <button @click="doLogout">Logout</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import API from "@/modules/api"
import AuthService from "@/modules/auth"
import { ref, reactive } from "vue"

const state = reactive({
  userData: {
    Active: false,
    CreatedAt: "",
    Energy: 0,
    HashedPassword: "",
    ID: 0,
    Password: "",
    Status: 0,
    TgChatId: 0,
    TgFirstName: "",
    TgId: 0,
    TgLanguageCode: "",
    TgLastName: "",
    TgUserName: "",
    UpdatedAt: "",
    Username: ""
  }
})
API.get("/user").then((response) => {
  state.userData = response.data.data;
}
)

const doLogout = () =>  AuthService.logout()
</script>

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;


.mainuserdate{
  color: clr.$clr-route-header;
  padding: 20px;
  background-color: clr.$bg-route-header;
}

li{
    padding:6px;
}
 li:before {
    padding-right:10px;
    font-weight: bold;
    color: #d1ceff;
    content: ".";
    transition-duration: 0.5s;
}
 li:hover:before {
    color: #024C6F;
    content: ".";
}    

.userdate{
  padding: 20px;
  border: 1px solid whitesmoke;
}


button {
  cursor: pointer;
  max-width: 70px;
  padding: 3px;
  color: clr.$clr-button;
  transition: 0.5s;
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
}
button:hover {
  box-shadow: 0px 0px 5px rgba(60, 41, 75, 0.5);
}

.logout button {
  padding: 10px;
  margin-top: 10px;
  margin-left: 10px;

  background-color: clr.$bg-button;
  transition: 0.5s;
}


</style>