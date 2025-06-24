<template>
  <div class="container-form-calculator">

    <div class="register-stepper">
      <div class="line-stepper"></div>
      <div class="stepperWrappers">
        <div v-for="n in 4" :key="n" class="step-wrapper">
          <div class="step" :class="{ 'step-active': step === n, 'step-done': step > n }">
            <span class="step-number">{{ n }}</span>
          </div>
          <span class="step-description" :class="{ 'step-description-active': step === n }">
            {{ stepDescriptions[n - 1] }}
          </span>
        </div>
      </div>
    </div>


    <div class="form-transitions-container">
      <transition name="slide-fade">
        <section v-if="step === 1">
          <form class="calculator-stepper">
            <div class="areaInformations">
              <TextInputMolecule text="Consultor:" />
              <TextInputMolecule text="Propriedade:" />
              <TextInputMolecule text="Área/talhão:" />
            </div>
          </form>
        </section>
      </transition>

      <transition name="slide-fade">
        <section v-if="step === 2">
          <form class="calculator-stepper">
            <div class="metodoChoice">
              <div class="metodos">
                <div v-for="(label, method) in metodoLabels" :key="method" class="metodo">
                  <button type="button" class="buttonChange" @click="changeMethod(method)">
                    {{ label }}
                  </button>
                </div>
              </div>
            </div>
          </form>
        </section>
      </transition>

      <transition name="slide-fade">
        <section v-if="step === 3" class="teste">
          <form class="calculator-stepper">
            <div class="analiseInformations">
              <div class="child" v-if="selectedMethod === 'saturacaoBases'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmolc/dm³):" />
                  <TextInputMolecule text="Saturação desejada (%V):" />
                </div>
                <div class="children">
                  <TextInputMolecule text="Saturação Atual (V%):" />
                  <TextInputMolecule text="PRNT do corretivo (%):" />
                </div>
              </div>

              <div class="child" v-else-if="selectedMethod === 'aluminioTrocavel'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmolc/dm³):" />
                  <TextInputMolecule text="Argila:" />
                  <TextInputMolecule text="Cálcio:" />
                </div>
                <div class="children">
                  <TextInputMolecule text="Magnésio:" />
                  <TextInputMolecule text="Alumínio:" />
                  <TextInputMolecule text="PRNT do corretivo (%):" />
                </div>
              </div>

              <div class="child" v-else-if="selectedMethod === 'saturacaoCalcio'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmolc/dm³):" />
                  <TextInputMolecule text="Teor de Ca da análise (cmolc/dm³): " />
                </div>
                <div class="children">
                  <TextInputMolecule text="PRNT do corretivo (%):" />
                  <TextInputMolecule text="CaO do corretivo (%):" />
                  <TextInputMolecule text="Porcentagem de Ca desejada (%):" />
                </div>
              </div>

              <div class="child" v-else-if="selectedMethod === 'saturacaoMagnesio'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmolc/dm³):" />
                  <TextInputMolecule text="Teor de Mg da análise (cmolc/dm³): " />
                </div>
                <div class="children">
                  <TextInputMolecule text="PRNT do corretivo (%):" />
                  <TextInputMolecule text="CaO do corretivo (%):" />
                  <TextInputMolecule text="Porcentagem de Mg desejada (%):" />
                </div>
              </div>
            </div>
          </form>
        </section>
      </transition>

      <transition name="slide-fade">
        <section v-if="step === 4">
          <form class="calculator-stepper">
            <div class="resultInformations">
              <h2>Dados da Análise</h2>
              <div class="analiseData" v-if="selectedMethod === 'saturacaoBases'">
                <TextDisplayMolecule label="CTC (cmolc/dm³):" :value="10" />
                <TextDisplayMolecule label="Saturação desejada (%V):" :value="20" />
                <TextDisplayMolecule label="Saturação Atual (V%):" :value="1" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="120" />
              </div>

              <div class="analiseData" v-if="selectedMethod === 'aluminioTrocavel'">
                <TextDisplayMolecule label="CTC (cmolc/dm³):" :value="10" />
                <TextDisplayMolecule label="Argila: ":value="20" />
                <TextDisplayMolecule label="Cálcio:" :value="1" />
                <TextDisplayMolecule label="Magnésio:" :value="120" />
                <TextDisplayMolecule label="Alumínio:" :value="120" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="120" />
              </div>

              <div class="analiseData" v-if="selectedMethod === 'saturacaoCalcio'">
                <TextDisplayMolecule label="CTC (cmolc/dm³):" :value="10" />
                <TextDisplayMolecule label="Teor de Ca da análise (cmolc/dm³): " :value="20" />
                <TextDisplayMolecule label="PRNT do corretivo (%):"  :value="1" />
                <TextDisplayMolecule label="CaO do corretivo (%):"  :value="120" />
                <TextDisplayMolecule label="Porcentagem de Ca desejada (%):" :value="120" />
              </div>

              <div class="analiseData" v-if="selectedMethod === 'saturacaoMagnesio'">
                <TextDisplayMolecule label="CTC (cmolc/dm³):" :value="10" />
                <TextDisplayMolecule label="Teor de Mg da análise (cmolc/dm³): " :value="20" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="120" />
                <TextDisplayMolecule label="MgO do corretivo (%):" :value="1" />
                <TextDisplayMolecule label="Porcentagem de Mg desejada (%):"  :value="120" />
              </div>

              <div class="resultadoCalagem">
                <h2>Resultado</h2>
                <TextDisplayMolecule label="Necessidade de Calagem (ton/ha):" :value="30" />
              </div>
            </div>
          </form>
        </section>
      </transition>
    </div>


    <div class="buttons-next-and-back">
      <div v-if="step === 1">
        <div class="next-and-back">
          <div class="cta"></div>
          <div class="cta">
            <ButtonAlterarAtom text="Avançar >" @click="next" />
          </div>
        </div>
      </div>

      <div v-else-if="step === 2 || step === 3">
        <div class="next-and-back">
          <div class="cta">
            <ButtonAlterarAtom text="< Voltar" @click="back" />
          </div>
          <div class="cta" v-if="step === 3">
            <ButtonAlterarAtom text="Avançar >" @click="next" />
          </div>
        </div>
      </div>

      <div v-else-if="step === 4">
        <div class="next-and-back">
          <div class="cta">
            <ButtonAlterarAtom class="back" text="< Voltar" @click="back" />
          </div>
          <div class="cta">
            <ButtonAlterarAtom text="Reiniciar" @click="reiniciar" />
          </div>
          <div class="cta-especial">
            <ButtonExportarAtom text="Exportar para Excel" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref } from 'vue'


