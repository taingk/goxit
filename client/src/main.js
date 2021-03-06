import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './registerServiceWorker';
import toasted from 'vue-toasted';

Vue.config.productionTip = false;

Vue.use(toasted, {
  action: {
    text: 'x',
    onClick: (e, toastObject) => {
      toastObject.goAway(0);
    }
  },
  iconPack: 'fontawesome'
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
