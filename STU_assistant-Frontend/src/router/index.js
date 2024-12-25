import HomeView from '@/views/HomeView.vue';
import UserLogin from '@/views/UserLogin.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { isAuthenticated } from '@/utils/auth';


const routes = [
  {
    path:'/',
    redirect:'/login'
  },
  {
    path: '/login',
    name: 'login',
    component: UserLogin,
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

// 全局路由守卫
router.beforeEach((to, from, next) => {
  if (to.path !== '/login' && !isAuthenticated()) {
    // 如果用户未认证，且不是在访问登录页，则重定向到登录页
    next({ path: '/login' })
  } else {
    // 否则继续导航
    next()
  }
})

export default router
