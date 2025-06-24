<template>
  <label class="checkbox-atom">
    <input class="checkbox" />
    <span class="label-text">
      <slot />
    </span>
    <ButtonExportarAtom class="button-exportar" text="Exportar para Excel" @click="$emit('exportar')">
    </ButtonExportarAtom>
  </label>
</template>

<script setup>
import ButtonExportarAtom from './ButtonExportarAtom.vue';

defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
})
defineEmits(['update:modelValue'])


</script>

<style scoped>

.button-exportar {
  display: flex;
  width: 20%;

}

.checkbox {
  display: flex;
}

.checkbox-atom {
  background-color: #B3CB9B;
  display: flex;
  align-items: center;
  justify-content: space-between;  
  width: 100%;
  padding: 2.5em;
  border-radius: 1em;
  text-align: start;
  cursor: pointer;
  margin: 1em 0 1em 0;
  transition: 300ms;
}

.checkbox-atom input {
  display: none;
}


.checkbox-atom:hover {
  border-color: #ffffff;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.8);
  background-color: #d4ebc0;
  transition: 300ms;
}

.checkbox-atom:focus,
.checkbox-atom:focus-visible {
  outline: 2px solid rgba(255, 255, 255, 0.5);
  transition: 300ms;
}


.checkbox-atom input:checked+.custom-checkbox::after {
  content: '';
  position: absolute;
  left: 4px;
  top: 0px;
  width: 0.5em;
  height: 12px;
  border: solid #fff;
  border-width: 0 3px 3px 0;
  transform: rotate(45deg);
}


.label-text {
  flex: 1;
}

@media screen and (max-width: 1450px) {

  .button-exportar{
    width: 30%;
  }
    
}


@media screen and (max-width: 1200px) {

.button-exportar{
  width: 40%;
}
  
}

@media screen and (max-width: 480px) {

.button-exportar{
  width: 50%;
}
  
}
</style>