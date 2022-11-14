<template>
  <div>
    <h2 @click="getUserId">用户ID {{userId}}</h2>
  </div>
</template>

<script>
export default {
  name: "UserDetail",
  created() {

  },
  methods: {
    getUserId() {
      console.log(this.$route.query.id)
      this.userId = this.$route.query.id
    }
  },
  data() {
    return {
      userId: undefined,
    }
  }
}
</script>

<style scoped>

</style>