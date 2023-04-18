
<template scoped>

<div class="root">

<!-- SEARCH AND TABLE -->
<div class="searchBox">
<label for="checkbox1">Only my projects: {{ onlyMyProjects }}</label>
<input type="checkbox" id="checkbox1" v-model="onlyMyProjects" />

<label for="filterInput">Search by dream/project_name:</label>
<input id="filterInput" v-model="searchTerm" />
<button @click="doSearch(0, 10, 'id', 'asc')">GO</button>
</div>

<div class="projectstable">
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
</div>
<!-- END SEARCH AND TABLE -->



<div class="graph-panel">


  <!-- GRAPH MODELER -->
  <div class="containerDraw">
      <p>Hello from graphinfinit! Command: Ctrl+[C, V, Z, X, A, Shift+Z], backspace(delete), zoom</p>
      <div ref="container"></div>
  </div>

  <div ref="stencilref" id="nodebar"></div>
  <!-- END GRAPH MODELER -->


  <!-- PROJECT -->

  <div class="project-panel" v-if=showRowProject>

    <div>
      <button v-if="!rowProject.Published" @click="showRowProject=!showRowProject" class="clonebutton">Clone</button>
      <button @click="showRowProject=!showRowProject;rowProject.value={}; rowProject.Published=false; graph.fromJSON({}); selected_cell.value={}" class="redbutton">Close</button>
    </div>
    <div></div>

    
    <div class="project-id">PROJECT ID: {{ rowProject.ID }}</div>
    <div class="project-name">Name: 
      {{rowProject.Name}}
    </div>
    <div class="project-description">Description: 
      {{rowProject.Info}}
    </div>
    <div class="project-published">Published: 
      {{rowProject.Published}} 
      <div class="if-not-published" v-if="!rowProject.Published">
        <button @click="publishProject" class="orangebutton">Publish</button>
        <button @click="deleteProject" class="redbutton">Delete</button>
      </div>
    </div>
    <div class="project-publishat">PublishAt: 
      {{rowProject.PublishAt}}
    </div>
    <div class="project-energy">Energy: 
      {{rowProject.Energy}}
    </div>
    <div class="project-status">Status: 
      {{rowProject.Status}}
    </div>
    <div class="project-creater">Creater: 
      {{rowProject.Creater}}
    </div>
    <div class="project-CreatedAt">CreatedAt: 
      {{rowProject.CreatedAt}}
    </div>
    <div class="project-UpdatedAt">UpdatedAt: 
      {{rowProject.UpdatedAt}}
    </div>

  </div>

  <div v-else class="project-panel">
      <div class="project-label">Create new project!</div>
      <div class="project-dream">For dream:
          <input  type="text" v-model="NewProject.DreamName">
      </div>
      <div class="project-name">Name: 
        <input  type="text" v-model="NewProject.Name">
      </div>
      <div class="project-description">Description: 
        <textarea  type="text" v-model="NewProject.Info"></textarea>
      </div>
      <div><button @click="createNewProject" id="savebutton">Validate and Save</button></div>
  </div>
  <!-- END PROJECT -->


</div>


<!-- CELL DATA AND CONTROL -->