import TextInputMolecule from '../molecules/TextInputMolecule.vue'
import TextDisplayMolecule from '../molecules/TextDisplayMolecule.vue'
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue'
import ButtonExportarAtom from '../atoms/ButtonExportarAtom.vue'

const step = ref(1)


const stepDescriptions = ['Informações da área', 'Escolha do método', 'Dados da análise', 'Resultado']


const metodoLabels = {
  saturacaoBases: 'Saturação por Bases',
  aluminioTrocavel: 'Alumínio Trocável',
  saturacaoCalcio: 'Saturação de Cálcio',
  saturacaoMagnesio: 'Saturação de Magnésio'
}

const selectedMethod = ref<'saturacaoBases' | 'aluminioTrocavel' | 'saturacaoCalcio' | 'saturacaoMagnesio'>('saturacaoBases')





function next() {
  if (step.value < 4) step.value++
}

function back() {
  if (step.value > 1) step.value--
}

function reiniciar() {
  step.value = 1

}

function changeMethod(method: keyof typeof metodoLabels) {
  selectedMethod.value = method
  next()
}
</script>


<style scoped>
.resultInformations {
  font-family: 'Konkhmer Sleokchher';
  display: flex;
  justify-content: center;
  flex-direction: column;
  width: 100%;
}


.analiseData {
  display: flex;
  width: 100%;
  justify-content: center;
  flex-direction: column;
}

.form-transitions-container {
  width: 100%;
  margin-bottom: 1em;
  min-height: 25em;
}


.calculator-stepper {
  display: flex;
  justify-content: center;
  flex-direction: column;
  height: 100%;
  min-height: 20em;
  width: 100%;
}



.buttons-next-and-back {
  margin-top: auto;
}

.next-and-back {
  display: flex;
  margin-top: auto;
  justify-content: space-between;
  flex-direction: row;

}

.child {
  display: flex;
  margin-top: 3em;
  justify-content: center;
  flex-direction: column;
}

.children {
  display: flex;
  justify-content: space-around;
  flex-direction: row;
  gap: 0.5em;
}

.buttonChange {
  font-family: 'Konkhmer Sleokchher';
  font-size: 1em;
  width: 100%;
  padding: 1em 2em;
  background-color: rgba(56, 64, 49, 1);
  transition: .3s;
  border: 2px solid;
  border-color: #94af8b;
}

.register-stepper {
  position: relative;
  width: 100%;
  margin-bottom: 1em;
}

