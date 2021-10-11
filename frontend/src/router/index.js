import store, {getToken} from "../store";
import goTo from "vuetify/es5/services/goto";
import {apiUrlDomain} from "../plugins/axios";
import VueRouter from 'vue-router';
import Vue from 'vue'

Vue.use(VueRouter);

const routes = [
    {
        path: '/material',
        name: 'MaterialLibrary',
        meta: {
            title: '素材库',
            authCheck: false
        },
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
    {
        path: '/channel/edit',
        name: 'EditChannel',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/channel/EditChannel.vue')
    },
    // {
    //   path: '/template',
    //   name: 'GraphicsTemplateEdit',
    //   // route level code-splitting
    //   // this generates a separate chunk (about.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import(/* webpackChunkName: "about" */ '../views/material/MaterialTemplateEdit.vue')
    // },
    {
        path: '/message',
        name: 'MessageMasonry',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageMasonry.vue')
    },
    {
        path: '/channel',
        name: 'ChannelList',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/channel/ChannelList.vue')
    },
    {
        path: '/user',
        name: 'UserList',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/user/UserList.vue')
    },
    {
        path: '/role',
        name: 'RoleList',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/role/RoleList.vue')
    },
    {
        path: '/message/user-show',
        name: 'MessageUserShow',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageUserShow.vue')
    },
    {
        path: '/login',
        name: 'Login',
        meta: {
            title: '登录',
            hideInMenu: true
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/login/oauth.vue')
    },
    {
        path: '/menu',
        name: 'Menu',
        meta: {
            title: '菜单',
            authCheck: false
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/menu/index.vue')
    },
]

const RouterConfig = {
    mode: 'history',
    routes: routes,
    scrollBehavior: (to, from, savedPosition) => {
        let scrollTo = 0

        if (to.hash) {
            scrollTo = to.hash
        } else if (savedPosition) {
            scrollTo = savedPosition.y
        }

        return goTo(scrollTo)
    },
};


const router = new VueRouter(RouterConfig)
const loginPage = apiUrlDomain + '/oauth/login';
Vue.prototype.$loginPage = loginPage;

router.beforeEach((to, from, next) => {
    const accessList = store.state.user.access;
    if (to.meta.authCheck) {
        if (store.state.user.hasGetInfo && store.state.user.token && getToken() && getToken() === store.state.user.token) {
            //查看是否有权限
            if (accessList.indexOf(to.path) > -1) {
                next()
            } else if (to.path !== '/login') {
                next({
                    path: '/401'
                })
            } else {
                next()
            }
        } else {// 未登录需要登录
            window.location.href = loginPage;
        }
    } else {
        next()
    }
})

export default router;
