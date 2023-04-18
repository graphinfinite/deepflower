nes (7 sloc)  136 Bytes

<template>
  <div class="wrapper">
      <div class="header">
        <div>Home</div>

        
        <div class="logout">
          <div @click="doLogout">out >></div>
        </div>
      </div>
      <div class="energyname">
        <ul>
          <li>ID: {{ state.userData.ID }}</li>
          <li>Name {{ state.userData.Username }}</li>
          <li>Energy {{ state.userData.Energy }}</li>
        </ul>
        <br>


        <h1>Telegram</h1>
        <br>
        <ul>
          <li>Id: {{ state.userData.TgId }}</li>
          <li>Username: {{ state.userData.TgUserName }}</li>
          <li>FLname: {{ state.userData.TgFirstName }} {{ state.userData.TgLastName }}</li>
          <li>Lang: {{ state.userData.TgLanguageCode }}</li>
        </ul>

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
    //HashedPassword: "",
    ID: 0,
    //Password: "",
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


.wrapper {

}



.header{

  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;

  text-shadow: 0px 0px 3px rgb(255, 255, 255);
  color: clr.$clr-route-header;
  padding: 20px;
  //background-color: clr.$bg-route-header;
  font-size: clr.$route-header-fontsize;
  background-image:linear-gradient(90deg, #000000cb,#24084dd7), url('@/assets/purple-moss-leaves-plant-girly-background-header.jpg'); //
}


.header .logout {
  align-self:flex-end;


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

.energyname {
  padding: 20px;
  height: 100%;
}
.logout {
  cursor: pointer;

}



</style>