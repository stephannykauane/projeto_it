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
              <TextInputMolecule text="Consultor:" v-model="consultor" />
              <TextInputMolecule text="Propriedade:" v-model="propriedade" />
              <TextInputMolecule text="Área/talhão:" v-model="area" />
            </div>
          </form>
        </section>
      </transition>

      <transition name="slide-fade">
        <section v-if="step === 2">
          <form class="calculator-stepper">
            <div class="metodoChoice">
              <div class="metodos">
                <div v-for="metodo in metodos" :key="metodo.id" class="metodo">
                  <button type="button" class="buttonChange" @click="changeMethod(metodo)">
                    {{ metodo.label }}
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
              <div class="child" v-if="selectedMethod?.key === 'saturacaoBases'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmol dm⁻³):" v-model="ctc" />
                  <TextInputMolecule text="Saturação desejada (V%):" v-model="sat_desejada" />
                </div>
                <div class="children">
                  <TextInputMolecule text="Saturação Atual (V%):" v-model="sat_atual" />
                  <TextInputMolecule text="PRNT do corretivo (%):" v-model="prnt" />
                </div>
              </div>

              <div class="child" v-else-if="selectedMethod?.key === 'saturacaoCalcio'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmol dm⁻³):" v-model="ctc" />
                  <TextInputMolecule text="Teor de Ca da análise (cmol dm⁻³): " v-model="teor_ca" />
                </div>
                <div class="children">
                  <TextInputMolecule text="PRNT do corretivo (%):" v-model="prnt" />
                  <TextInputMolecule text="CaO do corretivo (%):" v-model="caO" />
                  <TextInputMolecule text="Porcentagem de Ca desejada (%):" v-model="ca_desejada" />
                </div>
              </div>

              <div class="child" v-else-if="selectedMethod?.key === 'saturacaoMagnesio'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmol dm⁻³):" v-model="ctc" />
                  <TextInputMolecule text="Teor de Mg da análise (cmol dm⁻³): " v-model="teor_mg" />
                </div>
                <div class="children">
                  <TextInputMolecule text="PRNT do corretivo (%):" v-model="prnt" />
                  <TextInputMolecule text="MgO do corretivo (%):" v-model="mgO" />
                  <TextInputMolecule text="Porcentagem de Mg desejada (%):" v-model="mg_desejada" />
                </div>
              </div>

              <div class="child" v-else-if="selectedMethod?.key === 'aluminioTrocavel'">
                <div class="children">
                  <TextInputMolecule text="CTC (cmol dm⁻³):" v-model="ctc" />
                  <TextInputMolecule text="Cálcio (cmol dm⁻³):" v-model="calcio" />
                </div>
                <div class="children">
                  <TextInputMolecule text="Magnésio (cmol dm⁻³):" v-model="magnesio" />
                  <TextInputMolecule text="Alumínio (cmol dm⁻³):" v-model="aluminio" />
                  <TextInputMolecule text="PRNT do corretivo (%):" v-model="prnt" />
                </div>
                <div class="tooltip-aluminio">
                  <div class="tooltip-icon">
                    <img src="../../../src/assets/info-icon.svg">
                    <div>
                      <span class="tooltiptext">Método sugerido exclusivamente para solos com menos de 15% de
                        argila</span>
                    </div>
                  </div>
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
              <div class="analiseData" v-if="selectedMethod?.key === 'saturacaoBases'">
                <TextDisplayMolecule label="CTC (cmol dm⁻³):" :value="formatarDecimalBR(calculo?.ctc)" />
                <TextDisplayMolecule label="Saturação desejada (V%):" :value="formatarDecimalBR(calculo?.sat_desejada)" />
                <TextDisplayMolecule label="Saturação Atual (V%):" :value="formatarDecimalBR(calculo?.sat_atual)" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="formatarDecimalBR(calculo?.prnt)" />
              </div>

              <div class="analiseData" v-if="selectedMethod?.key === 'saturacaoCalcio'">
                <TextDisplayMolecule label="CTC (cmol dm⁻³):" :value="formatarDecimalBR(calculo?.ctc)" />
                <TextDisplayMolecule label="Teor de Ca da análise (cmol dm⁻³): " :value="formatarDecimalBR(calculo?.teor_ca)" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="formatarDecimalBR(calculo?.prnt)" />
                <TextDisplayMolecule label="CaO do corretivo (%):" :value="formatarDecimalBR(calculo?.caO)" />
                <TextDisplayMolecule label="Porcentagem de Ca desejada (%):" :value="formatarDecimalBR(calculo?.ca_desejada)" />
              </div>

              <div class="analiseData" v-if="selectedMethod?.key === 'saturacaoMagnesio'">
                <TextDisplayMolecule label="CTC (cmol dm⁻³):" :value="formatarDecimalBR(calculo?.ctc)" />
                <TextDisplayMolecule label="Teor de Mg da análise (cmol dm⁻³): " :value="formatarDecimalBR(calculo?.teor_mg)" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="formatarDecimalBR(calculo?.prnt)" />
                <TextDisplayMolecule label="MgO do corretivo (%):" :value="formatarDecimalBR(calculo?.mgO)" />
                <TextDisplayMolecule label="Porcentagem de Mg desejada (%):" :value="formatarDecimalBR(calculo?.mg_desejada)" />
              </div>


              <div class="analiseData" v-if="selectedMethod?.key === 'aluminioTrocavel'">
                <TextDisplayMolecule label="CTC (cmolc/dm³):" :value="formatarDecimalBR(calculo?.ctc)" />
                <TextDisplayMolecule label="Cálcio (cmol dm⁻³):" :value="formatarDecimalBR(calculo?.calcio)" />
                <TextDisplayMolecule label="Magnésio (cmol dm⁻³):" :value="formatarDecimalBR(calculo?.magnesio)" />
                <TextDisplayMolecule label="Alumínio (cmol dm⁻³):" :value="formatarDecimalBR(calculo?.aluminio)" />
                <TextDisplayMolecule label="PRNT do corretivo (%):" :value="formatarDecimalBR(calculo?.prnt)" />
              </div>

              <div class="resultadoCalagem">
                <h2>Resultado</h2>
                <TextDisplayMolecule label="Necessidade de Calagem (ton ha⁻¹):"
                  :value="formatarDecimalBR(calculo?.resultado)" />
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

      <div v-else-if="step === 2">
        <div class="next-and-back">
          <div class="cta-step2">
            <div>
              <ButtonAlterarAtom text="< Voltar" @click="back" />
            </div>
            <div class="tooltip-icon">
              <img src="../../../src/assets/info-icon.svg">
              <div>
                <span class="tooltiptext">Os cálculos sugeridos têm por base informações publicadas por Sousa e Lobato
                  (2004), no livro "Cerrado: Correção do Solo e Adubação"</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="step === 3">
        <div class="next-and-back">
          <div class="cta-step2">
            <div>
              <ButtonAlterarAtom text="< Voltar" @click="back" />
            </div>
          </div>
          <div class="cta">
            <ButtonAlterarAtom text="Avançar >" @click="sendAnaliseAndNext" />
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
            <ButtonExportarAtom text="Exportar para Excel" @click="exportarParaExcel()" />
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
import dataService from '../../services/dataService'

