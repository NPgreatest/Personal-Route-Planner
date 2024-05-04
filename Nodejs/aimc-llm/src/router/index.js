import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AppView from '../views/AppView.vue'
import AppCreate from '../views/AppCreate.vue'
import AppEdit from '../views/AppEdit.vue'
import AppUse from '../views/AppUse.vue'
import AppDetail from '@/views/AppDetail.vue'



const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: AppView,
  },
  {
    path: '/create',
    name:'create',
    component:AppCreate
  },
  {
    path:'/edit/:_id',
    name:'edit',
    component:AppEdit
  },
  {
    path:'/use/:_id',
    name:'use',
    component:AppUse
  },
  {
    path:'/detail/:_id',
    name:'detail',
    component:AppDetail
  }
]

// export default new router({
//   mode: 'history', // 尝试改为 'hash' 如果你遇到路由问题
//   routes
// })


const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
