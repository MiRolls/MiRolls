// noinspection NpmUsedModulesInstalled

import Vue from 'vue'
import App from './App.vue'
import VueRouter from "vue-router";
import router from "./router"
import VueCookies from "vue-cookies"
// import

Vue.use(VueRouter);
Vue.use(VueCookies);

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