const step = ref(1)
const consultor = ref('')
const propriedade = ref('')
const area = ref('')
const ctc = ref('')
const prnt = ref('')
const sat_atual = ref('')
const sat_desejada = ref('')
const calcio = ref('')
const magnesio = ref('')
const aluminio = ref('')
const caO = ref('')
const ca_desejada = ref('')
const mgO = ref('')
const mg_desejada = ref('')
const teor_mg = ref('')
const teor_ca = ref('')
const error = ref('')
const stepDescriptions = ['Informações da área', 'Escolha do método', 'Dados da análise', 'Resultado']
const calculo = ref<any>(null)

const metodos = [
  { id: 1, key: 'saturacaoBases', label: 'Saturação por Bases' },
  { id: 2, key: 'saturacaoMagnesio', label: 'Saturação de Magnésio na CTC' },
  { id: 3, key: 'saturacaoCalcio', label: 'Saturação de Cálcio na CTC' },
  { id: 4, key: 'aluminioTrocavel', label: 'Alumínio Trocável' }
]

const selectedMethod = ref<{
  id: number
  key: 'saturacaoBases' | 'saturacaoCalcio' | 'saturacaoMagnesio' | 'aluminioTrocavel'
  label: string
} | null>(null)


function next() {
  if (step.value === 1) {
    if (!consultor.value || !propriedade.value || !area.value) {
      alert("Por favor, preencha todos os campos obrigatórios antes de avançar.");
      return;
    }
  }

  step.value++
}

