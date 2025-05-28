import './css/main.scss';

import { createRouter, createWebHistory } from 'vue-router'
import { createApp } from 'vue'
import { setupTelegram } from '@/bootstrap/telegram'
import { setupAxios } from '@/bootstrap/axios'
import { createPinia } from 'pinia'

import routes from './routes'
import App from './App.vue'
import { setColorTheme } from '@/theme/colors.js'
import defaultComponents from '@/components'

setupTelegram()

const app = createApp(App)

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
    scrollBehavior() {
        return { top: 0 }
    },
})

app.use(router)

const pinia = createPinia()

app.use(pinia)

for (let componentName in defaultComponents) {
    app.component(componentName, defaultComponents[componentName])
}

app.mount('#app')

setupAxios()
setColorTheme('telegram')
