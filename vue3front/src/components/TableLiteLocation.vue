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
    label: "UpdatedAt",
    field: "UpdatedAt",
    width: "3%",
    sortable: true,
    display: (row) => {
        if (row.UpdatedAt) {
          return (row.UpdatedAt.slice(0, 19));
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
    label: "Geo",
    field: "Geolocation",
    width: "1%",
    sortable: true,
    display: (row) => {
        if (row.Geolocation) {
          if (row.Geolocation.length>30) {
            return (row.Info.slice(0, 31)+"...")
          };
          return (row.Geolocation);
        } else {
          return ("<b>Empty</b>")
        };
},
    },
    {
    label: "Active",
    field: "Active",
    width: "1%",
    sortable: false,
    },
    {
    label: "Radius",
    field: "Radius",
    width: "1%",
    sortable: false,
    },
    {
    label: "Height",
    field: "Height",
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

const onlyMyLocations = ref(true)
const searchTerm = ref("")
 // 
const doSearch = (offset, limit, order, sort) => {
  var searchData = {
    Offset: offset,
    Limit: limit,
    Order: order,
    Sort: sort,
    onlyMyLocations: onlyMyLocations.value,
    SearchTerm: searchTerm.value
    }
  console.log(JSON.stringify(searchData))
  table.isLoading = true;
  let url = '/locations';
  API.get(url, {params: searchData} ).then((response) => {
      if (response.data.status === "ok") {
        table.isLoading = false;
        // refresh table rows
        table.rows = response.data.data.Locations;
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

//doSearch(0, 10, "id", "asc");

const rowLocation = reactive({
    IdFiles: "",
    Height:0,
    Radius: 0,
    CreatedAt: "",
    Creater: 0,
    Energy: 0,
    ID: 0,
    Info: "",
    Geolocation: "",
    Name: "",
    UpdatedAt: "",
    Active: "",
})

const rowClicked = (row) => {
  Object.assign(rowLocation,toRaw(row) );
};


const deleteLocation = () => {
  if (rowLocation.Name ==="") {
    window.alert("Error: Row is empty!");
    return
  }
  let url = '/locations/'+rowLocation.ID;
  API.delete(url).then((response) => {
    if (response.data.status == "ok") {
      table.rows = table.rows.filter(function(elem) {
          if (elem.Name == rowLocation.Name) {return false;} else {return true;}
      });
      rowLocation.Name = "";
    }
    window.alert(response.data.message);

  });
};

const messageErr = ref("")
const locationName = ref("");
const locationInfo = ref("");
const locationGeo = ref("");
const locationHeight= ref("");
const locationRadius= ref("");


const newLocation = reactive({
    Name: locationName,
    Info: locationInfo,
    Geolocation:locationGeo,
    Height: locationHeight,
    Radius:locationRadius
})
const doSend = () => API.post("/locations", JSON.stringify(newLocation)).then((response) => {
    if (response.data.status === "ok") {
      doSearch(0, 10, "id", "asc")
      window.alert(response.data.message)
      return
    }
    window.alert(response.data.message)

})


const energyToLocation = ref(0)
const addEnergyToLocation = () => {
  if (energyToLocation.value === 0) {
    window.alert("add zero energy???")
    return
  }
  API.post("/locations/"+rowLocation.ID+"/energy", JSON.stringify({Energy: energyToLocation.value})).then((response) => {
    if (response.data.status === "ok") {
      rowLocation.Energy += energyToLocation.value
      doSearch(0, 10, "id", "asc")
      return
    }
    window.alert(response.data.message)

})
}


const mapurl = ref("https://static-maps.yandex.ru/1.x/?ll=39.620070,53.753630&spn=0.002,0.002&size=300,300&l=map")
const getMapImage = () => {

  console.log("https://static-maps.yandex.ru/1.x/?ll=" + locationGeo.value+ "&spn=0.002,0.002&size=300,300&l=map")
  mapurl.value = "https://static-maps.yandex.ru/1.x/?ll=" + locationGeo.value+ "&spn=0.02,0.02&size=300,300&l=map"
  

}

</script>

<template scoped>

<div class="searchBox">
  <label for="checkbox1">Only my locations: {{ onlyMyLocations }}</label>
  <input type="checkbox" id="checkbox1" v-model="onlyMyLocations" />

  <label for="filterInput">SearchBy:</label>
  <input id="filterInput" v-model="searchTerm" />
  <button @click="doSearch(0, 10, 'id', 'asc')">GO</button>
</div>


<div class="root">
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


  
  

<div v-if='rowLocation.Name !==""'>
  <div id="locationrow">
    <div class="row-name">  Name: {{ rowLocation.Name }}</div>
    <div class="row-location">Geolocation: {{ rowLocation.Geolocation }}</div>
    <div class="row-radiusr">Radius: {{ rowLocation.Radius }}</div>
    <div class="row-height">Height: {{ rowLocation.Height }}</div>
    <div class="row-creater">Creater: {{ rowLocation.Creater }}</div>
    <div class="row-energy">Energy: {{ rowLocation.Energy }}</div>

    <div class="row-other">
      ID: {{ rowLocation.ID }} UpdatedAt: {{ rowLocation.UpdatedAt }} CreatedAt: {{ rowLocation.CreatedAt }} Active: {{ rowLocation.Active }}
    </div>

    <div class="row-info">
      <div class="i-label">
        Information
      </div>
      <div class="i-data" >
        <span v-html="rowLocation.Info"></span>

      </div>

    </div>
  </div>
  <div class="control-location-panel">
        <h1>Локация</h1>  
        <div>
          <p>Удалить локацию.</p>
          <button @click="deleteLocation">Delete Location</button>
        </div>
        
        <div>
          <p>Повысить энергетический статус локации</p>
          <input type="number" id="energe-input" v-model="energyToLocation">
          <button @click="addEnergyToLocation">+{{ energyToLocation }} Energy</button>
        </div>
  </div>
</div>

<div id="locationinput">
        <h1>Create new location!</h1>
        <form @submit.prevent="doSend">
          <label for="locationName">Location name</label>
          <input type="text" id="locationName" v-model="locationName" placeholder="..." autocomplete="off">
          <label for="locationInfo">Location Info (Use simple html tags. No XSS Cross-Site Scripting)</label>&nbsp;
          <textarea id="locationInfo" v-model="locationInfo" placeholder="..."></textarea>

          <div class="geo">
            <div>
              <label for="locationGeo">Geolocation format: 39.620070,53.753630 
                <button type="button" @click="getMapImage">check</button>
              </label>
              <input type="text" id="locationGeo" v-model="locationGeo" placeholder="..." autocomplete="off">
            </div>
          <img v-bind:src=mapurl>
          </div>
          

          
          <label for="locationHeight">Location Height (m)</label>
          <input type="number" id="locationHeight" v-model="locationHeight" placeholder="..." autocomplete="off">
          <label for="locationRadius">Location Radius (m)</label>
          <input type="number" id="locationRadius" v-model="locationRadius" placeholder="..." autocomplete="off">



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

<style scoped>


.geo {
  display: flex;
flex-direction: row;

}
/* .locationroot {

} */


.searchBox {

  border: 1px solid white;
  padding: 10px;
  background-color:whitesmoke ;
}



.searchBox #checkbox1, #checkbox2 {
  cursor:pointer;
  border: 1px solid black;
  padding: 5px;
  background-color: blueviolet;
  margin-left: 3px;
  margin-right: 7px;
}


.searchBox label {
  color: #365778;
  padding-right: 5px;
  cursor:default;
}

.searchBox #checkbox1:checked {
  background-color: #365778;
}

.searchBox #filterInput {
  padding: 10px;
  background-color: white;

}

.searchBox button {
  color: azure;
  background-color: #172025;
  cursor: pointer;
  border: 1px solid #add8d8;
  padding: 10px;


}

::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  color: whitesmoke;
  background-color: #365778;
  border-color: #172025;
}
::v-deep(.vtl-table td),
::v-deep(.vtl-table tr) {
  border: 1px solid whitesmoke;
}
::v-deep(.vtl-paging-info) {
  color: #17324d;
}
::v-deep(.vtl-paging-count-label),
::v-deep(.vtl-paging-page-label) {
  color: #17324d;
}
::v-deep(.vtl-paging-pagination-page-link) {
  border: 1px solid whitesmoke;
}


#locationrow {
  display: flex;
  width: 100%;
  flex-direction: column;
  border: 1px solid whitesmoke;
  padding: 20px;
}

#locationrow div {
  padding-top: 10px;
}

