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
    sortable: false,
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
    sortable: true,
    },
    {
    label: "R",
    field: "Radius",
    width: "1%",
    sortable: true,
    },
    {
    label: "H",
    field: "Height",
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

const onlyMyLocations = ref(true)
const searchTerm = ref("")
 // 
const doSearch = (offset, limit, order, sort) => {
  var searchData = {
    Offset: offset,
    Limit: limit,
    Order: order,
    Sort: sort,
    OnlyMyLocations: onlyMyLocations.value,
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

doSearch(0, 10, "id", "asc");

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
  getRowMapImage()
  showLocationDream()
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
const locationHeight= ref(0);
const locationRadius= ref(0);


const newLocation = reactive({
    Name: locationName,
    Info: locationInfo,
    Geolocation:locationGeo,
    Height: locationHeight,
    Radius:locationRadius
})
const doSend = () => {
  if (locationGeo.value === "") {
    locationGeo.value = "0,0"
  }
  if (locationName.value === "") {
    window.alert("location name is empty")
    return
  }

  if (locationGeo.value !== "") {
    var arrayll = locationGeo.value.split(",", 2)
    if (arrayll.length == 1) {
      window.alert("no valid geodata")
      return
    }  
    if (
      parseFloat(arrayll[0])<90 &&
      parseFloat(arrayll[0])>-90 &&
      parseFloat(arrayll[0])<180 &&
      parseFloat(arrayll[0])>-180
      ) 
      {
            API.post("/locations", JSON.stringify(newLocation)).then((response) => {
            if (response.data.status === "ok") {
              doSearch(0, 10, "id", "asc")
              window.alert(response.data.message)
              return
            }
            window.alert(response.data.messag)
            })
    } else {
      window.alert("Неправильный формат геолокации")
    }
  }
}

const energyToLocation = ref(1)
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
const rowmapurl = ref("")
const getRowMapImage = () => {
  rowmapurl.value = "https://static-maps.yandex.ru/1.x/?ll=" + rowLocation.Geolocation.slice(1,-1)+ "&spn=0.02,0.02&size=300,300&l=map"
}
const mapurl = ref("https://static-maps.yandex.ru/1.x/?ll=39.620070,53.753630&spn=0.002,0.002&size=300,300&l=map")
const getMapImage = () => {
  console.log("https://static-maps.yandex.ru/1.x/?ll=" + locationGeo.value+ "&spn=0.002,0.002&size=300,300&l=map")
  mapurl.value = "https://static-maps.yandex.ru/1.x/?ll=" + locationGeo.value+ "&spn=0.02,0.02&size=300,300&l=map"
}

var dreams = ref(new Array())
var isVisibleLocationDreams = ref(false)

const showLocationDream = () => {
  API.get("/locations/"+rowLocation.ID+"/dreams").then((response) => {
    if (response.data.status === "ok") {
      dreams.value = Array()
      if (response.data.data.LocationDreams == null) {
        return
      }
      dreams.value.push(...response.data.data.LocationDreams)
      return
    }
    window.alert(response.data.message)
})

}

</script>

<template scoped>
<div class="searchBox">
  <label for="checkbox1">Only my locations: {{ onlyMyLocations }}</label>
  <input type="checkbox" id="checkbox1" v-model="onlyMyLocations" />

  <label for="filterInput">SearchByName:</label>
  <input id="filterInput" v-model="searchTerm" />
  <button @click="doSearch(0, 10, 'id', 'asc')">ᐅ</button>
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


<div class="location-row-control" v-if='rowLocation.Name !==""'>
  <div id="locationrow">
    <div class="row-name">{{ rowLocation.Name }} <span id="row-active" v-if="rowLocation.Active">✓</span></div>

    
    <div v-if="rowLocation.Geolocation != '(0,0)'">
      <img v-bind:src=rowmapurl alt="Геолокация не определена" width="300" height="300">
    </div>
    <div  class="row-location">Geolocation: {{ rowLocation.Geolocation }}</div>
    <div class="row-radiusr">Radius: {{ rowLocation.Radius }}</div>
    <div class="row-height">Height: {{ rowLocation.Height }}</div>
    <div class="row-creater">Creater: {{ rowLocation.Creater }}</div>
    <div class="row-createdat">
      CreatedAt: {{ rowLocation.CreatedAt }}
    </div>
    <div class="row-updatedat">UpdatedAt: {{ rowLocation.UpdatedAt }}</div>
    <div class="row-energy">Energy: {{ rowLocation.Energy }}</div>
    

    <div class="row-info">
      <div class="i-label">
        Information
      </div>
      <div class="i-data" >
        <span v-html="rowLocation.Info"></span>
      </div>
    </div>
    <button type="checkbox" id="checkbox3" v-on:click="() => {isVisibleLocationDreams=!isVisibleLocationDreams}">Dreams</button>
  </div>

  <div class="control-location-panel">
        <h1>Control</h1>  
        <div>
          <p>Удаляя локацию, вы удаляете также связи её с мечтами. Сами мечты остаются без изменения.</p>
          <button @click="deleteLocation">Delete Location</button>
        </div>
        
        <div>
          <p>Повысить энергетический статус</p>
          <input type="number" id="energe-input" min="1" step="1" v-model="energyToLocation">
          <button @click="addEnergyToLocation">+{{ energyToLocation }} Energy</button>
        </div>
  </div>
</div>

<div v-if="isVisibleLocationDreams" class="location_dreams">
    <ul>
      <li v-for="dream in dreams" :key="dream.ID">
      ID: {{ dream.ID }} <br>
      Name: {{ dream.Name }} <br>
      Energy: {{ dream.Energy }} <br>
      Creater: {{ dream.Creater }}
      </li>
    </ul>
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
          <input type="number" step="1" id="locationHeight" v-model="locationHeight" placeholder="..." autocomplete="off" >
          <label for="locationRadius">Location Radius (m)</label>
          <input type="number" id="locationRadius" minlength="0" step="1" v-model="locationRadius" placeholder="..." autocomplete="off" min="0">
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

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;


.geo {
display: flex;
flex-direction: row;
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

.searchBox button {
  color: clr.$clr-button;
  background-color:clr.$bg-button;
  cursor: pointer;
  padding: 10px;
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

.location-row-control{
  display: flex;
  flex-direction: row;
  border-top: 7px solid #0B0410;
}

#locationrow {
  width: 80%;
  padding: 20px;
}

#locationrow .row-name{
  font-size: 23px;
  margin-bottom: 15px;

}

#row-active{
  margin-left: 10px;
  color: green;
}
#locationrow div {
  margin-top: 10px;
}




#checkbox3 {
  color:blueviolet;
  cursor:pointer;
  padding: 10px;
  margin-top: 30px;
}
.location_dreams {
  width: 100%;
  border: 1px solid whitesmoke;
  padding: 20px;
  background-color: whitesmoke;  /* aliceblue */
}
.location_dreams ul li{
  background-color: white;
  padding:15px;
  border: 1px solid rgb(200, 205, 226);
}



