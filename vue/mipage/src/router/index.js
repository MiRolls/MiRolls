// noinspection NpmUsedModulesInstalled
import VueRouter from "vue-router";
import MakePage from "@/pages/MakePage.vue";
import IndexPage from "@/pages/IndexPage.vue";
import SearchPage from "@/pages/SearchPage.vue";

//创建路由器
export default new VueRouter({
    routes:[{
        path:"/make",
        component: MakePage
    },{
        path:"",
        component: IndexPage
    },{
        path:"/search",
        component: SearchPage,
    }],
})
