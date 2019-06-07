import Vue from 'vue'
import Router from 'vue-router'

import store from '@/store'

import DashboardPage from '@/components/pages/DashboardPage'
import ContainersPage from "@/components/pages/ContainersPage";
import ImagesPage from "@/components/pages/ImagesPage";

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
    },
    {
      path: '/containers',
      name: 'containers',
      component: ContainersPage,
      meta: {
        title: 'Containers'
      }
    },
    {
      path: '/images',
      name: 'images',
      component: ImagesPage,
      meta: {
        title: 'Images'
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