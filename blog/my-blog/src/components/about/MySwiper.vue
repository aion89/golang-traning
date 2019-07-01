<template>
<div class="swiper">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <agile class="main" ref="main" :options="options1" :as-nav-for="asNavFor1">
      <div class="slide" v-for="(slide, index) in slides" :key="index" :class="`slide--${index}`"><img :src="slide"/></div>
    </agile>
    <agile class="thumbnails"  ref="thumbnails" :options="options2" :as-nav-for="asNavFor2">
        <div class="slide slide--thumbniail" v-for="(slide, index) in slides" :key="index" :class="`slide--${index}`" @click="$refs.thumbnails.goTo(index)"><img :src="slide"/></div>
        <template slot="prevButton"><i class="fas fa-chevron-left"></i></template>
        <template slot="nextButton"><i class="fas fa-chevron-right"></i></template>
    </agile>
</div>
</template>

<script>
import { VueAgile } from 'vue-agile'
export default {
  name: 'MySwiper',
  components: {
    agile: VueAgile
  },
  data () {
    return {
      asNavFor1: [],
      asNavFor2: [],
      options1: {
        dots: false,
        fade: true,
        navButtons: false },

      options2: {
        autoplay: true,
        centerMode: true,
        dots: false,
        navButtons: false,
        slidesToShow: 5,
        responsive: [{
          breakpoint: 5000,
          settings: {
            slidesToShow: 3 }
        },
        {
          breakpoint: 5000,
          settings: {
            navButtons: true }
        }
        ]},

      slides: [
        '/static/images/about_swiper/1.jpg',
        '/static/images/about_swiper/2.jpg',
        '/static/images/about_swiper/3.jpg',
        '/static/images/about_swiper/4.jpg',
        '/static/images/about_swiper/5.jpg'
      ]}
  },
  mounted () {
    this.asNavFor1.push(this.$refs.thumbnails)
    this.asNavFor2.push(this.$refs.main)
  }
}
</script>

<style scoped>
.swiper {
   width: 100%;
}

.main {
  margin-bottom: 30px;
  text-align: center;
}
.thumbnails {
  margin: 0 -5px;
  display: inline-block;
  width: calc(60%)
}
.agile__nav-button {
  background: transparent;
  border: none;
  color: #ccc;
  cursor: pointer;
  font-size: 24px;
  transition-duration: 0.3s;
}
.thumbnails .agile__nav-button {
  position: absolute;
  top: 50%;
  -webkit-transform: translateY(-50%);
          transform: translateY(-50%);
}
.thumbnails .agile__nav-button--prev {
  left: -45px;
}
.thumbnails .agile__nav-button--next {
  right: -45px;
}
.agile__nav-button:hover {
  color: #888;
}
.agile__dot {
  margin: 0 10px;
}
.slide {
  align-items: center;
  box-sizing: border-box;
  color: #fff;
  display: flex;
  height: 800px;
  justify-content: center;
}
.slide--thumbniail {
  cursor: pointer;
  height: 110px;
  padding: 0 5px;
  transition: opacity 1s;
}
.slide--thumbniail:hover {
  opacity: 0.75;
}
.slide img {
  height: 100%;
  width: 100%;
  -o-object-fit: cover;
     object-fit: cover;
  -o-object-position: center;
     object-position: center;

}
</style>
