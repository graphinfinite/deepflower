<script setup>
import TableLite from "vue3-table-lite";

import { reactive } from "vue"
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
],
rows: [],
totalRecordCount: 0,
sortable: {
    order: "id",
    sort: "asc",
},
});




// ​
// CountG: 0  +
// ​​​
// CreatedAt: "2023-03-18T05:15:29.344611Z" +
// ​​​
// Creater: 1 +
// ​​​
// Energy: 0 +
// ​​​
// ID: 0 +
// ​​​
// Info: "info" +
// ​​​
// Location: "" +
// ​​​
// Name: "Best Dream" +
// ​​​
// Publised: false +
// ​​​
// PublishAt: "2023-03-18T05:15:29.344611Z" +
// ​​​
// Status: "created" +
  
/**
 * Table search event
 */

 // offset, limit, order, sort
const doSearch = () => {
table.isLoading = true;
  
// Start use axios to get data from Server
let url = '/dreams';
API.get(url)
.then((response) => {
// Point: your response is like it on this example.
//   {
//   rows: [{
//     id: 1,
//     name: 'jack',
//     email: 'example@example.com'
//   },{
//     id: 2,
//     name: 'rose',
//     email: 'example@example.com'
//   }],
//   count: 2,
//   ...something
// }

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

// return {
// table,
// doSearch,
// tableLoadingFinish,
// };
</script>

<template>
<table-lite
:is-loading="table.isLoading"
:columns="table.columns"
:rows="table.rows"
:total="table.totalRecordCount"
:sortable="table.sortable"
:messages="table.messages"
@do-search="doSearch"
@is-finished="table.isLoading = false"
/>
</template>

