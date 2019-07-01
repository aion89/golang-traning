<template>
  <div class="speeddial">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <md-speed-dial class="md-top-left" md-direction="bottom">
     <md-speed-dial-target @click="main">
        <md-icon>home</md-icon>
      </md-speed-dial-target>

      <md-speed-dial-content>
        <md-button class="md-icon-button" @click="blog">
          <md-icon>thumb_up</md-icon>
        </md-button>

        <md-button class="md-icon-button" @click="about">
          <md-icon>mail</md-icon>
        </md-button>
      </md-speed-dial-content>
    </md-speed-dial>
  </div>
</template>

<script>
import MenuIcon from 'vue-material-design-icons/Menu.vue'
import router from '../../router/index'
export default {
  name: 'SpeedDial',
  components: {
    MenuIcon
  },
  methods: {
    main () {
      router.push({ name: 'MainPage' })
    },
    about () {
      router.push({ name: 'About' })
    },
    blog () {
      router.push({ name: 'Blog' })
    }
  }
}
</script>

<style scoped>

</style>
