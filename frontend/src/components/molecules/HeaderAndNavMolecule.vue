<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import HeaderPrincipalMolecule from '../molecules/HeaderPrincipalMolecule.vue';
import NavMenuMolecule from '../molecules/NavMenuMolecule.vue';


onMounted(() => {
  const appDiv = document.getElementById('app');
  if (appDiv) {
    appDiv.style.setProperty('padding', '0', 'important');
    appDiv.style.setProperty('justify-content', 'flex-start', 'important');
    appDiv.style.setProperty('margin', '0', 'important');
    appDiv.style.setProperty('max-width', '100vw', 'important');

    
  }
});


onUnmounted(() => {
  const appDiv = document.getElementById('app');
  if (appDiv) {
    appDiv.style.removeProperty('padding');
    appDiv.style.removeProperty('justify-content');
    appDiv.style.removeProperty('margin');
    appDiv.style.removeProperty('max-width');
  }
});
</script>


<template>
    <div class="container-header-and-nav">
        <div class="header">
            <HeaderPrincipalMolecule/>
        </div>
        <div class="nav">
            <NavMenuMolecule/>
        </div>
    </div>
</template>


<style>

.nav {
    width: 100%;
    padding: 4vw 13vw;
    transition: 0.5em;
    margin-top: 2em;
}

@media screen and (max-width: 1500px){
    .nav{
      padding: 0 9vw;
    }
}


@media screen and (max-width: 480px){
    .nav{
      padding: 0 2vw;
    }
}

</style>