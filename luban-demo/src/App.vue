<template>
  <div id="app">
    <!--5. isRouterAlive刷新页面-->
    <router-view v-if="isRouterAlive"></router-view>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {
  },
  methods: {
    // 定义刷新页面方法
    reload(){
      this.isRouterAlive = false;
      this.$nextTick(function () {
        this.isRouterAlive = true
      })
    }
  },
  data() {
    return {
      isRouterAlive: true
    }
  }
}
</script>

<style>
</style>
