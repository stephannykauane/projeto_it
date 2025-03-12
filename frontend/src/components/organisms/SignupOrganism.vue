<script setup lang="ts">
import SignUpMolecule from '../molecules/SignUpMolecule.vue';

</script>

<template>
<div class="container-signup-form">
    <div class="signup-form">
        <SignUpMolecule></SignUpMolecule>
    </div>
</div>   



</template>

<style>

.container-signup-form{
    display: flex;
    width: 100%;
    justify-content: center;
    padding: 0 20em;
    
  
}

.signup-form{
    display: flex;
    width: 100%;
    justify-content: center;
    background-color: rgba(147, 168, 126, 0.6);
    padding: 8em 3em;
    border-radius: 0.5em ;
    box-shadow: 0.5px 5px 30px rgb(0, 0, 0, 0.25);;
}


@media screen and (max-width: 1025px){
    .container-signup-form{
        padding: 0em 10em;
        max-height: 800px;
   
    }

    .signup-form{
        padding: 8em 3em;
    }
}


@media screen and (max-width: 769px){
    .container-signup-form{
        padding: 0 6em;
        max-height: 650px ;
    }

    .signup-form{
        padding: 7em 3em;
    }
}


@media screen and (max-width: 480px){
    .container-signup-form{
        padding: 0;
    }

    .signup-form{
        padding: 6em 3em;
    }
}
</style>