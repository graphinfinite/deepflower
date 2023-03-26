
<template>
    <div>
        <div ref="container"></div>
    </div>





</template>


<script setup>

import { ref, onMounted} from "vue"

import { Graph, Node, Shape, Edge } from '@antv/x6'
const container = ref(null)

const graph = ref(null)

onMounted(() => {
        // the DOM element will be assigned to the ref after initial render
        console.log(container.value) // <div>This is a root element</div>

        const graph = new Graph({
            container: container.value,
            width: 1100,
            height: 600,
            panning: true,


            grid: true,

            background: {
                color: 'whitesmoke'
            },

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

                // https://x6.antv.vision/en/docs/tutorial/intermediate/interacting/#%E8%BF%9E%E7%BA%BF%E8%A7%84%E5%88%99
                allowBlank: false,
                allowMulti: false,
                allowLoop: false,
                allowEdge: false,

                allowNode: false,
                allowPort: true,

                highlight: true,
                snap: true,

                // https://x6.antv.vision/en/examples/showcase/practices#dag
                connector: 'algo-connector',
                connectionPoint: 'boundary',
                anchor: 'top',
                selecting: {
                        enabled: true,
                        multiple: false,
                        
                }
            }
        });


        graph.addNode({
            //shape: "vue3-shape",
             x: 100,
            y: 100,
            width: 150,
             height: 100,
             // attrs: {
             //   body: {       //     stroke: "#ebebeb",
             //   },
            // },
             });


      })





</script>