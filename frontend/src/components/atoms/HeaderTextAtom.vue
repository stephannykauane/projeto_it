<template>
    <div class="container-text-cabecalho">
         <h2>{{ text }}</h2>
    </div>
</template>

<script setup>
defineProps({
    text: String
})
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Arya:wght@400;700&display=swap');

.container-text-cabecalho {
    display: flex;
    justify-content: center;
    font-size: 40px;
    font-family: 'Arya', serif;
    text-transform: uppercase;
    white-space: nowrap; 
    margin: 0;
}


</style>
