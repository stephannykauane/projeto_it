<template>
  <div class="container-method-calculator">
    <div v-if="selectedMethod === 'saturacao'">
      <TextInputMolecule text="CTC (cmolc/dm³):"/>
      <TextInputMolecule text="Magnésio:"/>
      <TextInputMolecule text="Cálcio:"/>
      <TextInputMolecule text="Potássio:"/>
      <TextInputMolecule text="PRNT:"/>
    </div>

    <div v-else-if="selectedMethod === 'aluminio'">
      <TextInputMolecule text="CTC (cmolc/dm³):"/>
      <TextInputMolecule text="Argila:"/>
      <TextInputMolecule text="Cálcio:"/>
      <TextInputMolecule text="Magnésio:"/>
      <TextInputMolecule text="Alumínio:"/>
      <TextInputMolecule text="PRNT:"/>
    </div>

    <div class="button-calcular-exportar">
      <ButtonAlterarAtom text="CALCULAR"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import TextInputMolecule from '../molecules/TextInputMolecule.vue';
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue';

defineProps<{
  selectedMethod: string;
}>();
</script>

<style scoped>
.container-method-calculator {
  display: flex;
  flex-direction: column;
  background-color: rgba(56, 64, 49, 1);
  border-radius: 1em;
  padding: 5em 9em;
  margin-top: 3em;
  transition: .3s;
}

.button-calcular-exportar {
  width: 27%;
  margin-left: 50%;
  display: flex;
  justify-content: flex-end;
  margin-top: 1em;
  transition: .3s;
}

::v-deep(.profile-alterar){
  padding: 0.8em 2.5em;
}

::v-deep(.profile-alterar:hover){
  background-color: #738b5b8f;
}


@media screen and (max-width:480px) {
  .container-method-calculator{
    padding: 5em 4em;
  }

  
}
</style>
