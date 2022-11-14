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
                component: () => import('@/view/Test.vue')
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
    },
    {
        path: '/user/detail/:id',
        name: 'UserDetail',
        component: () => import('@/view/user/UserDetail.vue'),
    },
    {
        path: '/user/center',
        name: 'UserCenter',
        component: () => import('@/view/user/UserCenter.vue'),
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
