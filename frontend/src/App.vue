<template>
  <RouterView/>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Signika+Negative:wght@300..700&display=swap');

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  min-width: 100vw;
  min-height: 100vh;
  overflow-x: hidden;
}

body {  
  display: flex;
  flex-direction: column; 
  justify-content: center;
  background: linear-gradient(to bottom, #384031, #93A87E);
  font-family: 'Signika Negative', sans-serif;
}

div#app {
  width: 100%;
  min-height: 100%;
  display: flex;
  flex-direction: column;
 
  justify-content: center;
  flex-grow: 1;
 
}

@media screen and (max-width: 767px) {
  body {
    padding: 0 2px;
  }



}


</style>
