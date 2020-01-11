import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '@/views/Home.vue';
import UserShow from '@/views/User/Show.vue';
import Register from '@/views/User/Register.vue';
import Login from '@/views/User/Login.vue';
import Logout from '@/components/Logout.vue';
import VoteShow from '@/views/Vote/Show.vue';
import VoteEdit from '@/views/Vote/Edit.vue';
import VoteList from '@/views/Vote/List.vue';
import VoteCreate from '@/views/Vote/Create.vue';

import store from '@/store';

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
    component: VoteCreate
  },
  {
    path: '/vote/:uuid',
    name: 'show-vote',
    component: VoteShow,
    children: [
      {
        path: 'edit',
        component: VoteEdit
      }
    ]
  },
  {
    path: '/votes',
    name: 'list-votes',
    component: VoteList
  },
  {
    path: '/user/:uuid?',
    name: 'show-user',
    component: UserShow
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

router.beforeEach((to, from, next) => {
  if (to.path === '/vote' && !store.state.accessLevel) next('/votes');
  next();
});

export default router;