<div v-if="Object.keys(selected_cell).length !== 0" class="cell-panel">

    <!-- EDGE -->

    <div v-if="selected_cell.value.shape==='edge'" class="edgeForm">
      <div class="edge-id">EDGE ID: {{ selected_cell.value.id }}</div>
      <label for="edge-label">EDGE LABEL</label>
      <input  type="text" v-model="edgeUpdate.Label" id="edge-label">
      <button @click="updateEdge">Update</button>
    </div>


  <!-- DATA NODES: SLOW AND FAST -->

  <div v-else-if="selected_cell.value.shape==='slow-model' || selected_cell.value.shape==='fast-model'" class="node">
    <div class="node-data">
      <div class="node-id">NODE ID: {{ selected_cell.value.id }}</div>
      <div class="node-label">Label: {{ selected_cell.value.attrs.text.text }} </div>
      <div class="node-leadtime">Lead Time: {{ selected_cell.value.data.LeadTime }}</div>
      <div class="node-status">Status: {{ selected_cell.value.data.Status }}</div>
      <div class="node-energy">Energy: {{ selected_cell.value.data.Energy }}</div>
      <div class="node-energy">Performers: {{ selected_cell.value.data.Performers }}</div>
      <div class="node-description">Description: {{ selected_cell.value.data.Description }}</div>
    </div>

    <div class="node-control">
      <div class="node-change-created" v-if="!rowProject.Published">
        <label for="node-label">Label</label>
        <input v-model="nodeUpdate.Label" type="text" :placeholder="selected_cell.value.attrs.text.text"  id="node-label">
        <label for="leadtime">Lead Time(h)</label>
        <input  type="number" v-model="nodeUpdate.LeadTime" id="node-leadtime" min="1" step="1">
        <label for="node-description">Description</label>
        <textarea  type="text" v-model="nodeUpdate.Description" id="node-description" ></textarea>

        <button @click="updateNode">Set</button>
      </div>
      <div class="node-change-publiched" v-else>
        <div class="energy-to-task-form" v-if="selected_cell.value.getData().Status =='created'">
          <input type="number" min="1" step="1" v-model="energyToTask"><button @click="addEnergyToTask">+{{ energyToTask }} Energy</button>
        </div>
        <div class="close-task-form" v-if="selected_cell.value.getData().Status =='inwork'">
          <button @click="closeTask" >DONE</button>
        </div>
        <div class="close-task-form" v-if="selected_cell.value.getData().Status =='created'">
          <button @click="grabTask" >GRAB</button>
        </div>
      </div>
    </div>
  </div>

  <!-- CHOICE NODE -->

  <div v-else-if="selected_cell.value.shape==='CHOICE' " class="node" :key="componentKey">
    <div class="node-data">
      <div class="node-id">NODE ID: {{ selected_cell.value.id }}</div>
      <div class="node-label">Label: {{ selected_cell.value.attrs.text.text }} </div>
      <div class="node-status">Status: {{ selected_cell.value.data.Status }}</div>
      <div class="node-ways">Ways: {{ selected_cell.value.data.Ways }}</div>
      <div class="node-choisenway">Chosen Way: {{ selected_cell.value.data.ChosenWay }}</div>
      <div class="node-description">Description: {{ selected_cell.value.data.Description }}</div>
    </div>

    <div class="node-control">
      <div class="node-change-created" v-if="!rowProject.Published">
        <label for="node-label">Label</label>
        <input v-model="nodeUpdate.Label" type="text" :placeholder="selected_cell.value.attrs.text.text"  id="node-label">
        <label for="node-description">Description</label>
        <textarea  type="text" v-model="nodeUpdate.Description" id="node-description" ></textarea>
        <button @click="updateNodeChoice">Set</button>


        <div class="addchoice">
          <div>
            Way:<input v-model="choiceWaysObj.Key" type="text">
            Сondition:<input v-model="choiceWaysObj.Value" type="text">
            <button @click="addChoiceWay">+</button>
          </div>
          <div class="choiceWaislist" v-for="(value, name) in selected_cell.value.data.Ways">
            <ul>
              <li >{{ name }}: {{ value }}<button id="button-simple" @click="()=>{removeNodeWay(name)}">x</button></li>
            </ul>
          </div>
        </div>
        
      </div>
      <div class="node-change-publiched" v-else> TODO</div>
    </div>
  </div>

    <!-- OTHER NODES -->

  <div v-else-if="selected_cell.value.shape==='MULTI' " class="node">
    MULTI: don`t use. After ways from CHOICE. Syntactic sugar for multiplying a subsequent graph into multiple subgraphs according to the number of multiple choice inputs.
    ID: {{ selected_cell.value.id }}
  </div>
  <div v-else-if="selected_cell.value.shape==='START' " class="node">
    START: special node for start project.
    ID: {{ selected_cell.value.id }}
  </div>
  <div v-else-if="selected_cell.value.shape==='END' " class="node">
    END: special node for end project.
    ID: {{ selected_cell.value.id }}
  </div>
  <div v-else-if="selected_cell.value.shape==='FAIL' " class="node">
    FAIL: special node after CHOICE. This path means the project can never be completed.
    ID: {{ selected_cell.value.id }}
  </div>
  <div v-else class="node">{{selected_cell.value.shape}}</div>

  <!-- END OTHER NODES -->

<!-- END IF CELL -->
</div>
<div v-else class="hh">
  Select node or edge
</div>

<!-- END CELL -->

</div>
</template>


<script setup>
import { ref, onMounted, reactive, toRaw, setBlockTracking} from "vue"
import { Graph, Shape, Cell } from '@antv/x6'
//import { Cell } from '@antv/x6'
//import { computed } from 'vue';
import "@antv/x6-vue-shape";
import { Stencil } from '@antv/x6-plugin-stencil'
import { Transform } from '@antv/x6-plugin-transform'
import { Selection } from '@antv/x6-plugin-selection'
import { Snapline } from '@antv/x6-plugin-snapline'
import { Keyboard } from '@antv/x6-plugin-keyboard'
import { Clipboard } from '@antv/x6-plugin-clipboard'
import { History } from '@antv/x6-plugin-history'
//
import { defaultGraphOptions } from '@/modules/initGraphModeler'


import TableLite from "vue3-table-lite";
import { computed } from "vue"
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
    label: "UpdDate",
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

const onlyMyProjects = ref(true)
const searchTerm = ref("")
const doSearch = (offset, limit, order, sort) => {
  var searchData = {
    Offset: offset,
    Limit: limit,
    Order: order,
    Sort: sort,
    OnlyMyProjects: onlyMyProjects.value,
    SearchTerm: searchTerm.value
    }
  console.log(JSON.stringify(searchData))
  table.isLoading = true;
  let url = '/projects';
  API.get(url, {params: searchData} ).then((response) => {
      if (response.data.status === "ok") {
        table.isLoading = false;
        // refresh table rows
        table.rows = response.data.data.Projects;
        table.totalRecordCount = response.data.data.TotalRecordCount;
        table.sortable.order = order;
        table.sortable.sort = sort;
        return
      } 
      window.alert(response.data.message);
  }); 
};
  
const tableLoadingFinish = (elements) => {
table.isLoading = false;
};

doSearch(0, 10, "id", "asc");

const rowProject = reactive({
  ID: 0,
  Name: "",
  Info: "",
  CreatedAt: "",
  UpdatedAt: "",
  Creater: 0,
  Energy: 0,
  Published: false,
  PublishAt: "",
  Status: "",
})

const showRowProject = ref(false)
const rowClicked = (row) => {
  showRowProject.value = true
  Object.assign(rowProject,toRaw(row));
  graph.fromJSON(JSON.parse(row.Graph))
};


const publishProject = () => {
  let url = '/projects/'+rowProject.ID+ '/publish';
  console.log(url)
  API.post(url).then((response) => {
    if (response.data.status === "ok") {
      rowProject.Published = true;
      doSearch(0, 10, "id", "asc") 
    }
    window.alert(response.data.message);
  });

};


const deleteProject = () => {
  let url = '/projects/'+rowProject.ID;
  API.delete(url).then((response) => {
    if (response.data.status == "ok") {
      showRowProject.value=false; graph.fromJSON({}); selected_cell.value={};
      doSearch(0, 10, "id", "asc")
      return
    }
    window.alert(response.data.message);
  });
};



const NewProject = reactive({
  DreamName: "",
  Info:"",
  Name: "",
  Graph: ""
})
const createNewProject = () => {
  NewProject.Graph = JSON.stringify(graph.toJSON())
  NewProject.Name = NewProject.DreamName +"/"+ NewProject.Name
  let url = '/projects';
  API.post(url, NewProject).then((response) => {
      if (response.data.status === "ok") {
        console.log("graph load")
        NewProject.Name = ""; NewProject.Info = ""; NewProject.Graph = ""
        window.alert("Проект успешно сохранен!");
        return
      } 
      window.alert(response.data.message);
  }); 

}

// TASK AND CHOICE CONTROL

const energyToTask = ref(0)
const addEnergyToTask = ()=> {
  let obj = selected_cell.value.getData()
  console.log(obj)
  if (obj.Status == "created") {
    let url = '/projects/'+ rowProject.ID+"/node/"+ selected_cell.value.id+"/addenergy";
    API.post(url, JSON.stringify({Energy: energyToTask.value})).then((response) => {
      if (response.data.status === "ok") {
        selected_cell.value.setData({"Energy": obj.Energy + energyToTask.value})
        doSearch(0, 10, "id", "asc")
        return
      } 
      window.alert(response.data.message);
    }); 
    return
  } 
  window.alert("error. status != created");
}


const grabTask = ()=> {
  let obj = selected_cell.value.getData()
  if (obj.Status == "created") {
    let url = '/projects/'+ rowProject.ID+"/node/"+ selected_cell.value.id+"/grab";
    API.post(url).then((response) => {
      if (response.data.status === "ok") {
        console.log("task inwork")
        selected_cell.value.setData({"Status": "inwork", "Performer":"you" })
        doSearch(0, 10, "id", "asc")
        return
      } 
      window.alert(response.data.message);
      
    }); 
    return
    
  }
  window.alert("error. status != created");
}

const closeTask = ()=> {
  let obj = selected_cell.value.getData()
  if (obj.Status == "inwork") {
    let url = '/projects/'+ rowProject.ID+"/node/"+ selected_cell.value.id+"/close";
    API.post(url).then((response) => {
      if (response.data.status === "ok") {
        console.log("task in confirmation")
        selected_cell.value.setData({"Status": "confirmation"})
        doSearch(0, 10, "id", "asc")
        return
      } 
      window.alert(response.data.message);
      
    }); 
    return
    
  }
  window.alert("first you need to grab the task");
}
// END TASK AND CHOICE CONTROL


////////
/////// GRAPH INIT
///////
const container = ref(null)
const stencilref = ref(null)
const selected_cell = reactive({});
let graph = null
onMounted(() => { 

// https://x6.antv.antgroup.com/api/model/model
// https://x6.antv.vision/en/docs/tutorial/intermediate/events
graph = new Graph({
  ...defaultGraphOptions,
  container: container.value
})
     
graph.use(
    new Transform({
      resizing: true,
      //rotating: true,
    }),
  )
  .use(
    new Selection({
      enabled: true,
      rubberband: true,
      showNodeSelectionBox: true,
    }),
  )
  .use(
    new Snapline({
      enabled: true,
    }),
  )
  .use(
    new Keyboard({
      enabled: true,
    }),
  )
  .use(
    new Clipboard({
      enabled: true,
    }),
  )
  .use(
    new History({
      enabled: true,
    }),
  )

register_events(graph)

const stencil = new Stencil({
  title: 'Конструктор',
  target: graph,
  stencilGraphWidth: 130,
  stencilGraphHeight: 600,
  collapsable: false,
  groups: [
    {
      title: 'Формы1',
      name: 'group1',
    },
  ],
  layoutOptions: {
    columns: 1,
    columnWidth: 100,
    rowHeight: 60,
  },
});
stencilref.value.appendChild(stencil.container);
// #endregion

// #region
graph.bindKey(['meta+c', 'ctrl+c'], () => {
  const cells = graph.getSelectedCells()
  if (cells.length) {
    graph.copy(cells)
  }
  return false
})
graph.bindKey(['meta+x', 'ctrl+x'], () => {
  const cells = graph.getSelectedCells()
  if (cells.length) {
    graph.cut(cells)
  }
  return false
})
graph.bindKey(['meta+v', 'ctrl+v'], () => {
  if (!graph.isClipboardEmpty()) {
    const cells = graph.paste({ offset: 32 })
    graph.cleanSelection()
    graph.select(cells)
  }
  return false
})
graph.bindKey(['meta+z', 'ctrl+z'], () => {
  if (graph.canUndo()) {
    graph.undo()
  }
  return false
})
graph.bindKey(['meta+shift+z', 'ctrl+shift+z'], () => {
  if (graph.canRedo()) {
    graph.redo()
  }
  return false
})
graph.bindKey(['meta+a', 'ctrl+a'], () => {
  const nodes = graph.getNodes()
  if (nodes) {
    graph.select(nodes)
  }
})
graph.bindKey('backspace', () => {
  const cells = graph.getSelectedCells()
  if (cells.length) {
    graph.removeCells(cells)
  }
})
graph.bindKey(['ctrl+1', 'meta+1'], () => {
  const zoom = graph.zoom()
  if (zoom < 1.5) {
    graph.zoom(0.1)
  }
})
graph.bindKey(['ctrl+2', 'meta+2'], () => {
  const zoom = graph.zoom()
  if (zoom > 0.5) {
    graph.zoom(-0.1)
  }
})

// 控制连接桩显示/隐藏
const showPorts = (ports, show) => {
  for (let i = 0, len = ports.length; i < len; i += 1) {
    ports[i].style.visibility = show ? 'visible' : 'hidden'
  }
}

graph.on('node:mouseenter', () => {
  const ports = container.value.querySelectorAll(
    '.x6-port-body',
  ) 
  showPorts(ports, true)
})
graph.on('node:mouseleave', () => {
  const ports = container.value.querySelectorAll(
    '.x6-port-body',
  ) 
  showPorts(ports, false)
})
// #endregion

const ports = {
  groups: {
    top: {
      position: 'top',
      attrs: {
        circle: {
          r: 4,
          magnet: true,
          stroke: '#5F95FF',
          strokeWidth: 1,
          fill: '#fff',
          style: {
            visibility: 'hidden',
          },
        },
      },
    },
    right: {
      position: 'right',
      attrs: {
        circle: {
          r: 4,
          magnet: true,
          stroke: '#5F95FF',
          strokeWidth: 1,
          fill: '#fff',
          style: {
            visibility: 'hidden',
          },
        },
      },
    },
    bottom: {
      position: 'bottom',
      attrs: {
        circle: {
          r: 4,
          magnet: true,
          stroke: '#5F95FF',
          strokeWidth: 1,
          fill: '#fff',
          style: {
            visibility: 'hidden',
          },
        },
      },
    },
    left: {
      position: 'left',
      attrs: {
        circle: {
          r: 4,
          magnet: true,
          stroke: '#5F95FF',
          strokeWidth: 1,
          fill: '#fff',
          style: {
            visibility: 'hidden',
          },
        },
      },
    },
  },
  items: [
    {
      group: 'top',
    },
    {
      group: 'right',
    },
    {
      group: 'bottom',
    },
    {
      group: 'left',
    },
  ],
}


// slow
Graph.registerNode(
  'slow-model',
  {
    inherit: 'rect',
    width: 66,
    height: 36,
    data: { 
      Description: "",
      LeadTime: 0,
      Performer: "",
      Energy: 0,
      Status: "",
    },
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#EE0010',
        fill: '#EFF4FF',
      },
      text: {
        fontSize: 12,
        fill: '#262626',
      },
    },
    ports: { ...ports },
  },
  true,
)
const slow = graph.createNode({
  shape: 'slow-model',
  label: 'SLOW',
  data: { 
    Description: "empty",
    LeadTime: 0,
    Performer: "",
    Energy: 0,
    Status: "created",
    },
})


// fast
Graph.registerNode(
  'fast-model',
  {
    inherit: 'rect',
    width: 66,
    height: 36,
    data: { 
      Description: "",
      LeadTime: 0,
      Performer: "",
      Energy: 0,
      Status: "",
    },
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#0089C8',
        fill: '#EFF4FF',
      },
      text: {
        fontSize: 12,
        fill: '#262626',
      },
    },
    ports: { ...ports },
  },
  true,
)
const fast = graph.createNode({
  shape: 'fast-model',
  attrs: {
    body: {
      rx: 6,
      ry: 6,
    },
  },
  label: 'FAST',
  data: { 
    Description: "empty",
    LeadTime: 0,
    Performer: "",
    Energy: 0,
    Status: "created",
    },
})



// choice
Graph.registerNode(
  'CHOICE',
  {
    inherit: 'rect',
    width: 66,
    height: 36,
    data:{
      Ways: {},
      Description: "",
      Status: "created",
      ChosenWay: "",
    },
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#3CC5FF',
        fill: '#3CC5FF',
      },
      text: {
        fontSize: 12,
        fill: '#FFFFFF',
      },
    },
    ports: { ...ports },
  },
  true,
)
const choice = graph.createNode({
  shape: 'CHOICE',
  label: "CHOICE",
  attrs: {
    body: {
      rx: 20,
      ry: 20,
    },
  },
})


// multiplex
Graph.registerNode(
  'MULTI',
  {
    inherit: 'polygon',
    width: 66,
    height: 36,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#FFD400',
        fill: '#FFD400',
      },
      text: {
        fontSize: 12,
        fill: '#262626',
      },
    },
    ports: {
      ...ports,
      items: [
        {
          group: 'top',
        },
        {
          group: 'bottom',
        },
      ],
    },
  },
  true,
)
const multi = graph.createNode({
  shape: 'MULTI',
  attrs: {
    body: {
      refPoints: '10,0 40,0 30,20 0,20',
    },
  },
  label: 'MULTI',
})


// start end
Graph.registerNode(
  'START',
  {
    inherit: 'circle',
    width: 45,
    height: 45,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#5F95FF',
        fill: '#309A05',
      },
      text: {
        fontSize: 12,
        fill: '#F5F5F5',
      },
    },
    ports: { ...ports },
  },
  true,
)
Graph.registerNode(
  'END',
  {
    inherit: 'circle',
    width: 45,
    height: 45,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#6F35F6',
        fill: '#6F35F6',
      },
      text: {
        fontSize: 12,
        fill: '#F5F5F5',
      },
    },
    ports: { ...ports },
  },
  true,
)
const start = graph.createNode({
  shape: 'START',
  label: 'START',
})
const end = graph.createNode({
  shape: 'END',
  label: 'END',
})


// fail
Graph.registerNode(
  'FAIL',
  {
    inherit: 'polygon',
    width: 66,
    height: 36,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#EE0020',
        fill: '#EE0020',
      },
      text: {
        fontSize: 12,
        fill: '#FFFFFF',
      },
    },
    ports: {
      ...ports,
      items: [
        {
          group: 'top',
        },
        {
          group: 'bottom',
        },
      ],
    },
  },
  true,
)
const fail = graph.createNode({
  shape: 'FAIL',
  attrs: {
    body: {
      refPoints: '0,10 10,0 20,10 10,20',
    },
  },
  label: 'FAIL',
})

stencil.load([start, slow, fast, choice, fail, multi, end], 'group1')
// mount end
})
const register_events = (graph) => {
        graph.on('node:click', ({ e, x, y, node, view }) => { 
          console.log(node)
          if (selected_cell.value != node)
                selected_cell.value = node
        })
        graph.on('edge:click', ({ e, x, y, edge, view }) => { 
          console.log(edge)
          if (selected_cell.value != edge)
                selected_cell.value = edge
        })
}
///////
/////// END GRAPH INIT 
//////



//  UPDATE CELL
const nodeUpdate = reactive({
  Label: "",
  LeadTime: 0, 
  Description: "",
})
const edgeUpdate = reactive({
  Label: "",
})
const componentKey = ref(0);  /// для лучшей отрисовки
const choiceWaysObj = reactive({
  Key:"",
  Value:""
})


const updateEdge = () => {
  selected_cell.value.setLabels(edgeUpdate.Label)
  edgeUpdate.Label = "";
}
const updateNode = () => {
  selected_cell.value.setAttrs({
  text: { text: nodeUpdate.Label},
  })
  selected_cell.value.setData({"Description":nodeUpdate.Description, "LeadTime":nodeUpdate.LeadTime})
  nodeUpdate.Label = "";
  nodeUpdate.LeadTime = 0;
  nodeUpdate.Description = "";
}

const addChoiceWay = ()=> {
  selected_cell.value.setData({Ways: {[choiceWaysObj.Key]:choiceWaysObj.Value}});
  choiceWaysObj.Key="";choiceWaysObj.Value="";componentKey.value += 1;
 }
const removeNodeWay = (name)=> {
  let obj = selected_cell.value.getData()
  delete obj.Ways[name]
  selected_cell.value.updateData(obj, {overwrite: true}) 
  componentKey.value += 1;
}
const updateNodeChoice = ()=> {
  selected_cell.value.setAttrs({
  text: { text: nodeUpdate.Label},
  })
  selected_cell.value.setData({"Description":nodeUpdate.Description})
  nodeUpdate.Label = "";nodeUpdate.Description = "";
}
// END UPDATE CELL



</script>




<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;

.root {
}



.searchBox {
  padding: 20px;
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
  background-color: white;
  border: 1px solid whitesmoke;
}
.searchBox button {
  padding: 10px;
}


.graph-panel {
  display: flex;
  flex-direction: row;
  margin-top: 30px;
}
#nodebar {
    position: relative;
    width: 10%;
    height: 770px;
    border-top:10px solid #1D0505;
    border-right:10px solid #1D0505;
    border-bottom:10px solid #1D0505;
}
.containerDraw {
    border-left: 5px solid #ffffff;
    width: 70%;
    height: 770px;
    background-color: rgb(29, 5, 5);
}
.containerDraw p {
    color:whitesmoke;
}




// Стиль панели создание и обновление проекта
.project-panel {
    border-bottom:10px solid #1D0505;
    width: 20%;
    max-width: 20%;
    background-color: white;
    padding: 20px;
}
.project-panel input{
  padding: 5px;
  margin-top: 5px;
  border: 1px solid whitesmoke
}

.project-panel div {
  margin-top: 20px;
}

.project-panel .project-label {

  font-size: 17px;
  color: #2f156b;
}

.redbutton{
  color: red;
  margin-left: 10px;
}
.orangebutton{
  color:darkorange;
}


#savebutton {
  color:#309A05;
  width: 100%;
  max-width: 100%;
}
//////// EnD




////// ЯЧЕЙКА
.cell-panel {
  width: 100%;
  background-color: white;
}
.node{
  padding:10px;
  display: flex;
  flex-direction: row;
}
.node .node-data {
padding: 10px;
width: 50%;
}
.node .node-data div {
  margin-top: 10px;
}
.node .node-control{
  width: 50%;
}
.node .node-control .node-change-created {
  display: flex;
  flex-direction: column;
  padding:10px;
}
.node input, textarea {
  border: 1px solid whitesmoke;
  padding: 10px;
  margin-bottom: 10px;
}


.close-task-form {

  padding:30px;

  background-color: #ede9f8;
}
.close-task-form button {

}

.energy-to-task-form{
  border: 1px solid whitesmoke;
  margin-bottom: 10px;

  padding: 20px;
}



//// разобраться с этим
label {
color: #1D0505;
margin-top: 20px;
}
button {
  cursor: pointer;
  max-width: 70px;
  padding: 3px;
  color: clr.$clr-button;
  transition: 0.5s;
  box-shadow: 0 0 10px rgba(168, 164, 172, 0.5);
}
button:hover {
  box-shadow: 0px 0px 5px rgba(60, 41, 75, 0.5);
}
#button-simple {
  cursor: pointer;
  font-size: 15px;
  color: rgb(241, 6, 6);
  border:none;
  border-radius: 10px;
  background-color: white;
  transition: 1.5s;
  margin-left: 10px;
}

.addchoice {
  margin-top: 20px;
}

.edgeForm {
  padding: 20px;
}

.edgeForm input{
  padding: 5px;
  border: 1px solid whitesmoke;
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






</style>