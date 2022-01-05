import store from "../store";
import goTo from "vuetify/es5/services/goto";
import {apiUrlDomain} from "../plugins/axios";
import VueRouter from 'vue-router';
import Vue from 'vue'
import {getToken} from "../store/module/user"

Vue.use(VueRouter);

const routes = [
    {
        path: '/material',
        name: 'MaterialLibrary',
        meta: {
            title: '素材库',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/material/MaterialList.vue')
    },
    {
        path: '/material/graphics/edit',
        name: 'MaterialGraphicsEdit',
        meta: {
            title: '编辑素材',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/material/MaterialGraphicsEdit.vue')
    },
    {
        path: '/channel/edit',
        name: 'EditChannel',
        meta: {
            title: '编辑频道',
            authCheck: true
        },
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
        path: '/message-owner',
        name: 'MessageOwner',
        meta: {
            title: '消息管理列表',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageOwnerList.vue')
    },
    {
        path: '/message-user',
        name: 'MessageUser',
        meta: {
            title: '用户消息列表',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageUserList.vue')
    },
    {
        path: '/channel',
        name: 'ChannelList',
        meta: {
            title: '频道列表',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/channel/ChannelList.vue')
    },
    {
        path: '/user',
        name: 'UserList',
        meta: {
            title: '用户列表',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/user/UserList.vue')
    },
    {
        path: '/role',
        name: 'RoleList',
        meta: {
            title: '角色列表',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/role/RoleList.vue')
    },
    {
        path: '/permission',
        name: 'PermissionList',
        meta: {
            title: '权限列表',
            authCheck: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/permission/PermissionList.vue')
    },
    {
        path: '/message-user/show',
        name: 'MessageUserShow',
        meta: {
            title: '用户消息详情',
            authCheck: true,
            hideNav: true
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageUserShow.vue')
    },
    {
        path: '/message-owner/show',
        name: 'MessageOwnerShow',
        meta: {
            title: '消息详情',
            authCheck: true,
        },
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/message/MessageOwnerShow.vue')
    },
    {
        path: '/login',
        name: 'Login',
        meta: {
            title: '登录',
            hideNav: true
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/login/oauth.vue')
    },
    {
        path: '/menu',
        name: 'Menu',
        meta: {
            title: '菜单',
            authCheck: true
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/menu/index.vue')
    },
    {
        path: '/example',
        name: 'Example',
        meta: {
            title: '示例',
            authCheck: false,
            hideNav: true
        },
        component: () => import(/* webpackChunkName: "about" */ '../views/example/ShowMe.vue')
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

const setActiveMenu = function (to) {
    let path = '/'+(to.path.split('/')[1] || '')
    if (path === '/') {
        return;
    }

    const menuList = store.state.user.menu;
    if (!menuList){
        return;
    }
    for (let i = 0; i < menuList.length; i++) {
        let skip = false;
        if (menuList[i].children) {
            for (let j = 0; j < menuList[i].children.length; j++) {
                if (menuList[i].children[j].path === path) {
                    store.state.user.activeParentMenuId = i;
                    store.state.user.activeSubMenuId = j;
                    skip = true;
                    break;
                }
            }
        }
        if (skip) {
            break;
        }
    }
}

const router = new VueRouter(RouterConfig)
const loginPage = apiUrlDomain + '/oauth/login';
Vue.prototype.$loginPage = loginPage;

router.beforeEach((to, from, next) => {
    const accessList = store.state.user.access;
    if (to.meta.hideNav) {
        store.state.user.showNav = false;
    } else {
        store.state.user.showNav = true;
    }
    if (to.meta.authCheck) {
        if (store.state.user.hasGetInfo && store.state.user.token && getToken() && getToken() === store.state.user.token) {
            //查看是否有权限
            if (accessList.indexOf(to.path) > -1) {
                setActiveMenu(to);
                next()
            } else if (to.path !== '/login') {
                next({
                    path: '/401'
                })
            } else {
                setActiveMenu(to);
                next()
            }
        } else {// 未登录需要登录
            window.location.href = loginPage;
        }
    } else {
        setActiveMenu(to);
        next()
    }
})

export default router;
