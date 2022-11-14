## Vue 常用指令
1. v-bind, 可以简写为" :src ", v-bind:src

2. v-for 用于循环数组, 通常要求绑定key是唯一
```
 <li v-for="pod in podList" :key="pod.id">{{pod.name}}</li>
      podList: [
        {id: 1, "name": "nginx-1"},
        {id: 2, "name": "nginx-2"},
        {id: 3, "name": "nginx-3"},
      ]
```

3. v-model 用于数据双向绑定
```vue
    <input type="text" v-model="inputValue"/>
    <p>{{inputValue}}</p>
    <button @click="changData">改变数据</button>
```


```javascript
<script>
export default {
  name: "Test.vue",
  methods: {
    changData() {
      this.inputValue = "欢迎鲁班"
    }
  },
  data() {
    return {
      inputValue: "hello",
    }
  }
}
</script>
```

4. v-on 用于绑定事件， 指令参数为对应的事件名，可以简写为@ 
方式一： v-on:click
   <button v-on:click="changData">改变数据</button>
方式二： @click


5. v-show  用于控制元素display css属性， 可以展示或者隐藏元素
通常用于form表单，input等
```
    <h2 v-show="isShow">姓名</h2>
    <button @click="showBtn(true)">显示</button>
    <button @click="showBtn(false)">隐藏</button>
```


```javascript
<script>
export default {
  name: "Test.vue",
  methods: {
    showBtn(value) {
      this.isShow = value
    }
  },
  data() {
    return {
      isShow: true,
    }
  }
}
</script>
```


## Vue 常用声明周期
1。 created() DOM未完成时进行Ajax请求
2. mounted()
3. destroyed()


# 插槽使用slot
子组件
```
<template>
  <div>
    <!--定义插槽，同时指定插槽名称，如果不指定默认为default -->
    <slot name="slot1" :data1="data1">slot1 default value</slot>
  </div>
</template>

<script>
export default {
  name: "One",
  data() {
    return {
      data1: "test......."
    }
  }
}
</script>

<style scoped>

</style>
```


父组件通过插槽获取子组件数据
```
<template>
  <div>
    <One>
      <!--1。 指定子组件插槽名称-->
      <template #slot1="{data1}">
        <h1>父组件slot1 获取子组件作用域中的数据===> {{data1}}</h1>
      </template>
    </One>
  </div>
</template>

<script>
// 2。 导入子组件
import One from "./B.vue"
export default {
  name: "A",
  // 3。 声明组件
  components: {
    One,
  }
}
</script>

<style scoped>

```

该指令可以缩写成 #
如： v-slot:k8s,    v-slot,  #





# vue 路由使用
1。 安装vue-router
npm install --save vue-router@3
2。 创建路由文件 src/router/index.js
```javascript
需要在index.js 中导入并注册路由

import Vue from 'vue'
// 1。导入路由
import VueRouter from 'vue-router'

// 2。 注册路由
Vue.use(VueRouter);

// 3。 配置routes
const routes = [

   {
      path: '/',
      name: 'Home',
      component: () => import('@/view/Layout.vue'),
      children: [
         {
            path: '',
            name: 'Index',
            component: () => import('@/view/Index')
         },
         {
            path: '/k8s/cluster',
            name: 'ClusterManage',
            component: () => import('@/view/k8s/ClusterManage.vue')
         },
         {
            path: '/k8s/nodes',
            name: 'Nodes',
            component: () => import('@/view/k8s/Nodes.vue')
         },
         {
            path: 'test',
            name: 'Test',
            component: () => import('@/view/slot/A.vue')
         }
      ]
   },
   {
      path: '/user/login',
      name: 'Login',
      component: () => import('@/view/user/Login.vue'),
      meta: {
         title: '用户登录',
         module: "用户登录"
      }
   },
   {
      path: '/user/register',
      name: 'Register',
      component: () => import('@/view/user/Register.vue'),
      meta: {
         title: '用户注册'
      }
   }
];

// 4。 创建一个路由 router 实例，通过 routes 属性来定义路由匹配规则
const router = new VueRouter({
   routes,
   mode: 'history',
   base: process.env.BASE_URL,
});

// 导航守卫
// router.beforeEach((to, from, next) => {
//     // 从localStorage获取用户登陆标识
//     if (!localStorage.getItem("onLine")) {
//         if (to.path !== '/user/login') {
//             return next('/user/login')
//         }
//     }
//
//     next()
// });

// 5。 导出router
export default router


3。 挂载路由， App.vue中配置
<template>
<div id="app">
        <!--5. isRouterAlive刷新页面-->
<router-view v-if="isRouterAlive"></router-view>
        </div>
</template>


```




