import HomePage from "../components/pages/HomePage.vue";
import { createWebHistory, createRouter } from 'vue-router'
import LoginPage from "../components/pages/LoginPage.vue";
import SignupPage from "../components/pages/SignupPage.vue";
import ProfilePage from "../components/pages/ProfilePage.vue";
import CalculatorPage from "../components/pages/CalculatorPage.vue";
import HistoryPage from "../components/pages/HistoryPage.vue";

const routes = [{
    path: '/',
    redirect: 'home',
    children: [
        { path: 'home', name: 'home', component: HomePage, meta: { title: 'Caliming' } },
        { path: 'login', name: 'login', component: LoginPage, meta: { title: 'Login' } },
        { path: 'signup', name: 'signup', component: SignupPage, meta: { title: 'Cadastro' } },
        { path: 'profile', name: 'profile', component: ProfilePage, meta: { title: 'Perfil', requiresAuth: true } },
        { path: 'calculator', name: 'calculator', component: CalculatorPage, meta: { title: 'Calculadora', requiresAuth: true } },
        { path: 'history', name: 'history', component: HistoryPage, meta: { title: 'Histórico', requiresAuth: true } },
    ]
}]

const router = createRouter({
    history:createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const token = localStorage.getItem('token')
    if (to.meta.requiresAuth && !token) {
        return next({ name: 'login' })
    }

    next();
})

router.afterEach((to) => {
    if (to.meta.title) {
        document.title = to.meta.title;
    }
})

export { router }

