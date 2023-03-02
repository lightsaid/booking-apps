import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
    {
		path: '/login',
		name: 'login',
		component: import('@/views/Login.vue'),
		meta: {
			AuthRequired: false,
		},
	},
]

const router = createRouter({
	history: createWebHashHistory(),
	routes,
	scrollBehavior(to, from, savedPosition) {
		return (
			savedPosition ||
			new Promise((resolve) => {
				setTimeout(() => resolve({ top: 0 }), 300)
			})
		)
	},
})

// 路由拦截
// router.beforeEach((to, from) => {
// 	if (to.meta.AuthRequired) {
// 		// return { name: 'login', query: { redirect: to.fullPath } }
// 	}
// })

export default router