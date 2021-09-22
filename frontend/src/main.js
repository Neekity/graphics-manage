import Vue from 'vue'
import './plugins/axios'
import './plugins/toast'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import store from './store'
import VueCropper from 'vue-cropper'
import shnUI from 'shn-vue-ui'
Vue.use(shnUI)
Vue.use(VueCropper)
Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
