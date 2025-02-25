import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import piniaPluginPersistedState from "pinia-plugin-persistedstate"
import { createPinia } from 'pinia'
import { router } from './scripts/index.js'


const pinia = createPinia()
pinia.use(piniaPluginPersistedState)

createApp(App).use(pinia).use(router).mount('#app')
