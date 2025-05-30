<template>
  <div class="container-method-calculator">
    <div v-if="selectedMethod === 'saturacao'">
      <TextInputMolecule text="CTC (cmolc/dm³):" v-model="ctc"/>
      <TextInputMolecule text="Saturação desejada:" v-model="sat_desejada"/>
      <TextInputMolecule text="Magnésio:" v-model="magnesio"/>
      <TextInputMolecule text="Cálcio:" v-model="calcio"/>
      <TextInputMolecule text="Potássio:" v-model="potassio"/>
      <TextInputMolecule text="PRNT:" v-model="prnt"/>
    </div>

    <div v-else-if="selectedMethod === 'aluminio'">
      <TextInputMolecule text="CTC (cmolc/dm³):" v-model="ctc"/>
      <TextInputMolecule text="Argila:" v-model="argila"/>
      <TextInputMolecule text="Cálcio:" v-model="calcio"/>
      <TextInputMolecule text="Magnésio:" v-model="magnesio"/>
      <TextInputMolecule text="Alumínio:" v-model="aluminio"/>
      <TextInputMolecule text="PRNT:" v-model="prnt"/>
    </div>

    <div class="resultado">
      <div class="first">
        <h2>Resultado(NC):</h2>
      </div>
      <div class="result">
        <div v-if="resultado" class="aparecer-result">
          {{ resultado }} ton/ha 
        </div>
      </div>
    </div>

    <div class="button-calcular-exportar">
      <ButtonAlterarAtom text="CALCULAR" @click="calculando"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import TextInputMolecule from '../molecules/TextInputMolecule.vue';
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue';
import calimingAPI from '../../services/CalimingAPIClient';
import { ref, watch } from 'vue';

const props = defineProps<{
  selectedMethod: string;
}>();

const ctc = ref('');
const magnesio = ref('');
const calcio = ref('');
const argila = ref('');
const prnt = ref('');
const potassio = ref('');
const aluminio = ref('');
const sat_desejada = ref('');
const resultado = ref('');


const resetarCampos = () => {
  ctc.value = '';
  magnesio.value = '';
  calcio.value = '';
  argila.value = '';
  prnt.value = '';
  potassio.value = '';
  aluminio.value = '';
  sat_desejada.value = '';
  resultado.value = '';

};



watch(() => props.selectedMethod, () => {
  resetarCampos();
});

const calculando = async () => {
  const payload: Record<string, number | string> = {};

  if (props.selectedMethod === 'saturacao') {
    payload.id_metodo = 1;
    payload.ctc = parseFloat(ctc.value);
    payload.magnesio = parseFloat(magnesio.value);
    payload.calcio = parseFloat(calcio.value);
    payload.potassio = parseFloat(potassio.value);
    payload.prnt = parseFloat(prnt.value);
    payload.sat_desejada = parseFloat(sat_desejada.value);
  } else if (props.selectedMethod === 'aluminio') {
    payload.id_metodo = 2;
    payload.ctc = parseFloat(ctc.value);
    payload.argila = parseFloat(argila.value);
    payload.calcio = parseFloat(calcio.value);
    payload.magnesio = parseFloat(magnesio.value);
    payload.aluminio = parseFloat(aluminio.value);
    payload.prnt = parseFloat(prnt.value);
  }

  console.log("aluminio", aluminio.value)

  try {
    const response = await calimingAPI.gerarCalculo(payload);
    resultado.value = `${response}`;
  } catch (err) {
    console.error(err);
    resultado.value = 'Erro no cálculo. Tente novamente.';
  }
};


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
  margin-left: 48%;
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

.first {
  width: 20%;
  display: flex;
  justify-content: flex-end;
  text-align: end;
}
.result {
  display: flex;
  justify-content: center;
  background-color: #738b5b8f;
  width: 80%; 
  border-radius: 1em;
  padding: 0.7em;
  min-height: 3em;
  
}
.resultado{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  font-size: 1.2em;
  font-family: 'Konkhmer Sleokchher';
  border-radius: 1em;
  gap: 2em;
}

.aparecer-result{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  font-size: 1em;
  transition: .3s;
}
@media screen and (max-width:480px) {

  .container-method-calculator{
    padding: 5em 4em;
  }

}
</style>
