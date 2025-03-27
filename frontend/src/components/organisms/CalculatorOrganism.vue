<template>
    <div class="container-calculator-organism">
      <div class="header-and-nav">
        <HeaderAndNavMolecule />
      </div>
  
      <div class="calculator">
        <div class="button-method">
          <AccordionAtom :title="methodTitle">
              <p @click="changeMethod('saturacao')">SATURAÇÃO POR BASES</p>
              <p @click="changeMethod('aluminio')">ALUMÍNIO TROCÁVEL</p>
          </AccordionAtom>
        </div>
  
        <CalculatorMethodMolecule :selectedMethod="selectedMethod" />
  
        <div class="text-box-organism">
          <TextBoxMolecule />
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, computed } from 'vue';
  import HeaderAndNavMolecule from '../molecules/HeaderAndNavMolecule.vue';
  import AccordionAtom from '../atoms/AccordionAtom.vue';
  import CalculatorMethodMolecule from '../molecules/CalculatorMethodMolecule.vue';
  import TextBoxMolecule from '../molecules/TextBoxMolecule.vue';
  
  const selectedMethod = ref('saturacao');
  
  const methodTitle = computed(() => {
    return selectedMethod.value === 'saturacao' ? 'SATURAÇÃO POR BASES' : 'ALUMÍNIO TROCÁVEL';
  });
  
  const changeMethod = (method: string) => {
    selectedMethod.value = method;
  };
  </script>
  
  <style>
  .container-calculator-organism {
    display: flex;
    justify-content: center;
    width: 100%;
    flex-direction: column;
  }
  
  .calculator {
    display: flex;
    justify-content: center;
    flex-direction: column;
    margin: -2em 16em;

  }
  
  .button-method {
  display: flex;
  justify-content: flex-start;
  width: 20%;
  margin: 2em 0;
  position: relative; 
  z-index: 5;
  transition: .3s;
}

.button-method p {
    color: #ffffff;
    cursor: pointer;
    padding: 0.5em;
}

p:hover{
    background-color: #b6d49a7c;

}
  

@media screen and (max-width:1500px){
  .calculator{
    margin: 0em 9em;
  }

  .button-method{
    width: 23%;
  }

}

@media screen and (max-width:1200px){
  .calculator{
    margin: 0em 7em;
  }

  .button-method{
    width: 28%;
  }

}

@media screen and (max-width:900px){
  .calculator{
    margin: 0em 5.5em;
  }

  .button-method{
    width: 37%;
  }

}

@media screen and (max-width:769px){
  .calculator{
    margin: 0em 4.5em;
  }

  .button-method{
    width: 32%;
    margin: 2em 0 0.8em ; 
  }

}

@media screen and (max-width:480px){
  .calculator{
    margin: 0em 1em;
  }

  .button-method{
    width: 43%;
    margin: 2em 0 0.8em ; 
  }

}
  </style>
  