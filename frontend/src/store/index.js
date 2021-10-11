import Vue from 'vue'
import Vuex from 'vuex'
import {graphicsHttp} from "../plugins/axios";
import Cookies from "js-cookie";

Vue.use(Vuex)

const TOKEN_KEY = 'graphics_token';
const config = require('../config')
const getUserInfo = () => {
    return graphicsHttp('post', 'user/info');
}

const login = () => {
    return new Promise(resolve => {
        resolve({
            data: getGraphicsHttp()
        })
    })
}

export const getGraphicsHttp = () => {
    let url = location.search;
    let httpRequest = {};
    if (url.indexOf('?') !== -1) {
        let strs = url.substr(1).split('&');
        for (let i = 0; i < strs.length; i++) {
            httpRequest[strs[i].split('=')[0]] = unescape(strs[i].split('=')[1]);
        }
    }
    return httpRequest;
}

const logout = () => {
    return graphicsHttp('get', '/logout')
};

const {cookieExpires} = config;
const setToken = (token) => {
    Cookies.set(TOKEN_KEY, token, {expires: cookieExpires || 1})
};


export const getToken = () => {
    const token = Cookies.get(TOKEN_KEY)
    if (token) return token
    else return false
}

export default new Vuex.Store({
    state: {
        user: {
            userName: '',
            userId: '',
            avatarImgPath: '',
            token: getToken(),
            access: [],
            hasGetInfo: false,
            roles: [],
            rolesValues: [],
            subjects: [],
            menu: []
        },
    },
    mutations: {
        setAvatar(state, avatarPath) {
            state.user.avatarImgPath = avatarPath
        },
        setUserId(state, id) {
            state.user.userId = id
        },
        setUserName(state, name) {
            state.user.userName = name
        },
        setToken(state, token) {
            state.user.token = token
            setToken(token)
        },
        setHasGetInfo(state, status) {
            state.user.hasGetInfo = status
        },
        setMenu(state, menu) {
            state.user.menu = menu;
        },
    },
    actions: {
        // 登录
        handleLogin({commit}) {
            return new Promise((resolve, reject) => {
                login({}).then(res => {
                    if (res.data.token) { // 当有token返回时
                        commit('setToken', res.data.token)
                    }
                    resolve()
                }).then(() => {
                    graphicsHttp('post', '/menu/list').then(res => {
                        const menu = res.data.results;
                        commit('setMenu', menu);
                        // localStorage.setItem('menu', JSON.stringify(menu));
                        resolve()
                    }).catch(err => {
                        reject(err)
                    });
                }).catch(err => {
                    console.log(err)
                    reject(err)
                })
            })
        },
        // 退出登录，请求后端服务器
        handleLogout({state, commit}) {
            return new Promise((resolve, reject) => {
                logout(state.token).then(() => {
                    commit('setToken', '')
                    commit('setAccess', [])
                    commit('setHasGetInfo', false)
                    commit('setAvatar', "")
                    localStorage.clear();
                    // console.log('access clear');
                    resolve()
                }).catch(err => {
                    reject(err)
                })
            })
        },
        // 获取用户相关信息
        getUserInfo({state, commit}) {
            return new Promise((resolve, reject) => {
                try {
                    getUserInfo(state.token).then(res => {
                        const data = res.data;
                        console.log(data)
                        if (!data.id && !data.name) {
                            location.href = '/login';
                        }
                        commit('setAvatar', data.avatar)
                        commit('setUserName', data.name)
                        commit('setUserId', data.id)
                        commit('setAccess', data.access)
                        commit('setHasGetInfo', true)
                        resolve(data)
                    }).catch(err => {
                        reject(err)
                    })
                } catch (error) {
                    console.log(error)
                    reject(error)
                }
            })
        },

    },
    modules: {}
})
