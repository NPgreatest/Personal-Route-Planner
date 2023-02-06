import Vue from "vue"
import VueRouter from "vue-router"
import Index from "../views"
import Login from "../views/Login.vue"
import Home from "../views/Home.vue"
import Register from "../views/Register.vue"
import SiteDetail from "../views/SiteDetail.vue";

Vue.use(VueRouter);

const routes=[
    {
        path:"/",
        component: Index,
        meta:{
            auth:false
        },
        redirect:"/home",
        children:[
            {path:"/home",component:Home,meta:{auth:false}},
            {path: "sitedetails",component: SiteDetail,meta: {auth: false}}
        ]
    },
    {
        path:"/login",
        name:"Login",
        component:Login,
    },
    {
        path:"/register",
        name:"Register",
        component:Register,
    },
];

const router=new VueRouter({
    routes,
});

export default router
