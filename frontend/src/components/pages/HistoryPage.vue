<script setup lang="ts">
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue';
import CheckBoxAtom from '../atoms/CheckBoxAtom.vue';
import CalculatorOrganism from '../organisms/CalculatorOrganism.vue';
import HistoryOrganism from '../organisms/HistoryOrganism.vue';
import { ref } from 'vue'

const teste = ref(false);
</script>

<template>
    <CheckBoxAtom v-model="teste">
        <div class="testednv">
            <div>
        <p>data: dd/mm/yy </p>
        <p>Método: Alumínio Trocável</p>
        <p>Necessidade de calagem: XX ton/ha</p> 
        </div>
        <div class="af">
            <ButtonAlterarAtom text="Exportar para excel"/>
        </div>
        </div>


    </CheckBoxAtom>
</template>


<style scoped>
 p{
    color: #185C37;
 }

 .testednv{
    display: flex;
    width: 100%;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
 }

 .af{

    display: flex;
    justify-content: flex-end;
    

 }

</style>

