import HomePage from "../components/pages/HomePage.vue";
import { createWebHistory, createRouter } from 'vue-router'
import LoginPage from "../components/pages/LoginPage.vue";
import SignupPage from "../components/pages/SignupPage.vue";
import ProfilePage from "../components/pages/ProfilePage.vue";

const routes = [{
    path: '/',
    redirect: 'home',
    children: [
        { path: 'home', name: 'home', component: HomePage},
        { path: 'login', name: 'login', component: LoginPage},
        { path: 'signup', name: 'signup', component: SignupPage},
        { path: 'profile', name: 'profile', component: ProfilePage},
    ]
}]

const router = createRouter({
    history:createWebHistory(),
    routes
})

export { router }