function formatarDecimalBR(valor: number | string): string {
  if (valor === null || valor === undefined || isNaN(Number(valor))) return '-'
  return valor.toString().replace('.', ',')
}

function parseValorNumerico(valor: string): number {
  return parseFloat(valor.replace(',', '.'))
}


function back() {
  if (step.value > 1) step.value--
}

function reiniciar() {
  step.value = 1
  area.value = ""
  propriedade.value = ""
  consultor.value = ""
  ctc.value = ""
  prnt.value = ""
  sat_atual.value = ""
  sat_desejada.value = ""
  calcio.value = ""
  magnesio.value = ""
  aluminio.value = ""
  caO.value = ""
  ca_desejada.value = ""
  mgO.value = ""
  mg_desejada.value = ""
  teor_mg.value = ""
  teor_ca.value = ""

}

function changeMethod(method: typeof metodos[number]) {
  selectedMethod.value = method
  step.value++
}

async function sendAnaliseAndNext() {

  if (!selectedMethod.value) {
    alert("Método não selecionado.");
    return;
  }

  const camposComuns = [ctc.value, prnt.value]
  const isEmpty = (val: string | number) => val === '' || val === null || val === undefined

  const metodo = selectedMethod.value.key
  let camposObrigatorios: (string | number)[] = []

  if (metodo === 'saturacaoBases') {
    camposObrigatorios = [...camposComuns, sat_desejada.value, sat_atual.value]
  } else if (metodo === 'aluminioTrocavel') {
    camposObrigatorios = [...camposComuns, calcio.value, magnesio.value, aluminio.value]
  } else if (metodo === 'saturacaoCalcio') {
    camposObrigatorios = [...camposComuns, teor_ca.value, caO.value, ca_desejada.value]
  } else if (metodo === 'saturacaoMagnesio') {
    camposObrigatorios = [...camposComuns, teor_mg.value, mgO.value, mg_desejada.value]
  }

  const algumVazio = camposObrigatorios.some(isEmpty)

  if (algumVazio) {
    alert("Preencha todos os campos obrigatórios antes de continuar.");
    return;
  }

  await sendAnalise()
  step.value++
}

const sendAnalise = async () => {
  try {
    const resultado = await dataService.gerarCalculo({
      consultor: consultor.value,
      propriedade: propriedade.value,
      area: area.value,
      ctc: parseValorNumerico(ctc.value),
      prnt: parseInt(prnt.value),
      sat_desejada: parseValorNumerico(sat_desejada.value),
      sat_atual: parseValorNumerico(sat_atual.value),
      calcio: parseValorNumerico(calcio.value),
      aluminio: parseValorNumerico(aluminio.value),
      magnesio: parseValorNumerico(magnesio.value),
      caO: parseValorNumerico(caO.value),
      ca_desejada: parseValorNumerico(ca_desejada.value),
      mgO: parseValorNumerico(mgO.value),
      mg_desejada: parseValorNumerico(mg_desejada.value),
      teor_ca: parseValorNumerico(teor_ca.value),
      teor_mg: parseValorNumerico(teor_mg.value),
      id_metodo: selectedMethod.value?.id ?? 0

    })

    calculo.value = resultado
  } catch (err) {
    error.value = err.message
  }

}

