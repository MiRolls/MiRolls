import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import VueCookies from 'vue-cookies'
import {createRouter,createWebHashHistory} from 'vue-router'
import MakePage from "./pages/MakePage.vue";
import IndexPage from "./pages/IndexPage.vue";
import SearchPage from "./pages/SearchPage.vue";

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
app.use(router);
app.use(VueCookies);
app.mount('#app');
// mount