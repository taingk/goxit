import Vue from 'vue';
import Vuex from 'vuex';
import axios from '@/utils/axios';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: null,
    accessLevel: 0
  },
  mutations: {
    authenticate(state, payload) {
      state.token = payload.JWT;
      state.accessLevel = payload.AccessLevel;
      axios.defaults.headers.common['Authorization'] = state.token;
      axios.defaults.headers.common['accesslevel'] = state.accessLevel;
    }
  }
});