#locationrow .row-info .i-data{
background-color: #ffffff;
}


.control-location-panel {
  display: flex;
  width: 100%;
  flex-direction: column;
  border: 1px solid whitesmoke;
  padding: 20px;
}
.control-location-panel div {
  padding-top: 15px;
  padding-bottom: 15px;
  border-top: 1px solid whitesmoke;
}

.control-location-panel button {
  color: azure;
  background-color: #172025;
  cursor: pointer;
  border: 1px solid #add8d8;
  padding: 10px;
  transition: background-color 2s ease-in-out;

  margin-top: 30px;
  width: 20%;
}
.control-location-panel button:hover {
  background-color: #bdf750;
}

.control-location-panel #energe-input{
  padding: 10px;
  border: 1px solid rgb(233, 229, 229);
}


#locationinput {
    padding: 20px;
    border: 1px solid whitesmoke;
}

#locationinput h1 {
    margin-bottom: 30px;
    margin-top: 30px;
    color:#2C5662;
}

#locationinput form {
    display: flex;
    width: 100%;
    flex-direction: column;
}

#locationinput form label {
    margin-top: 20px;

}

#locationinput form input, textarea {
background-color: rgb(250, 249, 253);
width: 100%;
padding: 20px 20px;
border: 1px solid whitesmoke;
border-radius: 4px;
}


#locationinput form button {
  color: azure;
  background-color: #172025;
  cursor: pointer;
  border: none;
  padding: 10px;
  transition: background-color 2s ease-in-out;

  margin-top: 30px;
  width: 10%;
}
#locationinput form button:hover {
  background-color: #bdf750;
}

</style>