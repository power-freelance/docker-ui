import axios from 'axios'
import store from '@/store'

const $http = axios.create({
  baseURL: '/api/',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

$http.interceptors.request.use((config) => {
  store.commit('ui/setProgress', true);
  return config;
});

$http.interceptors.response.use((response) => {
  store.commit('ui/setProgress', false);

  return response;
}, (error) => {
  store.commit('ui/setProgress', false);

  return Promise.reject(error);
})

export default $http;
