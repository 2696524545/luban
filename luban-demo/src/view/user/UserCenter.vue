<template>
  <div>
    <button @click="goUserDetail">跳转到用户详情</button>
<!--    <router-link to="{ name: 'UserDetail', query: {id: 100}}"></router-link>-->
  </div>
</template>

<script>
export default {
  name: "UserCenter",
  methods: {
    goUserDetail() {
      this.$router.push(
          {
            name: "UserDetail",
            query: {id :100}
          }
      )
    }
  }
}
</script>

<style scoped>

</style>