const exportarParaExcel = async () => {
  try {
    if (!calculo.value) {
      error.value = 'Cálculo não encontrado.'
      return
    }

    const payload = {
      consultor: consultor.value,
      propriedade: propriedade.value,
      area: area.value,
      ctc: parseValorNumerico(ctc.value) || 0,
      prnt: parseInt(prnt.value) || 0,
      sat_desejada: parseValorNumerico(sat_desejada.value) || 0,
      sat_atual: parseValorNumerico(sat_atual.value) || 0,
      calcio: parseValorNumerico(calcio.value) || 0,
      aluminio: parseValorNumerico(aluminio.value) || 0,
      magnesio: parseValorNumerico(magnesio.value) || 0,
      caO: parseValorNumerico(caO.value) || 0,
      mgO: parseValorNumerico(mgO.value) || 0,
      teor_ca: parseValorNumerico(teor_ca.value) || 0,
      teor_mg: parseValorNumerico(teor_mg.value) || 0,
      ca_desejada: parseValorNumerico(ca_desejada.value) || 0,
      mg_desejada: parseValorNumerico(mg_desejada.value) || 0,
      id_metodo: selectedMethod.value?.id ?? 0,
      resultado: calculo.value.resultado ?? 0
    }

    await dataService.gerarExcel(payload)

  } catch (err: any) {
    console.error("Erro ao exportar Excel:", err)
    error.value = err.message || "Erro inesperado ao gerar Excel."
  }
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

.cta-step2 {
  display: flex;
  justify-content: space-between;
  width: 100%;

}

.tooltip-aluminio {
  display: flex;
  justify-content: flex-end;
  margin-top: 1em;
}


.tooltip-icon {
  position: relative;
  display: inline-block;
  align-self: flex-end;
  margin-right: 2.3em;
  width: 0;
}


.tooltip-icon .tooltiptext {
  font-family: 'Konkhmer Sleokchher';
  font-size: 0.8em;
  visibility: hidden;
  width: 200px;
  background-color: rgba(133, 175, 123, 0.699);
  color: #fff;
  text-align: center;
  border-radius: 6px;
  padding: 1em 0.6em;
  position: absolute;
  z-index: 1;
  bottom: 150%;
  left: 50%;
  margin-left: -6.7em;
}

.tooltip-icon .tooltiptext::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: rgba(133, 175, 123, 0.699) transparent transparent transparent;
}

.tooltip-icon:hover .tooltiptext {
  visibility: visible;
}


.tooltip-icon img {
  width: 1.8em;
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

  .tooltip-icon .tooltiptext{
    margin-left: -8.7em ;
  }

  .tooltip-aluminio .tooltiptext {
    margin-left: -7em;

  }

  .tooltip-icon {
    width: 0;
  }

  .tooltip-aluminio .tooltip-icon img {
    width: 1.7em;

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
  
  .tooltip-aluminio {
    font-size: 0.8em;
  }

  .tooltip-aluminio .tooltip-icon img {
    width: 1.3em;

  }


  .tooltip-aluminio .tooltiptext {
    margin-left: -6.5em;
    width: 150px;

  }

}

@media screen and (max-width:580px) {

  .next-and-back {
    flex-wrap: wrap;
    gap: 1em;

  }

  .cta-especial {
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



@media screen and (max-width:480px) {

  .container-method-calculator {
    padding: 5em 4em;
  }

  .tooltip-icon .tooltiptext {
    width: 150px;
    margin-left: -6.2em;
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