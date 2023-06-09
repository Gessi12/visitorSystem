import Vue from 'vue'
import Router from 'vue-router'

import Login from "../view/Login";
import Register from "../view/Register";
import Main from "../view/Main";
import Search from "../view/index/Search";
import Index from "../view/index/Index";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/login',
      component:Login
    },
    {
      path: '/register',
      component:Register
    },
    {
      path: '/main',
      component:Main,
      children:[
        {
          path:'/main/index',
          component:Index
        },
        {
          path:'/main/search',
          component:Search
        }
      ]
    },


  ]
})
