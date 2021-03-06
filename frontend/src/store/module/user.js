import {graphicsHttp} from "../../plugins/axios";
import Cookies from "js-cookie";

const TOKEN_KEY = 'graphics_token'
const config = require('../../config')
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

const logout = (token) => {
    return graphicsHttp('post', '/logout',{token:token})
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

export default {
    state: {
        userName: '',
        userId: '',
        avatarImgPath: '',
        token: getToken(),
        activeParentMenuId:0,
        activeSubMenuId:0,
        access: [],
        hasGetInfo: false,
        showNav: true,
        roles: [],
        rolesValues: [],
        subjects: [],
        menu: []
    },
    mutations: {
        setAvatar(state, avatarPath) {
            state.avatarImgPath = avatarPath
        },
        setUserId(state, id) {
            state.userId = id
        },
        setUserName(state, name) {
            state.userName = name
        },
        setToken(state, token) {
            state.token = token
            setToken(token)
        },
        setHasGetInfo(state, status) {
            state.hasGetInfo = status
        },
        setMenu(state, menu) {
            state.menu = menu;
        },
        setAccess(state, access) {
            state.access = access
        },
    },
    actions: {
        // ??????
        handleLogin({commit}) {
            return new Promise((resolve, reject) => {
                login({}).then(res => {
                    if (res.data.token) { // ??????token?????????
                        commit('setToken', res.data.token)
                    }
                    resolve()
                }).then(() => {
                    graphicsHttp('post', '/menu/list').then(res => {
                        const menu = res.data.data;
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
        // ????????????????????????????????????
        handleLogout({state, commit}) {
            return new Promise((resolve, reject) => {
                logout(state.token).then(() => {
                    commit('setToken', '')
                    commit('setAccess', [])
                    commit('setHasGetInfo', false)
                    commit('setAvatar', "")
                    commit('setMenu', {})
                    localStorage.clear();
                    // console.log('access clear');
                    resolve()
                }).catch(err => {
                    reject(err)
                })
            })
        },
        // ????????????????????????
        getUserInfo({state, commit}) {
            return new Promise((resolve, reject) => {
                try {
                    getUserInfo(state.token).then(res => {
                        const data = res.data.data;
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
}
