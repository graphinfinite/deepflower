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
          if (row.Name.length>25) {
            return (row.Name.slice(0, 25)+"...")
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
const doSend = () =>{

  newdream.Name =newdream.Location +"‗"+ newdream.Name
  API.post("/dreams", JSON.stringify(newdream)).then((response) => {
    if (response.data.status === "ok") {
      newdream.Name=""
      newdream.Info=""
      doSearch(0, 10, "id", "asc")
      window.alert(response.data.message)
      return
    }
    window.alert(response.data.message)

});


} 


const energyToDream = ref(1)
const addEnergyToDream = () => {
  if (energyToDream.value === 0) {
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

  <label for="filterInput">Search by location/dream:</label>
  <input id="filterInput" v-model="searchTerm" />
  <button @click="doSearch(0, 10, 'id', 'asc')">ᐅ</button>
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
  

<div class="dream-row-panel" v-if='rowDream.Name !==""'>
  <div id="dreamrow">
    <div class="row-name">{{ rowDream.Name }}<span v-if="!rowDream.Published" id="row-published">no published</span> </div>
    <div class="row-energy">Energy: {{ rowDream.Energy }}</div>
    <div v-if="rowDream.Published">PublishAt: {{ rowDream.PublishAt }}</div>
    <div class="row-creater">Creater: {{ rowDream.Creater }}</div>
    <div>CreatedAt: {{ rowDream.CreatedAt }}</div>


    <div class="row-status">  Status:{{ rowDream.Status }} </div>
    <div>G: {{ rowDream.CountG }}</div>

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
        <h1>Control</h1>  
        <div>
          <p>После публикации мечту нельзя будет изменить!</p> 
          <p>На публикацию расходуется 1 энергия.</p>
          <button @click="deleteDream">Delete Dream</button>
          <button @click="publishDream">Publish Dream</button>
        </div>
        
        <div>
          <p>Добавить энергию мечте</p>
          <input type="number" id="energe-input" min="1" step="1" v-model="energyToDream">
          <button @click="addEnergyToDream">+{{ energyToDream }} Energy</button>
        </div>
  </div>
</div>


<div id="dreaminput">
        <h1>Create new dream!</h1>
        <form @submit.prevent="doSend">
          <label for="location">Location name</label>&nbsp;
          <input  id="location" v-model="location" placeholder="...">
          <label for="dreamname">Dream name</label>
          <input type="text" id="dreamname" v-model="dreamname" placeholder="..." autocomplete="off">
          <label for="dreaminfo">Dream info (Use simple html tags. No XSS Cross-Site Scripting)</label>&nbsp;
          <textarea id="dreaminfo" v-model="dreaminfo" placeholder="..."></textarea>

          <button type="submit">Save</button> 
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
  padding: 10px;
}

.searchBox #checkbox1, #checkbox2 {
  cursor:pointer;
  padding: 5px;
  background-color: blueviolet;
  margin-left: 5px;
}
.searchBox #checkbox1:checked {
  background-color: #365778;
}
.searchBox label {
  cursor:default;
  margin-left: 10px;
}
.searchBox input {
  margin-left: 10px;
  border: 1px solid whitesmoke;
  padding: 10px;
  background-color: white;
  color: black;

}
.searchBox button {
  color: clr.$clr-button;
  background-color:clr.$bg-button;
  cursor: pointer;
  padding: 10px;
}

.dreamroot{
}


::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  color: clr.$clr-table-header;
  background-color: clr.$bg-table-header;
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


.dream-row-panel{
  display: flex;
  flex-direction: row;
  border-top: 7px solid #0B0410;
}


#dreamrow {

  width: 80%;
  padding: 20px;
}


#dreamrow #row-published {
  font-size: 10px;
  margin-left: 5px;
  color: rgb(235, 11, 78);

}

#dreamrow .row-name {
  font-size: 20px;
  margin-bottom: 20px;

}
#dreamrow div {
  padding-top: 10px;
}



.control-dream-panel {
  width: 20%;
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
  padding: 20px;
}
.control-dream-panel div {
  padding-top: 15px;
  padding-bottom: 15px;
  border-top: 1px solid whitesmoke;
}

.control-dream-panel input{
  margin-top: 10px;
  padding:10px;
}

.control-dream-panel h1{

  font-size: 20px;
}



button {
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
  background-color:clr.$bg-button;
  cursor: pointer;
  padding: 10px;
  transition: 0.5s;
  margin-top: 30px;
  color: clr.$clr-button;
  margin-left:5px;
}
button:hover {
  box-shadow: 0px 0px 5px rgba(60, 41, 75, 0.5);
  background-color:clr.$bg-button-hover;
}



///  new
#dreaminput {
    padding: 20px;
    border-top: 7px solid #0B0410;
    border-bottom: 20px solid #0B0410;
}

#dreaminput h1 {
    margin-bottom: 30px;
    margin-top: 30px;
    color:clr.$clr-button;
    font-size: 20px;
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
//background-color: rgb(250, 249, 253);
width: 100%;
padding: 20px 20px;
border: 1px solid whitesmoke;
}

</style>

