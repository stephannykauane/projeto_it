<template>
    <div class="container-text">
         <h2>{{ text }}</h2>
    </div>

</template>

<script setup>
defineProps({
    text: String
})
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Konkhmer Sleokchher');
.container-text {
    display: flex;
    font-family: 'Konkhmer Sleokchher'; 
    width: 100%;
    justify-content: center;
}

h2 {
    font-size: 1em ;
}



@media screen and (max-width: 600px) {
    .container-text{
        width: 100%;
    }
  h2{
    font-size: 0.8em;
    word-wrap: break-word;
    flex-grow: 1;
  }
}
</style>