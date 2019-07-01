<template>
  <div class="main">
    <h1>{{ msg }}</h1>
    <h2>Main Page</h2>
  </div>
</template>

<script>
export default {
  name: 'MainPage',
  data () {
    return {
      msg: 'Welcome to Main Page'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.main {
  min-height: 800px;
  height: auto !important;
  padding-top: 200px;
}
</style>
