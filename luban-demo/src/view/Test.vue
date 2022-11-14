<template>
  <div>
    <!--- v-bind -->
<!--    <img alt="logo" v-bind:src="logo">-->

<!--    <ol>-->
<!--      <li v-for="pod in podList" :key="pod.id">{{pod.name}}</li>-->

<!--    </ol>-->

    <input type="text" v-model="inputValue"/>
    <p>{{inputValue}}</p>
    <button v-on:click="changData">改变数据</button>
    <h2 v-show="isShow">姓名</h2>
    <button @click="showBtn(true)">显示</button>
    <button @click="showBtn(false)">隐藏</button>

    <button @click="getNodes('test')">获取node节点</button>
  </div>
</template>

<script>
import {ListNodes} from "@/api/k8s";
import {Test} from "@/api/test";

export default {
  name: "Test.vue",
  created() {
    this.changData()
  },
  mounted() {
    this.inputValue = "mounted"
  },
  methods: {
    changData() {
      this.inputValue = "欢迎鲁班"
      console.log("created.......")
    },
    showBtn(value) {
      this.isShow = value
    },
    getNodes(clusterName) {
      Test({"clusterName": clusterName}).then(res => {
        console.log(res, "Test result")
      })
      // ListNodes(clusterName).then(res => {
      //   console.log(res)
      // })
    }
  },
  data() {
    return {
      logo: "https://ask.kuailexingqiu.fun/template/wic_random/static/logo.png",
      podList: [
        {id: 1, "name": "nginx-1"},
        {id: 2, "name": "nginx-2"},
        {id: 3, "name": "nginx-3"},
      ],
      inputValue: "hello",
      isShow: true,
    }
  }
}
</script>

<style scoped>

</style>