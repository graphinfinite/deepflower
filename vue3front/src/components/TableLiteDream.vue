<script setup>
import TableLite from "vue3-table-lite";
import { reactive, ref, toRaw, computed } from "vue"
import API from "@/modules/api"
  // init table settings
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
    label: "Name",
    field: "Name",
    width: "20%",
    sortable: true,
    display: (row) => {
        if (row.Name) {
          if (row.Name.length>15) {
            return (row.Name.slice(0, 15)+"...")
          };
          return (row.Name);
        } else {
          return ("Empty")
        };
},
    },
    {
    label: "Information",
    field: "Info",
    width: "30%",
    sortable: false,
    display: (row) => {
        if (row.Info) {
          if (row.Info.length>30) {
            return (row.Info.slice(0, 31)+"...")
          };
          return (row.Info);
        } else {
          return ("<b>Empty</b>")
        };
},
    },
    {
    label: "Pub",
    field: "Published",
    width: "3%",
    sortable: true,
    },
    {
    label: "PubDate",
    field: "PublishAt",
    width: "3%",
    sortable: true,
    display: (row) => {
        if (row.PublishAt) {
          return (row.PublishAt.slice(0, 19));
        } else {
          return ("Empty")
        };
},
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
    display: (row) => {
        if (row.Location) {
          if (row.Location.length>30) {
            return (row.Info.slice(0, 31)+"...")
          };
          return (row.Location);
        } else {
          return ("<b>Empty</b>")
        };
},
    },
    {
    label: "S",
    field: "Status",
    width: "1%",
    sortable: false,
    },
],
rows: [],
totalRecordCount: 0,

sortable: {
    order: "id",
    sort: "asc",
},
});

const onlyMyDreams = ref(true)
const searchTerm = ref("")
 // 
const doSearch = (offset, limit, order, sort) => {
  var searchData = {
    Offset: offset,
    Limit: limit,
    Order: order,
    Sort: sort,
    OnlyMyDreams: onlyMyDreams.value,
    SearchTerm: searchTerm.value
    }
  console.log(JSON.stringify(searchData))
  table.isLoading = true;
  let url = '/dreams';
  API.get(url, {params: searchData} ).then((response) => {
      if (response.data.status === "ok") {
        table.isLoading = false;
        // refresh table rows
        table.rows = response.data.data.Dreams;
        table.totalRecordCount = response.data.data.TotalRecordCount;
        table.sortable.order = order;
        table.sortable.sort = sort;
        return
      } 
      window.alert(response.data.message);
  }); 
};
  
/**
 * Table search finished event
 */
const tableLoadingFinish = (elements) => {
table.isLoading = false;
};

doSearch(0, 10, "id", "asc");

const rowDream = reactive({
      CountG: 0,
      CreatedAt: "",
      Creater: 0,
      Energy: 0,
      ID: 0,
      Info: "",
      Location: "",
      Name: "",
      Published: false,
      PublishAt: "",
      Status: "",
})

const rowClicked = (row) => {
  Object.assign(rowDream,toRaw(row) );
};


const deleteDream = () => {
  if (rowDream.Name ==="") {
    window.alert("Error: Row is empty!");
    return
  }
  let url = '/dreams/'+rowDream.ID;
  API.delete(url).then((response) => {
    if (response.data.status == "ok") {
      table.rows = table.rows.filter(function(elem) {
          if (elem.Name == rowDream.Name) {return false;} else {return true;}
      });
      rowDream.Name = "";
    }
    window.alert(response.data.message);

  });
};

const publishDream = () => {
  if (rowDream.Name ==="") {
    window.alert("Error: Row is empty!");
    return
  }
  
  let url = '/dreams/'+rowDream.ID+ '/publish';

  console.log(url)
  API.post(url).then((response) => {
    if (response.data.status === "ok") {
      rowDream.Published = true;
      doSearch(0, 10, "id", "asc") 
    }
    window.alert(response.data.message);
  });
};



const messageErr = ref("")
const dreamname = ref("");
const dreaminfo = ref("");
const location = ref("");
const newdream = reactive({
    Name: dreamname,
    Info: dreaminfo,
    Location:location
})
const doSend = () => API.post("/dreams", JSON.stringify(newdream)).then((response) => {
    if (response.data.status === "ok") {
      doSearch(0, 10, "id", "asc")
      window.alert(response.data.message)
      return
    }
    window.alert(response.data.message)

})


const energyToDream = ref(0)
const addEnergyToDream = () => {
  if (energyToDream.value === 0) {
    window.alert("add zero energy???")
    return
  }
  API.post("/dreams/"+rowDream.ID+"/energy", JSON.stringify({Energy: energyToDream.value})).then((response) => {
    if (response.data.status === "ok") {
      rowDream.Energy += energyToDream.value
      doSearch(0, 10, "id", "asc")
      return
    }
    window.alert(response.data.message)

})
}



</script>

<template>

<div class="searchBox">

  <label for="checkbox1">Only my dreams: {{ onlyMyDreams }}</label>
  <input type="checkbox" id="checkbox1" v-model="onlyMyDreams" />

  <label for="filterInput">SearchBy:</label>
  <input id="filterInput" v-model="searchTerm" />
  <button @click="doSearch(0, 10, 'id', 'asc')">GO</button>
