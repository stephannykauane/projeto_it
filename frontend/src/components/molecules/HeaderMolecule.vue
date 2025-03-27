<template>
  <div class="header-molecule">
    <div class="text-caliming">
        <HeaderTextAtom text="Caliming"></HeaderTextAtom>
    </div>    
    <div class="button-caliming">
      <router-link to="login">
        <ButtonHomeAtom customClass="button-caliming" text="Entrar"></ButtonHomeAtom>
      </router-link>
  
    </div>
  </div>
</template>

<script setup lang="ts">
import ButtonHomeAtom from '../atoms/ButtonHomeAtom.vue';
import HeaderTextAtom from '../atoms/HeaderTextAtom.vue';
</script>

<style>
.header-molecule {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  margin:0rem 1rem;
  padding: 1rem;
  box-sizing: border-box;
}


.text-caliming{
    animation: text 1.5s ease-in-out;
}

@keyframes text {
    from {
        opacity: 0;
        transform: translateX(-10px); 
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}

.button-caliming{
    animation: button 1.5s ease-in-out;
}

@keyframes button {
    from {
        opacity: 0;
        transform: translateX(10px); 
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}
@media screen and (max-width: 600px) {
  .header-molecule{
    margin-top: 0.5rem;
    padding: 0;
  }
}


</style>