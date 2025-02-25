<template>
   <div class="container-description">
    <div>
        <h3>{{ text }}</h3>
    </div>
    <div>
        <p>{{ subtext }}</p>
    </div>
   </div>
</template>

<script setup>
defineProps({
    text: String,
    subtext: String
});
</script>


<style>

@import url('https://fonts.googleapis.com/css2?family=Konkhmer+Sleokchher&display=swap');

.container-description {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  background-color: hsla(90, 19%, 58%, 0.6);
  border-radius: 0.5rem;
  width: 100%;
  max-width: 600px; 
  min-height: 120px; 
  padding: 2.1em 2em;
  margin: 5px 0;

}

h3, p {
  font-family: 'Konkhmer Sleokchher', sans-serif;
  text-align: start;
  margin: 0;
  max-width: 100%;
  word-wrap: break-word;
}

h3 {
  font-size: 1.3em;
  display: inline-flex;
  align-items: center;
  white-space: wrap;
} 

p {
  font-size: 1em;
  color: #ffffff99;
  flex-grow: 1; 

}


@media screen and (max-width: 600px) {
  .container-description {
    padding: 1.2em;
    min-height: 120px; 
  }

  h3 {
    font-size: 1em;
  }

  p {
    font-size: 0.8em;
  }
}


</style>