</div>


<div class="dreamroot">
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
  />
  

<div v-if='rowDream.Name !==""'>
  <div id="dreamrow">
    <div class="row-name">  Name: {{ rowDream.Name }}</div>
    <div class="row-published">  Published: {{ rowDream.Published }}</div>
    <div class="row-location">Location: {{ rowDream.Location }}</div>
    <div class="row-creater">Creater: {{ rowDream.Creater }}</div>
    <div class="row-energy">Energy: {{ rowDream.Energy }}</div>

    <div class="row-other">
      ID: {{ rowDream.ID }} PublishAt: {{ rowDream.PublishAt }} CreatedAt: {{ rowDream.CreatedAt }} Status: {{ rowDream.Status }} G: {{ rowDream.CountG }}
    </div>

    <div class="row-info">
      <div class="i-label">
        Information
      </div>
      <div class="i-data" >
        <span v-html="rowDream.Info"></span>

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
          <input type="number" id="energe-input" v-model="energyToDream">
          <button @click="addEnergyToDream">+{{ energyToDream }} Energy</button>
        </div>
  </div>
</div>


<div id="dreaminput">
        <h1>Create new dream!</h1>
        <form @submit.prevent="doSend">
          <label for="dreamname">Dream name</label>
          <input type="text" id="dreamname" v-model="dreamname" placeholder="..." autocomplete="off">
          <label for="dreaminfo">Dream info (Use simple html tags. No XSS Cross-Site Scripting)</label>&nbsp;
          <textarea id="dreaminfo" v-model="dreaminfo" placeholder="..."></textarea>
          <label for="location">Location name</label>&nbsp;
          <input  id="location" v-model="location" placeholder="...">
          <button type="submit">->...</button> 
          <div class="form-group">
          <div v-if="messageErr" class="alert alert-danger" role="alert">
            {{ messageErr }}
          </div>
        </div>
        </form>
      </div>

</div>
</template>



<!-- .dreamroot {

} -->
<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;

.searchBox {
  color: rgb(0, 0, 0);

  padding: 10px;
  background-color:#ffffff ;
}

.searchBox #checkbox1, #checkbox2 {
  cursor:pointer;
  border: 1px solid rgb(219, 208, 208);
  padding: 5px;
  background-color: blueviolet;
  margin-left: 3px;
  margin-right: 7px;
}
.searchBox label {
  padding-right: 5px;
  cursor:default;
}
.searchBox #checkbox1:checked {
  background-color: #365778;
}
.searchBox #filterInput {
  padding: 10px;
  background-color: white;
  color: black;

}
.searchBox button {
  background-color: #172025;
  cursor: pointer;
  padding: 10px;
  color: whitesmoke;
}



.dreamroot{

}


::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  color: clr.$clr-table-header;
  background-color: clr.$bg-table-header;
  border-color: #172025;
}
::v-deep(.vtl-table td),
::v-deep(.vtl-table tr) {
  border: 1px solid clr.$clr-table-header;
}
::v-deep(.vtl-paging-info) {
  color: clr.$bg-table-header;
}
::v-deep(.vtl-paging-count-label),
::v-deep(.vtl-paging-page-label) {
  color: clr.$bg-table-header;
}
::v-deep(.vtl-paging-pagination-page-link) {
  border: 1px solid clr.$clr-table-header;
}


#dreamrow {
  display: flex;
  width: 100%;
  flex-direction: column;
  border: 1px solid whitesmoke;
  padding: 20px;
}

#dreamrow div {
  padding-top: 10px;
}

#dreamrow .row-info .i-data{
background-color: #ffffff;
}


.control-dream-panel {
  display: flex;
  width: 100%;
  flex-direction: column;
  border: 1px solid whitesmoke;
  padding: 20px;
}
.control-dream-panel div {
  padding-top: 15px;
  padding-bottom: 15px;
  border-top: 1px solid whitesmoke;
}

.control-dream-panel button {
  color: azure;
  background-color: #172025;
  cursor: pointer;
  border: 1px solid #add8d8;
  padding: 10px;
  transition: background-color 2s ease-in-out;

  margin-top: 30px;
  width: 20%;
}
.control-dream-panel button:hover {
  background-color: #bdf750;
}

.control-dream-panel #energe-input{
  padding: 10px;
  border: 1px solid rgb(233, 229, 229);
}


#dreaminput {
    padding: 20px;
    border: 1px solid whitesmoke;
}

#dreaminput h1 {
    margin-bottom: 30px;
    margin-top: 30px;
    color:#2C5662;
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
background-color: rgb(250, 249, 253);
width: 100%;
padding: 20px 20px;
border: 1px solid whitesmoke;
border-radius: 4px;
}


#dreaminput form button {
  color: azure;
  background-color: #172025;
  cursor: pointer;
  border: none;
  padding: 10px;
  transition: background-color 2s ease-in-out;

  margin-top: 30px;
  width: 10%;
}
#dreaminput form button:hover {
  background-color: #bdf750;
}

</style>

