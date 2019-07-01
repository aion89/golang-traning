<template>
  <div class="admin">
    <h1>{{ msg }}</h1>
    <h2>Admin</h2>
  </div>
</template>

<script>
export default {
  name: 'Admin',
  data () {
    return {
      msg: 'Welcome to Admin page'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.blog {
  min-height: 700px;
}
</style>
