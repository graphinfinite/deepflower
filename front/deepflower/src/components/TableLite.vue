<script setup>
import TableLite from "vue3-table-lite";

import { reactive, ref, toRaw } from "vue"
import API from "@/modules/api"
  // Init Your table settings
const table = reactive({
isLoading: false,
columns: [
    {
    label: "ID",
    field: "ID",
    width: "1%",
    sortable: true,
    isKey: true,
    },
    {
    label: "Name",
    field: "Name",
    width: "20%",
    sortable: true,
    },
    {
    label: "Information",
    field: "Info",
    width: "40%",
    sortable: true,
    },
    {
    label: "Pub",
    field: "Publised",
    width: "3%",
    sortable: true,
    },
    {
    label: "PubDate",
    field: "PublishAt",
    width: "3%",
    sortable: true,
    },
    {
    label: "E",
    field: "Energy",
    width: "5%",
    sortable: true,
    },
    {
    label: "Loc",
    field: "Location",
    width: "1%",
    sortable: true,
    },
    {
    label: "S",
    field: "Status",
    width: "1%",
    sortable: true,
    },
    {
          label: "",
          field: "quick",
          width: "10%",
          display: function (row) {
            return (
              '<button type="button" data-id="' +
              row.id +
              '" class="is-rows-el quick-btn">Look</button>'
            );
          },

    }
],
rows: [],
totalRecordCount: 0,
sortable: {
    order: "id",
    sort: "asc",
},
});

 // offset, limit, order, sort
const doSearch = () => {
table.isLoading = true;
  
// Start use axios to get data from Server
let url = '/dreams';
API.get(url)
.then((response) => {
// refresh table rows
table.rows = response.data.data;
//table.totalRecordCount = response.count;
//table.sortable.order = order;
//table.sortable.sort = sort;
});
        // End use axios to get data from Server
};
  
/**
 * Table search finished event
 */
const tableLoadingFinish = (elements) => {
table.isLoading = false;
};
// Get data first
doSearch();

const rowDream = reactive({
CountG: 0,
CreatedAt: "Error",
Creater: 0,
Energy: 0,
ID: 0,
Info: "Error",
Location: "Error",
Name: "Error",
Publised: false,
PublishAt: "Error",
Status: "Error",
})

const rowClicked = (row) => {
  console.log("Row clicked!", toRaw(row));
  Object.assign(rowDream,toRaw(row) );
};

// return {
// table,
// doSearch,
// tableLoadingFinish,
// };



</script>

<template>

<div class="dreamroot">
<table-lite
:has-checkbox="true"
:is-loading="table.isLoading"
:columns="table.columns"
:rows="table.rows"
:total="table.totalRecordCount"
:sortable="table.sortable"
:messages="table.messages"
@do-search="doSearch"
@is-finished="table.isLoading = false"
@row-clicked="rowClicked"
/>

<div id="dreamrow">
  <div class="row-name">  Name: {{ rowDream.Name }}</div>
  <div class="row-published">  Published: {{ rowDream.Publised }}</div>
  <div class="row-location">Location: {{ rowDream.Location }}</div>
  <div class="row-creater">Creater: {{ rowDream.Creater }}</div>
  <div class="row-energy">Energy: {{ rowDream.Energy }}</div>

  <div class="row-other">
    ID: {{ rowDream.ID }}
    PublishAt: {{ rowDream.PublishAt }}
    CreatedAt: {{ rowDream.CreatedAt }}
    Status: {{ rowDream.Status }}
    G: {{ rowDream.CountG }}
  </div>

  <div class="row-info">
    <div class="i-label">
      Information
    </div>
    <div class="i-data">
      {{ rowDream.Info }}
    </div>
  </div>


</div>

<div class="control-dream-panel">
      <h1>Панель взаимодействия c мечтой</h1>  

      <div>
        <p>После публикации мечту нельзя будет изменить!</p> 
        <p>На публикацию расходуется 1ед энергии.</p>
        <button @click="deleteDream">Delete Dream</button>
        <button @click="publishDream">Publish Dream</button>
        
      </div>



      <div>
        <p>Вы тратите свою личную энергию на мечту!</p>
        <input type="number" id="energe-input" v-model="EnergyToDream"> {{ EnergyToDream }}
        <button @click="addEnergyToDream">+Energy</button>
      </div>

</div>

</div>


</template>


<style scoped>

.dreamroot {
  font-family: Verdana, sans-serif;


}



::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  color: #fff;
  background-color: #2C4928;
  border-color: #2C4928;
}
::v-deep(.vtl-table td),
::v-deep(.vtl-table tr) {
  border: none;

}
::v-deep(.vtl-paging-info) {
  color: #2C4928;
}
::v-deep(.vtl-paging-count-label),
::v-deep(.vtl-paging-page-label) {
  color: #2C4928;
}
::v-deep(.vtl-paging-pagination-page-link) {
  border: none;
}



#dreamrow {
  display: flex;
  width: 100%;
  flex-direction: column;
  border: 1px solid #ebe8f0;
  padding: 20px;

}

#dreamrow div {

  padding-top: 10px;
}

#dreamrow .row-info .i-data{
background-color: #F8F7FA;
}


.control-dream-panel {
  display: flex;
  width: 100%;
  flex-direction: column;
  border: 1px solid #ebe8f0;
  padding: 20px;
}
.control-dream-panel div {


  padding-top: 15px;
  padding-bottom: 15px;
  border-top: 2px solid #192819;
}



.control-dream-panel button {
  color: azure;
  background-color: #99caf8;
  cursor: pointer;
  border: 1px solid #6294c2;
  padding: 10px;
  transition: background-color 1.5s ease-in-out;

  margin-top: 30px;
  width: 20%;
}
.control-dream-panel button:hover {
  background-color: #ffcaa4;
}

.control-dream-panel #energe-input{
  padding: 10px;

  border: 1px solid #FD781A;
}

</style>

