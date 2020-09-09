import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from './components/Home';
import Login from './components/Login';
import Oauth2success from './components/Oauth2success';

Vue.use(VueRouter);

const routes = [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/oauth2success', component: Oauth2success}
  ]
const router = new VueRouter({
    routes:routes
});

export default router;
