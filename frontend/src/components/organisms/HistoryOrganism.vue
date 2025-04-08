<script setup lang="ts">

import FilterMolecule from '../molecules/FilterMolecule.vue';
import HeaderAndNavMolecule from '../molecules/HeaderAndNavMolecule.vue';
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue';

</script>

<template>
 <div class="container-history-organism">
    <div class="header-and-nav">
        <HeaderAndNavMolecule/>
    </div>
    <div class="filter-things">
        <div class="history-text">
            <div class="history-text-icon">
                <div>
                   <h1>HISTÓRICO</h1>
                </div>
                <div class="history-icon">
                   <img class="icon-history" src="../../../src/assets/history-svgrepo-com.svg">
                </div>
            </div>
            <div class="filter-text-icon">
                <div>
                   <h2>FILTRAR</h2>
                </div>
                <div class="icon-2">
                   <img class="icon-filter" src="../../../src/assets/filter-svgrepo-com(1).svg">
                </div>
            </div>
        </div>
    </div>
    <div class="filter">
      <FilterMolecule/>
    </div>
    <div class="button-select">
       <ButtonAlterarAtom class="select" text="Selecionar todos"/>
    </div>
 </div>   
</template>


<style scoped>


.container-history-organism{
    display: flex;
    width: 100%;
    justify-content: center;
    flex-direction: column;
}

.history-text{
    font-family:'Konkhmer Sleokchher';
}

.history-text h1{
    font-size: 1.8em;

}

.filter {
    display: flex;
    justify-content: center;
    margin: 1em 4em;
    position: relative;
    z-index: 5;
}


.history-text-icon{
    width: 100%;
    display: flex;
    justify-content: flex-start;
    flex-direction: row;
    align-items: center;
}

.icon-history{
    width: 3.5%;
}

.icon-filter{
    width: 2.5%;
}


.filter-things{
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
    padding-left: 16.3em;
    margin-bottom: 1em;
    color: #ffffff;
}

.history-icon{
    display: flex;
    justify-content: flex-start;
    margin-left: 0.5em;
    margin-bottom: 0.3em;
}

.icon-2{
    display: flex;
    justify-content: flex-start;
    margin-left: 0.5em;
}

.filter-text-icon{
    width: 100%;
    text-align: left;
    display: flex;
    justify-content: flex-start;
    flex-direction: row;
    align-items: center;
    font-size: 1.5em;
    margin-top: 1em;
}

.select{
    width: 23.4%;
    margin-top: 3em;
}

::v-deep(.profile-alterar){
    background-color: #4f613f;
    padding: 0.4em 1em;
    margin-left: 17.8em;

}

::v-deep(.profile-alterar:hover){
    background-color: #344129;

}


</style>