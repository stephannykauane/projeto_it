<template>
    <nav class="container-menu">
        <ul class="menu">
            <a href="#" :style="{ color: route.path === '/profile' ? 'white' : 'white' }">
                    Sua conta
                </a>
            <li><a href="#">Calculadora</a></li>
            <li><a href="#">Painel Profissional</a></li> 
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
    transition: .3s;
    white-space: nowrap;
}


.menu a:hover {
    color: rgba(255, 255, 255, 0.781);
    transition: .3s;
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