.control-location-panel {
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
  width: 20%;
  padding: 20px;
}

.control-location-panel h1{
  font-size: 20px;
}
.control-location-panel div {
  padding-top: 15px;
  padding-bottom: 15px;
}

.control-location-panel button {
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
  background-color:clr.$bg-button;
  color: clr.$clr-button;
  cursor: pointer;
  padding: 10px;
  transition: 0.5s;
}
.control-location-panel button:hover {
  box-shadow: 0px 0px 5px rgba(60, 41, 75, 0.5);
  background-color:clr.$bg-button-hover;
}

.control-location-panel #energe-input{
  margin-top: 10px;

  padding: 10px;
  border: 1px solid rgb(233, 229, 229);
}


#locationinput {
    padding: 20px;
    border-top: 7px solid #0B0410;
    border-bottom: 20px solid #0B0410;
}

#locationinput h1 {
    margin-top: 30px;
    margin-bottom: 30px;
    color:clr.$clr-button;
    font-size: 20px;
}

#locationinput form {
    display: flex;
    width: 100%;
    flex-direction: column;
    margin-bottom: 10px;
}

#locationinput form label {
    margin-bottom: 20px;

}

#locationinput img{
  border: 3px solid rgb(165, 160, 160);
  margin-left: 20px;
  margin-top:  10px;

  width: 150px;
  height: 150px;
}

#locationinput form input, textarea {
width: 100%;
padding: 20px 20px;
border: 1px solid whitesmoke;
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




</style>