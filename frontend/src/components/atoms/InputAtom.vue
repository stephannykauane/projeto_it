<template>    
 <input 
    class='text-input' 
    v-model="modelValue"
    :type='type' 
    :placeholder='placeholder'>
</template>

<script setup>
const modelValue = defineModel() 
defineProps({
  placeholder: String,
  type: String,
})
</script>


<style>
.text-input{
  width: 100%;
  margin: 0.5em 0.2em;
  padding: 0.6em 2em;
  font-family:'Konkhmer Sleokchher';
  font-size: 1.1em;
  color: aliceblue;
  box-sizing: border-box;
  cursor: pointer;
  outline: none;
  border: 2px solid transparent;
  border-radius: 2em;
  background-color: rgb(184, 200, 170, 0.60);

}



.text-input:focus{
    border-color: rgba(255, 255, 255, 0.6);
    box-shadow: 0 0 15px rgba(238, 241, 237, 0.6);
} 


@media  screen and (max-width: 480px) {
   .text-input{
    padding: 0.7em 2em;
   }
  
}

</style>