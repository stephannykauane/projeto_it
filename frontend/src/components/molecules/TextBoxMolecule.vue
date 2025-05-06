<script setup lang="ts">
import TextBoxAtom from '../atoms/TextBoxAtom.vue';

</script>

<template>
    <div class="text-box-molecule">
        <TextBoxAtom text="Escolha o método"/>
        <TextBoxAtom text="Preencha os campos solicitados"/>
        <TextBoxAtom text="Calcule e obtenha os resultados "/>
        <TextBoxAtom text="Acesse o painel profissional para visualizar todos os cálculos"/>
    </div>
</template>

<style scoped>

.text-box-molecule{
    display: flex;
    width: 107.8%;
    justify-content: center;
    flex-direction: row;
    margin-left: -3.4em;
    margin-top: 2em;
    transition: .3s;
    
}

@media screen and (max-width:1500px){
    .text-box-molecule{
        width: 107.6%;
        margin-left: -3em;
    }
}

@media screen and (max-width:1200px){
    .text-box-molecule{
        width: 107.2%;
        margin-left: -2.2em;
    
    }
}

@media screen and (max-width:900px){
    .text-box-molecule{
        width: 104%;
        margin-left: -1em;
    
    }
}

@media screen and (max-width:769px){
    .text-box-molecule{
        width: 102%;
        margin-left: -0.5em;
    
    }
}

@media screen and (max-width:480px){
    .text-box-molecule{
        width: 102%;
        margin-left: -0.3em;
        margin-top: 1em;
    
    }
}
</style>