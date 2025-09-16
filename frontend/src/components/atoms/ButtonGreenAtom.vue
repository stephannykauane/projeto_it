<template>
<div class="button-green">
   <input class="green-button" type="button" :value="text"/>
</div>

</template>

<script>
export default {
    props: {
        text: String,
    }
};

</script>


<style>

.button-green {
    display: flex;
    justify-content: center;
    width: 100%;
}

.green-button {
    display: inline-block;
    font-size: 0.95em; 
    font-family:'Konkhmer Sleokchher'; 
    color: #ffffff !important;
    background-color: #384031;   
    min-width: 150px;   
    text-align: center;   
}


.green-button{
    border: 2px solid transparent;
    cursor: pointer;
    width: 100%;
    padding: 0.7em;
    margin: 0 6em;
    border-radius: 1em;
    transition: .5s;
 
}

.green-button:hover{
  transition: .5s;
  background-color: #CAFE9D;
  color: #384031;
}



@media  screen and (max-width: 769px) {
  .green-button{
    margin: 0 4em;
    font-size: 0.75em;
  }
  
}


@media  screen and (max-width: 480px) {
  .green-button{
    margin: 0 2em;
    font-size: 0.7em;
  }
  
}
</style>
