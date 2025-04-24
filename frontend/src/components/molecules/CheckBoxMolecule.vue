<script setup lang="ts">
import ButtonExportarAtom from '../atoms/ButtonExportarAtom.vue';
import CheckBoxAtom from '../atoms/CheckBoxAtom.vue';
import { ref } from 'vue'


const infos = ref([
  { date: '2025-04-01', method: 'PRNT', caliming: 'Alta' },
  { date: '2025-04-10', method: 'VNT', caliming: 'Média' },
  { date: '2025-04-20', method: 'Equação', caliming: 'Baixa' },
])
</script>

<template>
    <div class="checkbox-molecule">
        <div class="info" v-for="(info, index) in infos" :key="index">
            <CheckBoxAtom class="check">
            <p>Data: {{ info.date }} </p>
            <p>Método: {{ info.method }} </p>
            <p>Necessidade de calagem: {{ info.caliming }} </p>
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



