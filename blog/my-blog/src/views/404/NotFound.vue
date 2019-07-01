<template>
    <div class='notfound'>
        <h1>Page not found!!!</h1>
        <div>
            <router-link :to="{name: 'MainPage'}" class="btn btn-warning">Go Home</router-link>
        </div>
    </div>
</template>

<script>
export default {
  name: 'NotFound',
  data () {
    return {}
  }
}
</script>

<style scoped>
.notfound{
  min-height: 800px;
  padding-top: 200px;
}
</style>
