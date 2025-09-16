<script setup lang="ts">
import CheckBoxAtom from '../atoms/CheckBoxAtom.vue';
import { ref, computed, onMounted } from 'vue'
import { format } from 'date-fns'
import dataService from '../../services/dataService';

const infos = ref<Array<any>>([])
const currentPage = ref(1)
const totalPages = ref(1)
const limit = 10

const metodoLabels: Record<number, string> = {
  1: 'Saturação por Bases',
  2: 'Alumínio Trocável',
  3: 'Saturação de Cálcio',
  4: 'Saturação de Magnésio'
}

const buscarCalculos = async () => {
  try {
    const res = await dataService.getListaCalculos(currentPage.value, limit)
    infos.value = res.dados
    totalPages.value = res.totalPages
  } catch (err) {
    console.error('Erro ao buscar cálculos:', err)
  }
}

const makeExcel = async (info: any) => {
  try {
    const payload = {
      ctc: info.ctc,
      magnesio: info.magnesio,
      calcio: info.calcio,
      prnt: info.prnt,
      aluminio: info.aluminio,
      caO: info.caO,
      mgO: info.mgO,
      teor_ca: info.teor_ca,
      teor_mg: info.teor_mg,
      sat_atual: info.sat_atual,
      sat_desejada: info.sat_desejada,
      resultado: info.resultado,
      id_metodo: info.id_metodo,
      consultor: info.consultor,
      propriedade: info.propriedade,
      area: info.area,
      mg_desejada: info.mg_desejada,
      ca_desejada: info.ca_desejada,
    }

    await dataService.gerarExcel(payload)

  } catch (err) {
    console.error('Erro ao gerar excel no componente: ', err)
  }
}

const formatarData = (dataString: string) => {
  const data = new Date(dataString)
  return format(data, 'dd/MM/yyyy')
}

const paginaAnterior = () => {
  if (currentPage.value > 1) {
    currentPage.value--
    buscarCalculos()
  }
}

const proximaPagina = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    buscarCalculos()
  }
}

function formatarDecimalBR(valor: number | string): string {
  if (valor === null || valor === undefined || isNaN(Number(valor))) return '-'
  return valor.toString().replace('.', ',')
}


const infosReversos = computed(() => infos.value.slice().reverse())

onMounted(() => {
  buscarCalculos()
})
</script>

<template>
  <div class="checkbox-molecule">


    <div
      class="info"
      v-for="(info, index) in infosReversos.slice().reverse()"
      :key="index"
    >
      <CheckBoxAtom class="check" @exportar="makeExcel(info)">
        <p>Data: {{ formatarData(info.data_calculo) }}</p>
        <p>Método: {{ metodoLabels[info.id_metodo] }}</p>
        <p>Necessidade de calagem: {{ formatarDecimalBR(info.resultado) }} ton ha⁻¹</p>
      </CheckBoxAtom>
    </div>

    <div class="pagination" v-if="infos.length > 0">
      <button @click="paginaAnterior" :disabled="currentPage === 1"> < Anterior</button>
      <span>Página {{ currentPage }} de {{ totalPages }}</span>
      <button @click="proximaPagina" :disabled="currentPage === totalPages">Próxima ></button>
    </div>

    <div class="pagination" v-else>
      <p>Parece que ainda não há cálculos a serem exibidos...</p>
    </div>

  </div>
</template>




<style scoped>
.checkbox-molecule {
  width: 100%;
  justify-content: center;
  flex-direction: column;
  margin: 0 16.5em;
}

p {
  color: #185C37;
}

@media screen and (max-width: 1280px) {
  .checkbox-molecule {
    margin: 0 8em;
  }
}

.pagination {
  font-family: 'Konkhmer Sleokchher';
  font-size: 0.9em;
}

.pagination button {
  color: #ffffff;
  background-color: #617e4a;
  border: solid 1px transparent;
  transition: .3s ;
  margin: 0.3em;
}

.pagination button:hover {
  background-color: #8db36b;
  transform: .3s;
}

.pagination p {
  display: flex;
  justify-content: flex-start;
  font-size: 1.1em;
  color: #c5c2c2;
}

@media screen and (max-width: 769px) {
  .checkbox-molecule {
    margin: 0 5em;
  }

  .check {
    font-size: 0.7em;

  }

  .pagination {
    font-size: 0.7em;
  }
}

@media screen and (max-width: 540px) {
  .checkbox-molecule {
    margin: 0 3.6em;
  }


  .pagination {
    font-size: 0.4em;
  }
}


@media screen and (max-width: 480px) {
  .checkbox-molecule {
    margin: 0 1em;
  }

  .check {
    font-size: 0.7em;

  }
}
</style>
