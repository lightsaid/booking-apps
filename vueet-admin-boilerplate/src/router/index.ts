import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

export const routes: RouteRecordRaw[] = [
    {
        path: "/login",
        name: "login",
        component: import("@/views/Login.vue"),
        meta: {
            authRequired: false,
            title: "登录",
            hiddenMenu: true  // 在菜单中隐藏 
        },
    },
    {
        path: "/",
        component: () => import("@/components/Layout/Basic.vue"),
        redirect: "/dashboard",
        meta:{
            title: "Home"
        },
        children: [
            {
                path: "/dashboard",
                name: "dashboard",
                component: import("@/views/Dashboard.vue"),
                meta: {
                    authRequired: true,
                    title: "Dashboard",
                    icon: "Odometer"
                },
            },
        ],
    },
    {
        path: "/sys",
        component: () => import("@/components/Layout/Basic.vue"),
        name: "sys",
        redirect: "/sys/settings",
        meta: {
            authRequired: true,
            title: "系统管理",
            icon: "Monitor"
        },
        children: [
            {
                path: "/sys/settings",
                name: "settings",
                component: import("@/views/Settings.vue"),
                meta: {
                    authRequired: true,
                    title: "用户设置"
                },
            },
            {
                path: "/sys/role",
                name: "role",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "角色管理"
                },
            },
            {
                path: "/sys/user",
                name: "user",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "用户列表"
                },
            },
            {
                path: "/sys/order",
                name: "order",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "订单列表"
                },
            },
        ],
    },
    {
        path: "/theater",
        component: () => import("@/components/Layout/Basic.vue"),
        name: "theater",
        redirect: "/theater/index",
        meta: {
            authRequired: true,
            title: "电影院管理",
            icon:"Camera"
        },
        children: [
            {
                path: "/theater/index",
                name: "theaterIndex",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "影院列表"
                },
            },
            {
                path: "/theater/seat",
                name: "seat",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "座位列表"
                },
            },
        ],
    },
    {
        path: "/movie",
        component: () => import("@/components/Layout/Basic.vue"),
        name: "movie",
        redirect: "/movie/index",
        meta: {
            authRequired: true,
            title: "电影管理",
            icon:"VideoPlay"
        },
        children: [
            {
                path: "/movie/index",
                name: "movieIndex",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "电影列表"
                },
            },
            {
                path: "/movie/showtime",
                name: "showtime",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "场次列表"
                },
            },
            {
                path: "/movie/ticket",
                name: "ticket",
                component: import("@/views/Todo.vue"),
                meta: {
                    authRequired: true,
                    title: "电影票列表"
                },
            },
        ],
    },
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
    scrollBehavior(to, from, savedPosition) {
        return (
            savedPosition ||
            new Promise((resolve) => {
                setTimeout(() => resolve({ top: 0 }), 300);
            })
        );
    },
});

// 路由拦截
// router.beforeEach((to, from) => {
// 	if (to.meta.authRequired) {
// 		// return { name: 'login', query: { redirect: to.fullPath } }
// 	}
// })

export default router;
