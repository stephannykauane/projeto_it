<script setup lang="ts">
import InputAtom from '../atoms/InputAtom.vue';
import TextAtom from '../atoms/TextAtom.vue';

defineProps({
  text: String,        
  modelValue: String, 
  placeholder: String, 
  type: {               
    type: String,
    default: "text"
  }
});

defineEmits(["update:modelValue"]);
</script>

<template>
  <div class="text-input-molecule">
    <div class="text-textInput">
      <h2>{{ text }}</h2>  
    </div>
    <div class="input-textInput">
        <InputAtom 
      style="text-align:left;"
      :type="type" 
      :placeholder="placeholder" 
      :modelValue="modelValue"
      @update:modelValue="$emit('update:modelValue', $event)"
       />
    </div>
  
  </div>
</template>

<style scoped>

@import url('https://fonts.googleapis.com/css?family=Konkhmer Sleokchher');

.text-input-molecule {
  border: 2px solid;
  display: flex;
  justify-content: center;
  flex-direction: column;
  width: 100%;
  padding: 1em;
  border-radius: 0.5em;
  border-color: #8aa07b;
  margin: 0.2em 0;
  transition: .3s;
}

.text-textInput {
  font-family: 'Konkhmer Sleokchher'; 
  font-size: 1em;
  display: flex;
  text-align: left;
  align-items: center;
  width: 100%;
  transition: .3s;
}

.input-textInput {
  width: 100%;
  font-size: 1em;
  transition: .3s;
}


:deep(.text-input) {
    border-radius: 0.7em;
    background-color: rgba(103, 120, 87, 1) ;
    text-align: center;
}


@media screen and (max-width:769px) {
  .text{
    font-size: 0.8em;
  }

  .input{
    font-size: 0.8em;

  }
  
  
}
</style>