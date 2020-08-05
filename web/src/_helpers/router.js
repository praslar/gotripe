import Vue from 'vue';
import Router from 'vue-router';

import LoginPage from '../login/LoginPage'
import RegisterPage from '../register/RegisterPage'
import NotFoundComponent from '../app/404.vue'
import HomePage from '../home/HomePage'
Vue.use(Router);

export const router = new Router({
  mode: 'history',
  routes: [
    { path: '/', component: HomePage, beforeEnter(to,from,next) {
      const loggedIn = localStorage.getItem('user');
      if (!loggedIn) {
        return next('/login');
      }
      next();
    } },
    { path: '/login', component: LoginPage },
    { path: '/register/', component: RegisterPage },
    { path: '/', component: LoginPage },
    // otherwise redirect to home
    { path: '*/', component: NotFoundComponent }
  ]
});
