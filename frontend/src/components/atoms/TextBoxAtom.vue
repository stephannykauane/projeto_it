<template>
    <div class="text-box">
        <h3>{{ text }}</h3>

    </div>
</template>

<script setup>

defineProps ({
    text: {
        type: String
    }
})
</script>

<style scoped>

.text-box {
   display: flex;
   justify-content: center;
   text-align: center;
   width: 100%;
   background-color: #B3D297;
   border-radius: 1em;
   padding: 3em 3em;
   margin: 1em 3.5em;
   transition: .3s;
   max-height: 200px;
}

h3 {
    text-align: center;
    align-self: center;
    font-family:'Konkhmer Sleokchher';
    font-size: 1em;
    color: #5D6B50;
    transition: .3s;
}

@media screen and (max-width: 1500px){
    .text-box{
        padding: 3em 2.6em;
        margin: 1em 3em;
        border-radius: 0.8em;
        max-height: 170px;
    }

    h3 {
        font-size: 1em;
    }
}

@media screen and (max-width: 1200px){
    .text-box{
        padding: 1em 2.5em;
        margin: 1em 2.2em;
        border-radius: 0.8em;
        max-height: 160px;
    }

    h3 {
        font-size: 0.9em;
    }
}


@media screen and (max-width: 900px){
    .text-box{
        padding: 2em 1em;
        margin: 1em 1em;
        border-radius: 0.8em;
        max-height: 150px;
    }

    h3 {
        font-size: 0.8em;
    }
}

@media screen and (max-width: 769px){
    .text-box{
        padding: 2em 1em;
        margin: 1em 0.5em;
        border-radius: 0.8em;
    }

    h3 {
        font-size: 0.8em;
    }
} 

@media screen and (max-width: 480px){
    .text-box{
        padding: 1em 0.2em;
        margin: 1em 0.3em;
        border-radius: 0.6em;
    }

    h3 {
        font-size: 0.7em;
    }
} 
</style>