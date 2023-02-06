import Vue from "vue"
import VueRouter from "vue-router"
//import Index from "../views"
import Login from "../views/Login.vue"
import Home from "../views/Home.vue"

Vue.use(VueRouter);

const routes=[
    {
        path:"/",
        component: Home,
        meta:{
            auth:false
        },
        redirect:"/home",
        children:[
            {path:"/home",component:Home,meta:{auth:false}}
        ]
    },
    {
        path:"/login",
        name:"Login",
        component:Login,
    },
];

const router=new VueRouter({
    routes,
});

export default router