## Vue 路由传参
1。 params
```
{
   path: '/user/detail/:id',  // :id占位符
   name: 'UserDetail',
   component: () => import('@/view/user/UserDetail.vue'),
},

   // 1。 @click 执行事件方法
    <button @click="goUserDetail">跳转到用户详情</button>
    // 2。 使用router-link to
    <router-link to="{ name: 'UserDetail', params: {id: 100}}"></router-link>
```



2。 query
不需要提前声明占位符, 跳转时使用query， 此时url为/user/detail?id=xxxx
```javascript
      this.$router.push(
          {
            name: "UserDetail",
            query: {id :100}
          }
      )
```

## 总结
#### 获取参数
this.$route.query  // 获取query参数
this.$route.params // 获取params参数


route与router区别
route 可以获取路由信息， 如： path， params， query， name等
router 可以对路由进行跳转，可以通过push， replace进行路由跳转



#### 路由跳转
// 1。 @click 执行事件方法
<button @click="goUserDetail">跳转到用户详情</button>
// 2。 使用router-link to
<router-link to="{ name: 'UserDetail', params: {id: 100}}"></router-link>

router-link可以根据路由名称（name）， 或者path方式进行跳转。 跳转时可以通过
params， query进行参数的传递



#### 路由守卫
```javascript
// 导航守卫
router.beforeEach((to, from, next) => {
    // 从localStorage获取用户登陆标识
    if (!localStorage.getItem("onLine")) {
        
        if (to.path !== '/user/login') {
            return next('/user/login')
        }
    }

    next()
});
```



#### axios封装
```javascript
import axios from 'axios'
import {Message} from 'element-ui'
import store from '../store'
// 默认请求头信息
axios.defaults.headers.post['Content-Type'] = 'application/json;charset=UTF-8';
axios.defaults.headers.put['Content-Type'] = 'application/json;charset=UTF-8';

// 创建axios实例
const instance = axios.create({
    baseURL: process.env.BASE_URL,
    timeout: 3000,
})


// request 拦截器
instance.interceptors.request.use(config => {
    if (store.getters.token) {
        config.headers['token'] = 'jwt ' + store.getters.token
    }
    return config
    },
    error => {
        Promise.reject(error)
})

// 响应拦截器
instance.interceptors.response.use(
    response => {
        if (response.status === 200) {
            return Promise.resolve(response)
        } else {
            Message({
                message: response.msg,
                type: 'error',
                duration: 5 * 1000,
            })
            return Promise.reject(response)
        }
    },
    error => {
        const {response} = error
        if (response) {
            Message({
                message: response.msg,
                type: 'error',
                duration: 5 * 1000
            })
            return Promise.reject(response)
        }else {
            Message({
                message: "系统异常",
                type: 'error',
                duration: 5 * 1000
            })
            return Promise.reject(error)
        }
    }
)


export default function (method, url, data = null) {
    method = method.toLowerCase()
    if (method === 'post') {
        return instance.post(url, data)
    } else if (method === 'get') {
        return instance.get(url, { params: data })
    } else if (method === 'delete') {
        return instance.delete(url, { params: data })
    } else if (method === 'put') {
        return instance.put(url, data)
    } else {
        console.error("未知的方法:" + method)
        return false
    }
}
```


#### 全局注册
```javascript
在main.js中新增
import axios from "axios";

Vue.prototype.axios = axios;
```