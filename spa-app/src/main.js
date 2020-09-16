import Vue from 'vue';
import Vuex from 'vuex'
import App from './App.vue';
import vuetify from './plugins/vuetify';
import router  from './router.js';
import Cookies from 'js-cookie';
import VueYoutube from 'vue-youtube';

Vue.config.productionTip = false

Vue.use(VueYoutube)

new Vue({
  Vuex,
  Cookies,
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
