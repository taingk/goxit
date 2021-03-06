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

const preventGuest = next => {
  if (store.state.token) {
    next();
  } else {
    next({ name: 'home' });
  }
};

const preventUser = next => {
  if (store.state.token) {
    next({ name: 'list-votes' });
  } else {
    next();
  }
};

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    redirect: '/login'
  },
  {
    path: '/vote',
    name: 'create',
    component: VoteCreate,
    beforeEnter(to, from, next) {
      preventGuest(next);
    }
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
    ],
    beforeEnter(to, from, next) {
      preventGuest(next);
    }
  },
  {
    path: '/votes',
    name: 'list-votes',
    component: VoteList,
    beforeEnter(to, from, next) {
      preventGuest(next);
    }
  },
  {
    path: '/user/:uuid?',
    name: 'show-user',
    component: UserShow,
    beforeEnter(to, from, next) {
      preventGuest(next);
    }
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
    beforeEnter(to, from, next) {
      preventUser(next);
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    beforeEnter(to, from, next) {
      preventUser(next);
    }
  },
  {
    path: '/logout',
    name: 'logout',
    component: Logout,
    beforeEnter(to, from, next) {
      preventGuest(next);
    }
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
