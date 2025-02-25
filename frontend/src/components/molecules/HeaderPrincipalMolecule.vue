<script setup lang="ts">
import ButtonGreenAtom from '../atoms/ButtonGreenAtom.vue';
import ButtonSairAtom from '../atoms/ButtonSairAtom.vue';
import HeaderTextAtom from '../atoms/HeaderTextAtom.vue';

</script>

<template>
    <div class="container-header-molecule">
        <div class="header-molecule">
            <div class="text-caliming">
             <HeaderTextAtom class="caliming" text="Caliming"/>
            </div>    
            <div class="button-Sair">
              <ButtonSairAtom text="Sair"/> 
            </div>       
        </div>
    </div>
  </template>

<style scoped>
.button-sair{
    color: beige;
}

.header-molecule{
    padding: 1em 20em;
}
.container-header-molecule{
    background-color: #93A87E;
    width: 100vw;;
}


</style>