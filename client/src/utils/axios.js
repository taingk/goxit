import axios from 'axios';
import store from '@/store';

let instance = axios.create({
  baseURL: 'http://localhost:8080'
});

instance.interceptors.request.use(config => {
  config.headers.common['Authorization'] = store.state.token;
  config.headers.common['AccessLevel'] = store.state.accessLevel;
  return config;
});

export default instance;
