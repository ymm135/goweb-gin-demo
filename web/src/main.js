import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import Antd from 'ant-design-vue'
// 原ant-design css
import 'ant-design-vue/dist/antd.css'
// 自定义全局主题样式
import './theam/ant-design-vue.js'
import cryptos from './utils/crypto'
import global_variable from "./utils/global_variable";
// 路由
import router from './router'
// vuex
import store from './store/index.js'
import Vuex from 'vuex'

Vue.config.productionTip = false;

const instance = axios.create({
  baseURL: '/week'
});

// 配置过滤response
instance.interceptors.response.use((response) => {
  if (response) {
    if (response.data.code === 'FALSE') {
      return 'FALSE';
    }
    return response;
  }
}, error => {
  return Promise.reject(error);
});

// 跨域请求时是否需要使用凭证
instance.defaults.withCredentials = true;

instance.defaults.headers.post['Content-Type'] = 'application/json';
instance.defaults.headers.get['Content-Type'] = 'application/json';

// 配置请求头
instance.defaults.headers.common['x-token'] = sessionStorage.getItem('login');
instance.defaults.headers.common['x-user-id'] = sessionStorage.getItem('userId');

Vue.prototype.$axios = instance;
Vue.use(Antd);
Vue.prototype.cryptos = cryptos;
Vue.prototype.BASEURL = global_variable.baseURL;
Vue.use(Vuex);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
