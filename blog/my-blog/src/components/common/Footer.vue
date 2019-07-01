<template>
  <v-footer dark class='footer' height=auto absolute>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <v-card class="flex" flat tile >
      <v-card-title class="teal">
        <strong class="subheading">{{ footer_message }}</strong>
        <v-spacer></v-spacer>
      </v-card-title>
    <v-card-text>
      <v-btn v-for='icon in icons' :key='icon' dark icon @click="open(icon)">
        <v-icon size="24px" :class='icon'></v-icon>
      </v-btn>
    </v-card-text>

      <v-card-actions class="grey darken-3 justify-center">
        &copy;2019<strong></strong>
      </v-card-actions>
    </v-card>
  </v-footer>
</template>

<script>
import router from '../../router/index'
export default {
  name: 'Footer',
  data: () => ({
    footer_message: 'Footer Message',
    icons: [
      'fa fa-facebook',
      'fa fa-twitter',
      'fa fa-google-plus',
      'fa fa-linkedin',
      'fa fa-instagram'
    ]
  }),
  methods: {
    open (item) {
      switch (item) {
        case 'fa fa-facebook':
          window.location.href('https://www.facebook.com/profile.php?id=100000869414682')
          break

        case 'fa fa-instagram':
          router.push('Facebook')
          break

        default:
          break
      }
    }
  }
}
</script>

<style scoped>
.footer {
    bottom: 0;
    width:100%;
}
</style>
