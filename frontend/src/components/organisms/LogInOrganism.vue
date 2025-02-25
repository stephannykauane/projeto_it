<script setup lang="ts">
import LogInMolecule from '../molecules/LogInMolecule.vue';

</script>

<template>
<div class="container-login-form">
    <div class="login-form">
        <LogInMolecule></LogInMolecule>
    </div>
</div>   



</template>

<style>

.container-login-form{
    display: flex;
    width: 100%;
    height: auto;
    justify-content: center; 
    padding: 2.3em 20em;
}

.login-form{
    display: flex;
    width: 100%;
    justify-content: center;
    background-color: rgba(147, 168, 126, 0.6);
    padding: 12.25em 2.7em;
    border-radius: 0.5em ;
    box-shadow: 0.5px 5px 30px rgb(0, 0, 0, 0.25);;
}


@media screen and (max-width: 1025px){
    .container-login-form{
        padding: 5em 13em;

    }

    .login-form{
        padding: 8em 3em;
    }
}


@media screen and (max-width: 769px){
    .container-login-form{
        padding: 5em 6em;
    }

    .login-form{
        padding: 7em 3em;
    }
}


@media screen and (max-width: 480px){
    .container-login-form{
        padding: 0;
    }

    .login-form{
        padding: 6em 3em;
    }
}
</style>