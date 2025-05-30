<template>
   <label class="checkbox-atom">
     <input
       type="checkbox"
       :checked="modelValue"
       @change="$emit('update:modelValue', $event.target.checked)"
     />
     <span class="custom-checkbox"></span>
     <span class="label-text"><slot /></span>
     <ButtonExportarAtom class="button-exportar" text="Exportar para Excel" @click="$emit('exportar')"></ButtonExportarAtom>
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
 .checkbox-atom {
   background-color: #B3CB9B;
   display: flex;
   align-items: center;
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
  background-color:#d4ebc0;
  transition: 300ms;
}

.checkbox-atom:focus,
.checkbox-atom:focus-visible {
  outline: 2px solid rgba(255, 255, 255, 0.5);
  transition: 300ms;
}

 .custom-checkbox {
   width: 1.2em;
   height: 1.2em;
   background-color: #70885A;
   border-radius: 0.25em;
   display: inline-block;
   position: relative;
   margin-right: 1em;
 }



 .checkbox-atom input:checked + .custom-checkbox::after {
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
 </style>
 