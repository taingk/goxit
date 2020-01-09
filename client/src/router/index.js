import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';
import Register from '../views/Register.vue';
import Login from '../views/Login.vue';
import Logout from '../components/Logout.vue';
import Show from '../views/Show.vue';
import Edit from '../views/Edit.vue';
import List from '../views/List.vue';
import Create from '../views/Create.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/vote',
    name: 'create',
    component: Create
  },
  {
    path: '/vote/:uuid',
    name: 'show-vote',
    component: Show,
    children: [
      {
        path: 'edit',
        component: Edit
      }
    ]
  },
  {
    path: '/votes',
    name: 'list-votes',
    component: List
  },
  {
    path: '/register',
    name: 'register',
    component: Register
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/logout',
    name: 'logout',
    component: Logout
  }
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router;
