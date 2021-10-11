import Vue from 'vue'
import './plugins/axios'
import './plugins/toast'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import store from './store'
import VueCropper from 'vue-cropper'

Vue.use(VueCropper)
Vue.config.productionTip = false

const setMenuToLocal = function (item, submenu) {
    localStorage.setItem('menuItem', item);
    localStorage.setItem('submenuActive', submenu);
};
Vue.prototype.setMenuToLocal = setMenuToLocal;

new Vue({
    vuetify,
    router: router,
    store: store,
    render: h => h(App)
}).$mount('#app')
