<script setup lang="ts">
import DescriptionAtom from '../atoms/DescriptionAtom.vue';

</script>

<template>
    <div class="container-description-molecule">
      <div>
        <DescriptionAtom text="✅ Intuitivo" subtext="Fácil em sua utilização, Caliming promete uma interface intuitiva e descomplicada" />
      </div>
      <div>
        <DescriptionAtom text="🤝 Amigável" subtext="Seu objetivo principal e valor inegociável é a construção de um vínculo forte entre o agro e a tecnologia" />
      </div>
      <div>
        <DescriptionAtom text="⚡ Boa performance" subtext="É rápida e eficiente, evitando atrasos e colaborando na previsão de prazos" />
      </div>
      <div>
        <DescriptionAtom text="✨ Apresenta resultados detalhados" subtext="A plataforma oferece uma análise completa a partir do detalhamento oferecido" />
      </div>
      <div>
        <DescriptionAtom text="🎲 Garantia à proteção de dados" subtext="Caliming possui um ambiente seguro e confiável quanto a armazenamento de dados e informações pessoais" />
      </div>      
    </div>
  </template>
  
  <style>
  
.container-description-molecule{
  width: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
}


.container-description-molecule {
    animation: right 1.5s ease-in-out;
}

@keyframes right {
    from {
        opacity: 0;
        transform: translateX(10px); 
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}

  </style>
  