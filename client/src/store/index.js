import Vue from 'vue';
import Vuex from 'vuex';
import axios from '@/utils/axios';
import VuexPersist from 'vuex-persist';

Vue.use(Vuex);

const vuexPersist = new VuexPersist({
  key: 'goxit',
  storage: window.localStorage
});

export default new Vuex.Store({
  plugins: [vuexPersist.plugin],
  state: {
    token: null,
    accessLevel: 0
  },
  mutations: {
    authenticate(state, payload) {
      state.token = payload.JWT;
      state.accessLevel = payload.AccessLevel;
      axios.defaults.headers.common['Authorization'] = state.token;
      axios.defaults.headers.common['AccessLevel'] = state.accessLevel;
    }
  }
});
