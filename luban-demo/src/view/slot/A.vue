<template>
  <div>
    <One>
      <template #slot1="{data1}">
        <h1>父组件slot1 获取子组件作用域中的数据===> {{data1}}</h1>
      </template>
    </One>
  </div>
</template>

<script>
import One from "./B.vue"
export default {
  name: "A",
  components: {
    One,
  }
}
</script>

<style scoped>

</style>