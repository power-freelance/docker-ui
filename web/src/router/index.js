import Vue from 'vue'
import Router from 'vue-router'

import store from '@/store'

import DashboardPage from '@/components/pages/DashboardPage'

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: DashboardPage,
      meta: {
        title: 'Dashboard'
      }
    }
  ]
});

router.beforeResolve((to, from, next) => {
  store.commit('ui/setProgress', true);
  next();
});

router.afterEach(() => {
  store.commit('ui/setProgress', false);
});

export default router;