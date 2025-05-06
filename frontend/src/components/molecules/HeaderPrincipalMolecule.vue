<script setup lang="ts">

import calimingAPI from '../../services/CalimingAPIClient';
import ButtonSairAtom from '../atoms/ButtonSairAtom.vue';
import HeaderTextAtom from '../atoms/HeaderTextAtom.vue';
import SeparatorAtom from '../atoms/SeparatorAtom.vue';
import { useRouter } from 'vue-router'

const router = useRouter()

const handleLogout = () => {
    calimingAPI.logout()
    router.push('/home')
}

</script>

<template>
    <div class="container-header-molecule">
        <div class="header-molecule">
            <div class="text-caliming">
             <HeaderTextAtom class="caliming" text="Caliming"/>
            </div>    
            <div class="button-Sair">
              <ButtonSairAtom text="Sair" @click="handleLogout"/> 
            </div>       
        </div>   
    </div>
    <div class="container-separator">
        <SeparatorAtom/>
    </div>
  </template>

<style scoped>


.header-molecule {
    display: flex;
    align-items: center;
    width: 100%;
    padding: 1rem;
    margin: 0 14.2em;
    box-sizing: border-box;
    transition: 0.5s;
}


.container-header-molecule {
    background-color: #93A87E;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    box-sizing: border-box;
}


@media screen and (max-width: 1200px) {
  .header-molecule {
    margin-top: 0.5rem;
    padding: 0;
    gap: 5em;
  }
}

@media screen and (max-width: 769px) {
  .header-molecule {
    margin-top: 0.5rem;
    padding: 0;
    gap: 10em;
  }
}

@media screen and (max-width: 480px) {
  .header-molecule {
    margin-top: 0.5rem;
    padding: 0;
    gap: 2em;
  }
}
</style>


