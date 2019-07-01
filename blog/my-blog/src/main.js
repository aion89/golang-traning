// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store/store.js'
import VueSocketIO from 'vue-socket.io'
import VueMaterial from 'vue-material'
import VueAgile from 'vue-agile'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default-dark.css'
import 'vue-material-design-icons/styles.css'
import 'vuetify/dist/vuetify.min.css'
import Vuetify from 'vuetify'

Vue.use(Vuetify)
Vue.use(VueMaterial)

Vue.config.productionTip = false
Vue.use(VueAgile)
/* eslint-disable */
Vue.use(new VueSocketIO({
  debug: true,
  connection: 'http://127.0.0.1:9000/',
  vuex: {
    store,
    actionPrefix: "socket_",
    mutationPrefix: "SOCKET_",
  }
}))

new Vue({
  render: h => h(App),
  store,
  router,
  sockets: {
    connect: function () {
        /* eslint-disable */
        console.log('socket connected')
    },
    test: function (data) {
      this.$socket.emit("emit_ping", 'test');
    },
    pong: function (data) {
      console.log('pong' + data)
    }
  }
}).$mount('#app')
