import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

const router = new Router(
    {
        mode: 'history',
        routes: [
            {
                path: '/',
                redirect: '/mainComponent',
                meta: { title: '第一个展示的页面' }
            },
            {
                path: '/main',
                component: () => import('../components/navigation/Navigation.vue'),
                meta: { title: '侧边栏仪表板' },
                children: [
                    {
                        path: '/mainComponent',
                        component: () => import('../components/main/MainComponent.vue'),
                        meta: { title: '主页面' }
                    },
                    {
                        path: '/viewWeeklyReport',
                        component: () => import('../components/main/MainComponent.vue'),
                        meta: { title: '查看周报' }
                    },
                    {
                        path: '/writeWeeklyReport',
                        component: () => import('../components/main/WriteWeeklyReport.vue'),
                        meta: { title: '写周报' }
                    },
                    {
                        path: '/editWeeklyReport',
                        component: () => import('../components/main/InfoComponent.vue'),
                        meta: { title: '编辑周报' }
                    },
                    {
                        path: '/statisticalExport',
                        component: () => import('../components/main/StatisticalExport.vue'),
                        meta: { title: '统计导出' }
                    },
                    {
                        path: '/templateEditing',
                        component: () => import('../components/main/TemplateEditing.vue'),
                        meta: { title: '模板编辑' }
                    },
                    {
                        path: '/userPassword',
                        component: () => import('../components/setUp/UserPassword.vue'),
                        meta: { title: '用户密码' }
                    },
                    {
                        path: '/statisticalRules',
                        component: () => import('../components/setUp/StatisticalRules.vue'),
                        meta: { title: '统计规则' }
                    },
                    {
                        path: '/userManagement',
                        component: () => import('../components/setUp/UserManagement.vue'),
                        meta: { title: '用户管理' }
                    },
                    {
                        path: '/test',
                        component: () => import('../components/test/TestComponent.vue'),
                        meta: { title: '测试页面' }
                    },
                ]
            },
            {
                path: '/login',
                component: () => import('../components/login/LoginComponent.vue'),
                meta: { title: '登录页' }
            },
            {
                path: '*',
                redirect: '/404'
            }
        ]
    }
);

router.beforeEach((to, from, next) => {
    // 获取session中存储的用户登录数据
    let isLogin = sessionStorage.getItem('login');
    if (isLogin) {
        next();
    } else {
        if (to.path === '/login') {
            next();
        } else {
            next('/login');
        }
    }
});

// 解决重复路由报红问题
const originalPush = Router.prototype.push;
Router.prototype.push = function push (location) {
    return originalPush.call(this, location).catch(err => err)
};

export default router
