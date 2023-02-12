import Vue from "vue"
import VueRouter from "vue-router"
import Index from "../views"
import Login from "../views/Login.vue"
import Home from "../views/Home.vue"
import Register from "../views/Register.vue"
import SiteDetail from "../views/SiteDetail.vue";
import Tags from "../views/Tags.vue"
import Info from "../views/users/Info.vue";
import UserInfo from "../views/users/infos/UserInfo.vue";

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
            {path: "sitedetails",component: SiteDetail,meta: {auth: false}},
            {path:"tags",component: Tags,meta: {auth: false}}
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
    {
      path: "/info",
      name:"Info",
      component: Info,
        meta:{auth: false},
        children:[
            {path: "/userinfo",component: UserInfo,meta: {auth: true}}
        ]
    },
];

const router=new VueRouter({
    base:process.env.BASE_URL,
    routes,
});

export default router




