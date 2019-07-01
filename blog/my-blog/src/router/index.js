import Vue from 'vue'
import Router from 'vue-router'
import About from '@/views/About'
import Admin from '@/views/Admin'
import Blog from '@/views/Blog'
import MainPage from '@/views/MainPage'
import NotFound from '@/views/404/NotFound'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'MainPage',
      component: MainPage
    },
    {
      path: '/about',
      name: 'About',
      component: About
    },
    {
      path: '/admin',
      name: 'Admin',
      component: Admin
    },
    {
      path: '/blog',
      name: 'Blog',
      component: Blog
    },
    {
      path: 'https://www.facebook.com/profile.php?id=100000869414682',
      name: 'Facebook'
    },
    {
      path: '*',
      name: 'NotFound',
      component: NotFound,
      meta: {
        title: 'Page Not Found'
      }
    }
  ]
})
