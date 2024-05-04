import { createApp } from 'vue'
import App from './App.vue'
import naive from "naive-ui"
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/theme-chalk/index.css'
// import { upload } from './api/test_services'

createApp(App).use(store).use(router).use(naive).use(ElementPlus).mount('#app')
