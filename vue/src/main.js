import Vue from "vue";
import router from "./router";
import ElementUI from "element-ui";
import App from './App.vue';
import 'element-ui/lib/theme-chalk/index.css';
import {bubbles} from 'vue-canvas-effect';
import axios from "axios"
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import { BootstrapVue, BootstrapVueIcons} from 'bootstrap-vue'
import * as echarts from 'echarts';
Vue.prototype.$echarts = echarts;


Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)

axios.defaults.baseURL="http://localhost:8082"
Vue.component('bubbles-effect', bubbles);


//配置请求拦截器，用于在访问后端服务器时携带token令牌
axios.interceptors.request.use(config =>{
    let requestUrl = config.url
    let start = requestUrl.indexOf("/user")
    if (start == 0) {   // 如果访问后台，需要带上token
        config.headers.Authorization = window.sessionStorage.getItem("token")
    }

    return config           //必须return config
})

// 配置响应拦截器，该拦截器用于拦截后端数据的响应
axios.interceptors.response.use(config => {
    if(config.status == 500) {
        router.push("/internalError");
    } else if(config.status == 404) {
        router.push("/notFound");
    }

    if(config.data.status == 300) {    // 返回码为300时，说明需要登录，才能访问后台
        router.push("/login");
    }

    return config;
})

Vue.prototype.$axios = axios

Vue.config.productionTip = false


Vue.use(ElementUI);
Vue.config.productionTip=false;

new Vue({
    router,
    render:(h)=>h(App),
}).$mount("#app");

