<template>
  <div>
    <button @click="listPods">跳转集群管理</button>

  </div>
</template>

<script>
import {ListPods} from "@/api/demo";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Nodes",
  methods: {
     toManage() {
       this.$router.push({
         name: 'ClusterManage'
       })
     },
    listPods(clusterName, namespace) {
       clusterName = "test"
      namespace = "default"
      ListPods(clusterName, namespace).then(res => {
        console.log(res, "pods all")
      })
    },

  }
}
</script>

<style scoped>

</style>