.metodos {
  display: flex;
  justify-content: center;
  flex-direction: column;
  margin: 2em;
  background-color: rgba(56, 64, 49, 1);
  border: 2px solid;
  border-color: #94af8b;
  border-radius: 1em;
}

.buttonChange:hover {
  background-color: #94af8b;
  transition: .3s;

}

.metodo {
  margin: 0.5em;
}

.stepperWrappers {
  font-family: 'Konkhmer Sleokchher';
  display: flex;
  justify-content: space-between;
  position: relative;
  z-index: 2;
}

.button-alterar-profile {
  width: 100%;
  justify-content: flex-end;
}

.step-wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  min-height: 70px;
  position: relative;
}



.line-stepper {
  position: absolute;
  top: 12px;
  left: 0;
  right: 0;
  height: 2px;
  background-color: #797c76;
  z-index: 1;
}

.register-stepper .step {
  z-index: 2;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  border: 3px solid #87927a;
  color: #919291;
  background-color: white;
  border-radius: 50%;
  width: 2em;
  height: 2em;
  font-size: 16px;
  transition: .3s;
}

.register-stepper .step-done {
  color: #a7e4b5;
  border-color: #112c17;
  transition: .3s;
}

.register-stepper .step-number {
  font-weight: 800;
  line-height: 1;
  transition: .3s;
}

.step-description {
  margin-top: 0.5em;
  text-align: center;
  color: #cdcfcbd7;
  font-size: 0.86em;
}


.register-stepper .step-active {
  color: #048104;
  background-color: #7fe07b;
  border-color: #4a9e5e;
  font-size: 18px;
  transition: .3s;
}

.register-stepper .step-done {
  color: #7bb388;
  border-color: #78d68e;
}

.register-stepper .step-number {
  font-weight: 800;
  line-height: 1;
  vertical-align: middle;
}

.step-description-active {
  color: #ffffff;
  font-size: 16px;
  transition: .3s;
  font-weight: bold;
  text-shadow: 0 0 5px #7fe07b;
}

.container-form-calculator {
  display: flex;
  justify-content: center;
  flex-direction: column;
  background-color: rgba(56, 64, 49, 1);
  border-radius: 1em;
  padding: 5em 9em;
  min-height: 40em;
  margin-top: 3em;
  margin-bottom: 2em;
}

.button-calcular-exportar {
  width: 27%;
  margin-left: 48%;
  display: flex;
  justify-content: flex-end;
  margin-top: 1em;
  transition: .3s;
}

.texto-informativo {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  text-align: center;
  font-family: 'Konkhmer Sleokchher';
  width: 80%;
}

.areaInformations {
  display: flex;
  justify-content: center;
  width: 100%;
  flex-direction: column;
}

::v-deep(.profile-alterar) {
  padding: 0.8em 2.5em;
}

::v-deep(.profile-alterar:hover) {
  background-color: #738b5b8f;
}





@media screen and (max-width:1025px) {

  .container-form-calculator {
    padding: 5em 4em;
  }

  .step-description,
  .buttons-next-and-back {
    font-size: 0.8em;
  }

}


@media screen and (max-width:769px) {

  .form-transitions-container {
    min-height: 15em;
  }
  .child {
    margin-top: 0;
  }

  .register-stepper .step {
    font-size: 0.8em;
  }

  .areaInformations,
  .metodos,
  .children {
    font-size: 0.8em;
  }

  .step-description {
    font-size: 0.7em;
  }

  .container-form-calculator {
    padding: 5em 3em;
  }

}

@media screen and (max-width:580px) {

  .next-and-back {
    flex-wrap: wrap;
    
  }

  .cta-especial{
    margin-top: 1em;
  }


}


@media screen and (max-width:550px) {


  .register-stepper .step {
    font-size: 0.7em;
  }

  .step-description {
    font-size: 0.6em;
  }

  .container-form-calculator {
    padding: 5em 2em;
  }

}




@media screen and (max-width:480px) {

  .container-method-calculator {
    padding: 5em 4em;
  }

  @media screen and (max-width:769px) {

    .register-stepper .step {
      font-size: 0.8em;
    }

    .step-description {
      font-size: 0.7em;
    }

    .container-form-calculator {
      padding: 5em 3em;
    }

  }


}

@media screen and (max-width:375px) {

.container-method-calculator {
  padding: 5em 2em;
}

@media screen and (max-width:769px) {

  .register-stepper .step {
    font-size: 0.7em;
  }

  .step-description {
    font-size: 0.6em;
  }

  .container-form-calculator {
    padding: 5em 2em;
  }

}


}
</style>