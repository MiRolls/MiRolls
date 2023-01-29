import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import VueCookies from 'vue-cookies'
import {createRouter,createWebHashHistory} from 'vue-router'
import MakePage from "./pages/MakePage.vue";
import IndexPage from "./pages/IndexPage.vue";
import SearchPage from "./pages/SearchPage.vue";
import zh from "./assets/langs/zh";
import en from "./assets/langs/en";
import {createI18n} from "vue-i18n";
import QueryPage from "./pages/QueryPage.vue";

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
    },{
        path:"/query",
        component: QueryPage,
    }]
});

fetch("/get/site",{
    method:"POST"
}).then(res => res.json()).then(data=>{
    let lang = data.lang;2
    const i18n = new createI18n({
        locale: lang,
        messages:{
            en,
            zh,
        }
    })
    const app =createApp(App);
    app.use(router);
    app.use(VueCookies);
    app.use(i18n)
    app.mount('#app');
}).catch(()=>{
    const i18n = new createI18n({
        locale: "zh",
        messages:{
            en,
            zh,
        }
    })

    const app =createApp(App);
    app.use(router);
    app.use(VueCookies);
    app.use(i18n)
    app.mount('#app');
})

// const app =createApp(App);
// app.use(router);
// app.use(VueCookies);
// app.mount('#app');
// Dev mode