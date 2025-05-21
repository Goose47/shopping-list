import HomeView from '@/views/protected/HomeView.vue'
import HTTP403View from "@/views/public/HTTP403View.vue";
import {useAuthStore} from "@/stores/authStore.js";

const getUser = async () => {
    const userStore = useAuthStore()
    if (!userStore.user) {
        await userStore.getUser()
    }
    return userStore.user
}

const checkAuth = async (to, from, next) => {
    const user = await getUser()

    if (!user) {
        next({
            name: '403',
        });
    } else {
        next();
    }
}

const routes = [
    {
        path: '/',
        name: 'home',
        component: HomeView,
        beforeEnter: [checkAuth],
    },
    {
        path: '/403',
        name: '403',
        component: HTTP403View,
    },
]

export default routes