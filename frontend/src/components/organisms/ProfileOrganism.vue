<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import AltProfileMolecule from '../molecules/AltProfileMolecule.vue';
import HeaderAndNavMolecule from '../molecules/HeaderAndNavMolecule.vue';

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
    <div class="container-profile-organism">
      <div class="header-principal-organism">
        <HeaderAndNavMolecule/>
      </div>     
      <div class="alt">
        <AltProfileMolecule/>
      </div>
    </div>
</template>

<style scoped>

.header-principal-organism {
    display: flex;
    width: 100%;
    flex-direction: column;

}

.container-profile-organism {
    width: 100%;
}

.alt {
  display: flex;
  justify-content: center;
  align-items: center;
  padding:0 23vw;
  width: 100%;
  transition: 0.5s;
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
