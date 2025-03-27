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
    <div class="text">
      <h2>{{ text }}</h2>  
    </div>
    <div class="input">
        <InputAtom 
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
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
  width: 100%;
  margin: 1em 0;
  gap: 2em;
  transition: .3s;
}

.text {
  font-family: 'Konkhmer Sleokchher'; 
  font-size: 1.2em;
  text-align: end;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  width: 20%;
  transition: .3s;
}

.input {
  width: 80%;
  font-size: 1em;
  transition: .3s;
}


:deep(.text-input) {
    border-radius: 1em;
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
