
<template>
    <div class="graph-panel">
        <div class="containerDraw">
            <p>Command: Ctrl+[C, V, Z, X, A, Shift+Z], backspace(delete), zoom</p>
            <div ref="container"></div>
        </div>
        <div ref="stencilref" id="nodebar"></div>
        <div class="control-panel">

            <button @click="graphToJson()" >tojson</button>            
        </div>
    </div>
</template>


<script setup>
import { ref, onMounted, reactive} from "vue"
import { Graph, Shape } from '@antv/x6'
 import { Stencil } from '@antv/x6-plugin-stencil'
 import { Transform } from '@antv/x6-plugin-transform'
 import { Selection } from '@antv/x6-plugin-selection'
 import { Snapline } from '@antv/x6-plugin-snapline'
 import { Keyboard } from '@antv/x6-plugin-keyboard'
 import { Clipboard } from '@antv/x6-plugin-clipboard'
 import { History } from '@antv/x6-plugin-history'
//import insertCss from 'insert-css'
const container = ref(null)
const stencilref = ref(null)
const graphRef = reactive({})


const graphToJson = () => {
    console.log("sadsdasd")
    console.log(graphRef.value.toJSON())
}

onMounted(() => { 
console.log("onmount")
// #region 
graphRef.value = new Graph({
  container: container.value,
  //autoResize: true,
  height: 650,
  grid: true,
      panning:true,
  mousewheel: {
    enabled: true,

    zoomAtMousePosition: true,
    modifiers: 'ctrl',
    minScale: 0.5,
    maxScale: 3,
  },
  scroller: {
                enabled: true,
                pannable: true,
  },
  connecting: {
    router: 'manhattan',
    connector: {
      name: 'rounded',
      args: {
        radius: 8,
      },
    },
    anchor: 'center',
    connectionPoint: 'anchor',
    allowBlank: false,
    snap: {
      radius: 20,
    },
    createEdge() {
      return new Shape.Edge({
        attrs: {
          line: {
            stroke: '#A2B1C3',
            strokeWidth: 2,
            targetMarker: {
              name: 'block',
              width: 12,
              height: 8,
            },
          },
        },
        zIndex: 0,
      })
    },
    validateConnection({ targetMagnet }) {
      return !!targetMagnet
    },
  },
  highlighting: {
    magnetAdsorbed: {
      name: 'stroke',
      args: {
        attrs: {
          fill: '#5F95FF',
          stroke: '#5F95FF',
        },
      },
    },
  },
})



const graph = graphRef.value



            
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
// #endregion

// #region 初始化 stencil
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
    // {
    //   title: '系统设计图',
    //   name: 'group2',
    //   graphHeight: 250,
    //   layoutOptions: {
    //     rowHeight: 70,
    //   },
    // },
  ],
  layoutOptions: {
    columns: 2,
    columnWidth: 100,
    rowHeight: 100,
  },
})
stencilref.value.appendChild(stencil.container)
// #endregion

// #region 快捷键与事件
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

// undo redo
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

// select all
graph.bindKey(['meta+a', 'ctrl+a'], () => {
  const nodes = graph.getNodes()
  if (nodes) {
    graph.select(nodes)
  }
})

// delete
graph.bindKey('backspace', () => {
  const cells = graph.getSelectedCells()
  if (cells.length) {
    graph.removeCells(cells)
  }
})

// zoom
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

// #region 初始化图形
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
  label: 'TaskW',
  payload: {qwe:""},
  attrs: {
    body: {
      rx: 20,
      ry: 26,
    },
  },
})

console.log(r1.store.data.payload)
console.log(r1.store.data.attrs.text.text)
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

//return { graph }


// mount end
})









// this.graph.toJSON()Get all the contents of the node in the current canvas
// JSON.stringify(this.graph.toJSON())You can convert all the node contents of the current canvas into JSON String save to local or background
// JSON.parse(json); hold json The data is shaped into a data format and then passed fromJSON Method is then rendered onto the canvas.
// this.graph.fromJSON(json); You can get it from the background or locally json The data is shaped and rendered on the canvas.



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
</style>