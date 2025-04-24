
<template>
    <div class="filter-molecule">
        <div>
        <FilterAccordionAtom title="Dia">
          <button
            v-for="day in days"
            :key="day"
            @click="selecionar('day', day)"
            class="dia-filter" 
           >
            {{ day }}
          </button>
        </FilterAccordionAtom> 
        </div>
        <div>
        <FilterAccordionAtom title="Mês">
          <button
            v-for="month in months"
            :key="month"
            @click="selecionar('month', month)"
            class="mes-filter" 
           >
            {{ month }}
          </button>
        </FilterAccordionAtom>  
        </div>
        <div>
        <FilterAccordionAtom title="Ano ">
          <button
            v-for="year in years"
            :key="year"
            @click="selecionar('year', year)"
            class="ano-filter" 
           >
            {{ year }}
          </button>
        </FilterAccordionAtom>
        </div>
        <div>
        <FilterAccordionAtom class="method-accordion" title="Método">
          <button
            v-for="method in methods"
            :key="method"
            @click="selecionar('method', method)"
            class="method-filter" 
           >
            {{ method }}
          </button>
        </FilterAccordionAtom>
        </div>
    </div>
</template>


<script>


import FilterAccordionAtom from '../atoms/FilterAccordionAtom.vue';

export default {
    components: { FilterAccordionAtom },
    data() {
        return {
            filtros: {
                day:"",
                month:"",
                year:"",
                method:"",
            },
            days: Array.from({ length: 31 }, (_, i) => i + 1),
            months: Array.from({ length: 12 }, (_, i) => i + 1),
            years: Array.from({ length: 50 }, (_, i) => 2025 + i),
            methods:["Saturação por bases", "Alumínio Trocável"],
        };
    },
    methods: {
        selecionar(tipo, valor) {
            this.filtros[tipo] = this.filtros[tipo] === valor ? "" : valor;
        },
        aplicar() {
            this.$emit("applyFilters", this.filtros);
        },
    },
};

</script>


<style scoped>

.filter-molecule{
    font-family: 'Konkhmer Sleokchher';
    width: 78%;
    display: flex;
    margin-left: 12.3em;
    justify-content: space-between;
    flex-direction: row;
    position: absolute;
    top: 0;
    left: 0; 
    z-index: 10;
}


.mes-filter, .dia-filter, .ano-filter, .method-filter{
    background-color: #9cb387cb;
    width: 100%;
}


.mes-filter:hover, .dia-filter:hover, .ano-filter:hover, .method-filter:hover{
    background-color: #75866483;
    border: 0.1px solid transparent;

}

::v-deep(.accordion-content){
    background-color: #9cb387;
}


</style>