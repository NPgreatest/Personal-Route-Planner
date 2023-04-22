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
import UserRate from "../views/users/infos/Rating.vue";
import Plan from "../views/Plan.vue";
 import Chat from "../views/ChatBase.vue";
 import Route from "../views/Route.vue";
 import MsgBoard from "../views/MsgBoard.vue";
import {InUseAttributeError} from "core-js/internals/dom-exception-constants";
import msgBoard from "@/views/MsgBoard";
import Summary from "../views/Summary.vue";
import VueAMap  from 'vue-amap';
Vue.use(VueAMap);
Vue.use(VueRouter);

VueAMap.initAMapApiLoader({
    key: 'a7f50827d4048c0a04ceb22a45142253',
    plugin: ['AMap.Weather','AMap.PlaceSearch','AMap.Driving'],
    v: '1.4.4'
});
window._AMapSecurityConfig = {
    securityJsCode:'1dc99b3b717938eb7873a7fe779f38cc',
}

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
            {path:"tags",component: Tags,meta: {auth: false}},
            {path:"plan",component: Plan,meta: {auth: false}},
             {path:"chat",component: Chat,meta: {auth: false}},
            {path:"msgBoard",component: msgBoard,meta: {auth: false}},
            {path:"route",component: Route,meta: {auth: false}},

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
            {path: "/userinfo",component: UserInfo,meta: {auth: true}},
            {path: "/userrate",component: UserRate,meta: {auth: true}}
        ]
    },
    {
        path:"/summary",
        name:"Summary",
        component: Summary,
    }
];

const router=new VueRouter({
    base:process.env.BASE_URL,
    routes,
});

export default router




