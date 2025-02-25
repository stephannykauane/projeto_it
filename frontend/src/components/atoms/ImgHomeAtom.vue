<template>
   <div class="container-img">
      <img src="/home/stephanny/projeto_it/frontend/src/assets/maos.jpg">
   </div>
</template>

<style>
    img {
      width: 50%;
      height: auto; 
      image-rendering: crisp-edges;
      border-radius: 50%;
      box-shadow: 0.5px 5px 30px rgb(0, 0, 0, 0.25);

    }
    
    .container-img{
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;

    }


    @media screen and (max-width: 980px) {
      .container-img{
        width: 70%;
      }
    }  




   

  
</style>
