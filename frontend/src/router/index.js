import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/material',
    name: 'MaterialLibrary',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/material/MaterialList.vue')
  },
  {
    path: '/material/graphics/edit',
    name: 'MaterialGraphicsEdit',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/material/MaterialGraphicsEdit.vue')
  },
  // {
  //   path: '/template',
  //   name: 'GraphicsTemplateEdit',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/material/MaterialTemplateEdit.vue')
  // },
  // {
  //   path: '/message/list',
  //   name: 'MessageMasonry',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageMasonry.vue')
  // },
  {
    path: '/message/user-show',
    name: 'MessageUserShow',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageUserShow.vue')
  },
]

const router = new VueRouter({
  routes,
  mode: 'history',
})

export default router
