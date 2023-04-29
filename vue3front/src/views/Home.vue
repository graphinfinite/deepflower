nes (7 sloc)  136 Bytes

<template>
<div class="wrapper">
  <div class="header">
    <div>Home</div>

        
  <div class="logout">
      <div @click="doLogout">🢖</div>
    </div>
  </div>



<div class="userinfo">
  <div class="userinfo-main">
    <div id="username"><span @click="()=>{visibleTgInfo=!visibleTgInfo}">{{ state.userData.Username }}</span></div>
    <div id="userenergy">E = <span>{{ state.userData.Energy }}</span></div>
  </div>

<div class="info-menu-item">
  <div v-if="visibleTgInfo" class="userinfo-tgsettings">
    Telegram 
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

  <div class="searchBox">

<label for="checkbox1">OnlyActive: {{ onlyActive }}</label>
<input type="checkbox" id="checkbox1" v-model="onlyActive" />

<label for="filterInput">Search by status:</label>
<input id="filterInput" v-model="searchTerm" />
<button @click="doSearch(0, 10, 'id', 'asc')">ᐅ</button>
</div>

  <div class="process-table">
  <table-lite
  :max-width=300
  :is-loading="table.isLoading"
  :columns="table.columns"
  :rows="table.rows"
  :total="table.totalRecordCount"
  :sortable="table.sortable"
  :messages="table.messages"
  @do-search="doSearch"
  @is-finished="tableLoadingFinish"
  @row-clicked="rowClicked"
  /></div>

</div>
</div>




</template>

<script setup>
import API from "@/modules/api"
import AuthService from "@/modules/auth"
import { ref, reactive } from "vue"

import TableLite from "vue3-table-lite";
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





const onlyActive = ref(true)
const searchTerm = ref("")
const table = reactive({
isLoading: false,
columns: [
    {
    label: "ID",
    field: "ID",
    width: "5%",
    sortable: false,
    isKey: true,
    },
    {
    label: "LeadTime",
    field: "LeadTime",
    width: "5%",
    sortable: true,
    },
    {
    label: "EnergyTotal",
    field: "EnergyTotal",
    width: "5%",
    sortable: true,
    },

    {
    label: "InspectorsTotal",
    field: "InspectorsTotal",
    width: "5%",
    sortable: true,
    },
    {
    label: "InspectorsConfirmed",
    field: "InspectorsConfirmed",
    width: "5%",
    sortable: true,
    },
    {
    label: "S",
    field: "Status",
    width: "7%",
    sortable: true,
    },
    {
    label: "UpdatedAt",
    field: "UpdatedAt",
    width: "3%",
    sortable: true,
    },
],
rows: [],
totalRecordCount: 0,

sortable: {
    order: "id",
    sort: "asc",
},
});
 // 
const doSearch = (offset, limit, order, sort) => {
  var searchData = {
    Offset: offset,
    Limit: limit,
    Order: order,
    Sort: sort,
    OnlyActive: onlyActive.value,
    SearchTerm: searchTerm.value
    }
  console.log(JSON.stringify(searchData))
  table.isLoading = true;
  let url = '/processes';
  API.get(url, {params: searchData} ).then((response) => {
      if (response.data.status === "ok") {
        table.isLoading = false;
        // refresh table rows
        table.rows = response.data.data.Processes;
        table.totalRecordCount = response.data.data.TotalRecordCount;
        table.sortable.order = order;
        table.sortable.sort = sort;
        return
      } 
      window.alert(response.data.message);
  }); 
};



doSearch(0, 10, 'id', 'asc')


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
  cursor: pointer;
}

.userinfo #userenergy span{
  color: #3b035c;
}


.info-menu-item{
  color: grey;
  padding-top: 10px;
}


.process-panel{
  display: flex;
  flex-direction: column;
  border-bottom: 7px solid #0B0410;

  padding-bottom: 30px;
}


.process-panel div{

  margin-top: 10px;
}

.process-panel h1{
  color: #3b035c;
  font-size: 20px;
  margin-left: 20px;
  margin-top: 20px;
}


.searchBox {
  border: 1px solid whitesmoke;
  padding: 10px;
  background-color:white ;
}
.searchBox #checkbox1, #checkbox2  {
  cursor:pointer;
  border: 1px solid black;
  padding: 5px;
  background-color: blueviolet;
  margin-left: 3px;
  margin-right: 7px;
}

.searchBox #checkbox1:checked {
  background-color: #365778;
}
.searchBox label {
  cursor:default;
}
.searchBox #filterInput {
  margin-left: 15px;
  padding: 10px;
  border: 1px solid whitesmoke;
}


.searchBox button{
  font-size: 13px;
  padding:5px;
  border: 1px solid whitesmoke;
  margin: 3px;

  cursor: pointer;
  color: clr.$clr-button;
  transition: 0.5s;
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);



}


.searchBox button:hover{
  box-shadow: 0px 0px 5px rgba(60, 41, 75, 0.5);

}




</style>