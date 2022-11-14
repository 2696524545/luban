import Vue from 'vue'
import App from './App.vue'
import axios from "axios";
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false
Vue.prototype.axios = axios;
Vue.use(ElementUI)

new Vue({
  //通过 router 配置参数注入路由，从而让整个应用都有路由功能
  router,
  render: h => h(App),
}).$mount('#app')
