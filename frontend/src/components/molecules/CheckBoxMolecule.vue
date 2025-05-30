<script setup lang="ts">
import calimingAPI from '../../services/CalimingAPIClient';
import CheckBoxAtom from '../atoms/CheckBoxAtom.vue';
import { ref, onMounted } from 'vue'
import { format } from 'date-fns'


const infos = ref<Array<{ data_calculo: string; id_metodo: number; resultado: string }>>([])


const buscarCalculos = async () => {
  try {
    infos.value = await calimingAPI.getListaCalculos()
    
  } catch (err) {
    console.error('Erro ao buscar cálculos:', err)
  }
}

const makeExcel = async (info) => {
  try {
    const payload = {
      ctc: info.ctc,
      magnesio: info.magnesio,
      argila: info.argila,
      calcio: info.calcio,
      prnt: info.prnt,
      potassio: info.potassio,
      aluminio: info.aluminio,
      sat_desejada: info.sat_desejada,
      resultado: info.resultado,
      id_metodo: info.id_metodo
    }
    console.log(payload)
    const resp = await calimingAPI.gerarExcel(payload)

    if (resp.ok) {
      console.log('Excel gerado com sucesso')
    }
  } catch (err) {
    console.error('Erro ao gerar excel no componente: ', err)
  }
}

const formatarData = (dataString: string) => {
  const data = new Date(dataString)
  return format(data, 'dd/MM/yyyy')
}
onMounted(() => {
  buscarCalculos()
})
</script>

<template>
    <div class="checkbox-molecule">
        <div class="info" v-for="(info, index) in infos.slice().reverse()" :key="index">
          <CheckBoxAtom class="check" @exportar="makeExcel(info)">
            <p>Data: {{ formatarData(info.data_calculo) }} </p>
            <p>Método: {{ info.id_metodo }} </p>
            <p>Necessidade de calagem: {{ info.resultado }} </p>
         </CheckBoxAtom>
        </div>
    </div>
</template>


<style scoped>

.checkbox-molecule{
    width: 100%;
    justify-content: center;
    flex-direction: column;
    margin: 0 16.5em;
}

p{
    color: #185C37;
}
</style>



