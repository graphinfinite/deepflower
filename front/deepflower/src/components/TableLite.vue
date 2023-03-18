<script>
import TableLite from "vue3-table-lite";

import { reactive } from "vue"

import API from "@/modules/api"

const table = reactive({
  isLoading: false,
  columns: [
    {
      label: "ID",
      field: "id",
      width: "3%",
      sortable: true,
      isKey: true,
    },
    {
      label: "Name",
      field: "name",
      width: "10%",
      sortable: true,
    },
    {
      label: "Email",
      field: "email",
      width: "15%",
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


const doSearch = (offset, limit, order, sort) => {
  table.isLoading = true;
  //let url = "/dreams"
//   API.get(url)
//   .then((response) => {
    // refresh table rows
    //table.rows = response.data.data;
    // table.totalRecordCount = response.count;
    // table.sortable.order = order;
    // table.sortable.sort = sort;
 // });
};



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

