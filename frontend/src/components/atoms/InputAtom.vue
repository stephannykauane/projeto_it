<template>    
 <input 
    class='text-input' 
    v-model="modelValue"
    :type='type' 
    :placeholder='placeholder'
    >
</template>

<script setup>
const modelValue = defineModel() 
defineProps({
  placeholder: String,
  type: String,
})
</script>


<style scoped>
.text-input{
  width: 100%;
  margin: 0.1em 0.1em;
  padding: 0.6em 1em;
  font-family:'Konkhmer Sleokchher';
  font-size: 1.1em;
  color: aliceblue;
  box-sizing: border-box;
  cursor: pointer;
  outline: none;
  border: 2px solid transparent;
  text-align: left;
  border-radius: 0.7em;
  background-color: rgb(184, 200, 170, 0.60);

}

input {
  text-align: left;
  color: #ffffff !important;
}


.text-input::placeholder {
  color: #ffffffc5;
  opacity: 1; 
}


.text-input::-webkit-input-placeholder {
  color: #ffffffc5;
}

.text-input:focus{
    border-color: rgba(255, 255, 255, 0.6);
    box-shadow: 0 0 15px rgba(238, 241, 237, 0.6);
} 

@media  screen and (max-width: 769px) {
   .text-input{
    padding: 0.7em 2em;
    font-size: 0.8em;
   }
  
}



</style>