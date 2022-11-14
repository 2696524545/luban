<template>
  <div>
    <slot name="slot1" :data1="data1">slot1 default value</slot>
  </div>
</template>

<script>
export default {
  name: "One",
  data() {
    return {
      data1: "子组件slot1",
    }
  }
}
</script>

<style scoped>

</style>