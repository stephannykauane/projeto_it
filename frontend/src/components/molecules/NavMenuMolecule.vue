<template>
    <nav class="container-menu">
        <ul class="menu">
            <li>
                <router-link to="/profile" :class="{ active: route.path === '/profile' }">
                    Sua conta
                </router-link>
            </li>
            <li>
                <router-link to="/calculator" :class="{ active: route.path === '/calculator' }">
                    Calculadora
                </router-link>
            </li>
            <li>
                <router-link to="/history" :class="{ active: route.path === '/history' }">
                    Histórico
                </router-link>
            </li>
        </ul>
    </nav>
</template>

<script setup>
import { useRoute } from 'vue-router';
const route = useRoute();
</script>

<style scoped>
.container-menu {
    width: 100%;
    display: flex;
    justify-content: center;
    border: 2px solid transparent;
    border-radius: 10em;
    background-color: #70885A;
    padding: 0.5em 0;
    transition: .3s;
}

.menu {
    list-style: none;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    flex-wrap: nowrap;
    padding: 0;
    margin: 0;
    transition: .3s;
}

.menu li {
    margin: 0 10px;
    transition: .3s;
}

.menu a {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0.5em 9em;
    border: 2px solid transparent;
    border-radius: 10em;
    text-decoration: none;
    font-family: 'Konkhmer Sleokchher';
    text-transform: uppercase;
    color: rgba(202, 254, 157, 0.8);
    transition: 0.3s;
    white-space: nowrap;
}


.menu a.active {
    color: white;
    font-weight: bold;
}


.menu a:hover {
    color: rgba(255, 255, 255, 0.781);
    transition: 0.3s;
}

@media screen and (max-width: 1500px) {
    .menu {
        flex-direction: row;
    }

    .menu a {
        padding: 0.5em 5em;
        text-align: center;
    }
}

@media screen and (max-width: 1200px) {
    .menu a {
        padding: 0.5em 2em;
        text-align: center;
        font-size: 0.9em;
    }
}

@media screen and (max-width: 769px) {
    .menu a {
        font-size: 0.7em;
    }
}
</style>
