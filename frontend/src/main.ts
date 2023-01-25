// import { createApp } from 'vue'
// import './style.css'
// import App from './App.vue'
//
// createApp(App).mount('#app')

import {createApp, provide} from 'vue'
import App from './App.vue'
import './style.css'
import VueCookies from 'vue-cookies'
import {createRouter,createWebHashHistory} from 'vue-router'
import MakePage from "./pages/MakePage.vue";
import IndexPage from "./pages/IndexPage.vue";
import SearchPage from "./pages/SearchPage.vue";
// import

const router = createRouter({
    history:createWebHashHistory(),
    routes:[{
        path:"/make",
        component: MakePage
    },{
        path:"",
        component: IndexPage
    },{
        path:"/search",
        component: SearchPage,
    }]
});

const app =createApp(App);
provide("cookies",VueCookies)
app.use(router);
app.use(VueCookies);
app.config.globalProperties.$cookies = app.config.globalProperties.$cookies = app.config.globalProperties.$cookies || VueCookies;
app.mount('#app');
// mount
