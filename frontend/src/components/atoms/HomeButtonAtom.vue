<template>
  <div class="button-container">
    <input class="home-button" type="button" :value="text" />
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
@import url('https://fonts.googleapis.com/css2?family=Arya:wght@400;700&family=Signika:wght@300..700&display=swap');

.button-container {
  display: flex;
  justify-content: center;
  width: 100%;
}

.home-button {
  display: flex;
  justify-content: center;  
  background: transparent;
  outline: none;
  cursor: pointer;
  box-sizing: border-box;
  padding: 0.2em 2.5em;
  max-width: 200px;
  height: auto;
  text-align: center;
  margin: 0;
}

.home-button{
  font-size: 1.2em; 
  font-family: 'Arya', serif;
  color: white;
  border: 2px solid #ffffffc7;
  border-radius: 50px;
  transition: box-shadow 200ms ease-in-out, border-color 200ms;
}

.home-button:hover {
  border-color: #ffffff;
  box-shadow: 0 0 8px rgba(255, 255, 255, 0.8);
  transition: 300ms;
}

.home-button:focus,
.home-button:focus-visible {
  outline: 2px solid rgba(255, 255, 255, 0.5);
}


@media screen and (max-width: 600px) {
  .home-button {
    font-size: 1em; 
    padding: 0.3em 2.5em; 
    min-width: 90px;
  }
}

</style>
