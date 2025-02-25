<template>
    <div class="button-sair">
       <input class="sair-button" type="button" :value="text"/>
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

.button-sair {
    display: flex;
    justify-content: center;
    width: 100%;
}

.sair-button {
    display: inline-block;
    color: aliceblue;
    background-color: #384031;
    font-family:'Konkhmer Sleokchher';
    font-size: 0.90em;
    min-width: 150px;
    text-align: center;

}

.sair-button{
    border: 2px solid transparent;
    cursor: pointer;
    width: 100%;
    border-radius: 10em;
    padding: 0.4em 5em;
    margin: 5px;
    transition: .5s;
}

.sair-button:hover{
    transition: .5s;
    background-color: #c02e2e;
}


</style>
    