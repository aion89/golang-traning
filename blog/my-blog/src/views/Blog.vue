<template>
  <div class="blog">
    <h1>{{ msg }}</h1>
    <h2>Blog</h2>
  </div>
</template>

<script>
export default {
  name: 'Blog',
  data () {
    return {
      msg: 'Welcome to Blog page'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.blog {
  min-height: 800px;
  padding-top: 200px;
}
</style>
