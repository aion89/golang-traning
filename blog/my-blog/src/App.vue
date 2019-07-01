<template>
  <div id="app">
    <SpeedDial/>
    <router-view/>
    <Footer/>
  </div>
</template>

<script>
import SpeedDial from './components/common/SpeedDial'
import Footer from './components/common/Footer'
export default {
  name: 'App',
  components: {
    SpeedDial, Footer
  }
}
</script>

<style>
html{ height:100%; }
body{ min-height:100%; padding:0; margin:0; position:relative; }
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #FFFFFF;
  align-items: center;
}
h1, h2 {
  font-weight: normal;
  margin-block-start: 0em;
  margin-block-end: 0em
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
