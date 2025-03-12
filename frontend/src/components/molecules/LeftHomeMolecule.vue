<script setup lang="ts">
import TextAtom from '../atoms/TextAtom.vue';
import ImgHomeAtom from '../atoms/ImgHomeAtom.vue';
import ButtonHomeAtom from '../atoms/ButtonHomeAtom.vue';

</script>

<template>
   <div class="container-left">
      <div class="imagem">
        <ImgHomeAtom/>
      </div>
      <div class="texto">
         <div class="nome-caliming">
            <TextAtom text="CALIMING"></TextAtom>
         </div>
         <div class="frase-ferramenta">
            <TextAtom text="A ferramenta perfeita para o cálculo de calagem"></TextAtom>
         </div>
         <div class="frase-plataforma">
            <TextAtom text="Caliming é uma plataforma desenvolvida para facilitar a rotina do agronegócio no cálculo de calagem"></TextAtom>
         </div>
         <div class="botao-comecar">
            <ButtonHomeAtom text="Começar"/>
         </div>
      </div> 
   </div>
</template>

<style scoped>

.container-left {
   display: flex;
   width: 100%;
   height: auto;
   flex-direction: column;
   align-items: center;  
   justify-content: flex-start; 
   box-sizing: border-box;
   text-align: center;
   margin-top: -0.5em;
  
}


.imagem {
   display: flex;
   justify-content: center;
   width:900px;
}



.texto {
   display: flex;
   flex-direction: column;
   align-items: center; 
   text-align: center;
   padding: 1 0em;

}

.nome-caliming {
   font-family: 'Konkhmer Sleokchher', sans-serif;
   font-size: 2em;
   margin-top: 1em;

}

.frase-ferramenta {
   font-family:'Konkhmer Sleokchher', sans-serif;
   font-size: 1.2em;
   color:#CAFE9D;
   margin-top: 0.7em;
}

.frase-plataforma {
   font-family:'Konkhmer Sleokchher', sans-serif;
   font-size: 1em;
   color: #ffffff99;
   margin: 0.5em 1em;
   margin-top: 1.5em;
}

.botao-comecar{
   margin-top: 1em;
}

.container-left {
    animation: left 1.5s ease-in-out;
}

@keyframes left {
    from {
        opacity: 0;
        transform: translateX(-10px); 
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}

.container-right {
    animation: right 1.5s ease-in-out;
}

@keyframes right {
    from {
        opacity: 0;
        transform: translateX(-10px); 
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}

@media (max-width: 767px) {
   .container-left {
      padding: 20px;
   }

   .frase-ferramenta{
      font-size: 1em;
   }
   .frase-plataforma {
      font-size: 0.8em;
   }
}
</style>

