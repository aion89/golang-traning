<template>
  <div class="about">
    <MySwiper/>
  </div>
</template>

<script>
import MySwiper from '@/components/about/MySwiper'
export default {
  name: 'About',
  components: {
    MySwiper
  },
  data () {
    return {
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.about{
  min-height: 1200px;
}
</style>
