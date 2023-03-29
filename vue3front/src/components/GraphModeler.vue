
<template>
<div class="graph-panel">
  <div class="containerDraw">
      <p>Command: Ctrl+[C, V, Z, X, A, Shift+Z], backspace(delete), zoom</p>
      <div ref="container"></div>
  </div>
  <div ref="stencilref" id="nodebar"></div>
  <div class="control-panel">
      <button @click="graphToJson()" >ToJson</button> 
  </div>
</div>

<div v-if="Object.keys(selected_cell).length !== 0" class="cell-panel">
    isEdge:{{ selected_cell.value.shape==='edge' }}

    <div v-if="selected_cell.value.shape==='edge'" class="edgeForm">
      {{ selected_cell.value }}
      <input  type="text">
      <button @click="updateEdge"></button>
    </div>

    <div v-else class="nodeForm">
      {{ selected_cell.value.attrs.text.text }}
      <input v-model="newLabel" type="text" :placeholder="selected_cell.value.attrs.text.text" >
      <button @click="updateNode">Update</button>
    </div>

</div>

<div v-else class="hhh">
  Select node or edge
</div>

</template>


<script setup>
import { ref, onMounted, reactive} from "vue"
import { Graph } from '@antv/x6'
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


///////   https://x6.antv.antgroup.com/api/model/model
        // https://x6.antv.vision/en/docs/tutorial/intermediate/events
const container = ref(null)
const stencilref = ref(null)
const selected_cell = reactive({});
let graph = null

// Change data model
const newLabel = ref("")

const updateNode = () => {
  console.log(selected_cell.value.getAttrs())

  selected_cell.value.setAttrs({
  label: { text: newLabel.value },
})
  //const nodes = graph.getNodes()
  
}

// validation and upload to server
const graphToJson = () => {
    console.log("sadsdasd")
    console.log(graphRef.value.toJSON())
}

// init GraphModeler
onMounted(() => { 
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
  title: 'Формы',
  target: graph,
  stencilGraphWidth: 200,
  stencilGraphHeight: 300,
  collapsable: false,
  groups: [
    {
      title: 'Формы1',
      name: 'group1',
    },
  ],
  layoutOptions: {
    columns: 2,
    columnWidth: 100,
    rowHeight: 100,
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
Graph.registerNode(
  'custom-rect',
  {
    inherit: 'rect',
    width: 66,
    height: 36,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#5F95FF',
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
Graph.registerNode(
  'custom-polygon',
  {
    inherit: 'polygon',
    width: 66,
    height: 36,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#5F95FF',
        fill: '#EFF4FF',
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
Graph.registerNode(
  'custom-circle',
  {
    inherit: 'circle',
    width: 45,
    height: 45,
    attrs: {
      body: {
        strokeWidth: 1,
        stroke: '#5F95FF',
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
Graph.registerNode(
  'custom-image',
  {
    inherit: 'rect',
    width: 52,
    height: 52,
    markup: [
      {
        tagName: 'rect',
        selector: 'body',
      },
      {
        tagName: 'image',
      },
      {
        tagName: 'text',
        selector: 'label',
      },
    ],
    attrs: {
      body: {
        stroke: '#5F95FF',
        fill: '#5F95FF',
      },
      image: {
        width: 26,
        height: 26,
        refX: 13,
        refY: 16,
      },
      label: {
        refX: 3,
        refY: 2,
        textAnchor: 'left',
        textVerticalAnchor: 'top',
        fontSize: 12,
        fill: '#fff',
      },
    },
    ports: { ...ports },
  },
  true,
)

const r1 = graph.createNode({
  shape: 'custom-rect',
  label: "qwe",
  payload: {qwe:""},
  attrs: {
    body: {
      rx: 20,
      ry: 26,
    },
  },
})


const r2 = graph.createNode({
  shape: 'custom-rect',
  label: 'TaskR',
})
const r3 = graph.createNode({
  shape: 'custom-rect',
  attrs: {
    body: {
      rx: 6,
      ry: 6,
    },
  },
  label: 'TaskT',
})
const r4 = graph.createNode({
  shape: 'custom-polygon',
  attrs: {
    body: {
      refPoints: '0,10 10,0 20,10 10,20',
    },
  },
  label: 'TaskS',
})
const r5 = graph.createNode({
  shape: 'custom-polygon',
  attrs: {
    body: {
      refPoints: '10,0 40,0 30,20 0,20',
    },
  },
  label: 'TaskL',
})
const r6 = graph.createNode({
  shape: 'custom-circle',
  label: 'TaskY',
})
stencil.load([r1, r2, r3, r4, r5, r6], 'group1')
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
</script>



<style>
.graph-panel {
  display: flex;
  flex-direction: row;
}

#nodebar {
    position: relative;
    width: 20%;
    height: 670px;
    border: 2px dashed black;
}

.containerDraw {
    width: 70%;
    height: 670px;
    border: 1px dashed black;
    background-color: rgb(29, 5, 5);
}
.containerDraw p {
    color:whitesmoke;

}

.control-panel {
    width: 10%;
    background-color: rgb(29, 5, 5);
}

.control-panel button {

    background-color: white;
    padding: 20px;
}

.cell-panel{


  padding: 20px;
  border: 1px solid black;


}

.cell-panel input {

  border: 1px solid black;
}

.cell-panel button {
  border: 1px solid black;
  background-color: blue;
  padding: 5px;


}

.edgeForm, .nodeForm {
  display: flex;
  flex-direction: column;



}





</style>