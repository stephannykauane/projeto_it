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
        { path: 'home', name: 'home', component: HomePage},
        { path: 'login', name: 'login', component: LoginPage},
        { path: 'signup', name: 'signup', component: SignupPage},
        { path: 'profile', name: 'profile', component: ProfilePage},
        { path: 'calculator', name: 'calculator', component: CalculatorPage},
        { path: 'history', name: 'history', component: HistoryPage},
    ]
}]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export { router }

