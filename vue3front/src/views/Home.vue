nes (7 sloc)  136 Bytes

<template>
<div class="wrapper">
  <div class="header">
    <div>Home</div>

        
  <div class="logout">
      <div @click="doLogout">ᛪ</div>
    </div>
  </div>



<div class="userinfo">
  <div class="userinfo-main">
    <div id="username"><span>{{ state.userData.Username }}</span></div>
    <div id="userenergy">E = <span>{{ state.userData.Energy }}</span></div>
  </div>


<div class="info-menu">
  <div @click="()=>{visibleTgInfo=!visibleTgInfo}">᧠</div>
  <div>᧡</div>
  <div>᧰</div>
  <div>᧽</div>
  <div>᧯</div>
</div>

<div class="info-menu-item">
  <div v-if="visibleTgInfo" class="userinfo-tgsettings">
    <ul>
      <li>Id: {{ state.userData.TgId }}</li>
      <li>Username: {{ state.userData.TgUserName }}</li>
      <li>FLname: {{ state.userData.TgFirstName }} {{ state.userData.TgLastName }}</li>
      <li>Lang: {{ state.userData.TgLanguageCode }}</li>
      <li>ChatId: {{ state.userData.TgChatId }}</li>
    </ul>
  </div>
</div>


</div>


<div class="process-panel">
  

  <h1>Processes</h1>

  <div class="ert">No active processes</div>
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
    ID: 0,
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


const visibleTgInfo = ref(false)


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
  color: clr.$clr-route-header;
  padding: 20px;
  font-size: clr.$route-header-fontsize;
  background-image:linear-gradient(90deg, #000000cb,#24084dd7), url('@/assets/purple-moss-leaves-plant-girly-background-header.jpg'); //
}


.header .logout {
  align-self:flex-end;
  cursor: pointer;
  font-size: 15px;


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

.userinfo {
  padding: 20px;
  height: 100%;

  border-bottom: 7px solid #0B0410;
  
}

.userinfo-main {
  margin-bottom: 20px;
}
.userinfo-tgsettings {


}

.userinfo-tgsettings h1 {

  margin-left: 5px;
  margin-bottom: 10px;

  
}

.userinfo #username,#userenergy {
  font-size: 20px;
  padding: 10px;
}

.userinfo #username span{
  color: #043b02;
}

.userinfo #userenergy span{
  color: #3b035c;
}



.info-menu {


  display: flex;
  flex-direction: row;
}

.info-menu div{
  font-size: 13px;
  padding:5px;
  border: 1px solid whitesmoke;
  margin: 3px;

  cursor: pointer;
  color: clr.$clr-button;
  transition: 0.5s;
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
}


.info-menu div:hover {
  box-shadow: 0px 0px 5px rgba(60, 41, 75, 0.5);
}

.info-menu-item{

  color: grey;
  padding-top: 10px;
}


.process-panel{
  display: flex;
  flex-direction: column;
  padding: 20px;
  border-bottom: 7px solid #0B0410;
}

.process-panel div{

  margin-top: 10px;
}

.process-panel h1{
  color: #3b035c;
  font-size: 20px;
}



</style>