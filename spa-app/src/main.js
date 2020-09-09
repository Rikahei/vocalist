import Vue from 'vue';
import Vuex from 'vuex'
import App from './App.vue';
import vuetify from './plugins/vuetify';
import router  from './router.js';
import Cookies from 'js-cookie';

Vue.config.productionTip = false

new Vue({
  Vuex,
  Cookies,
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
