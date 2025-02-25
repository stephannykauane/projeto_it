<script setup lang="ts">
import ButtonGreenAtom from '../atoms/ButtonGreenAtom.vue';
import InputAtom from '../atoms/InputAtom.vue';
import TextAtom from '../atoms/TextAtom.vue';

</script>

<template>
<div class="login-molecule">
    <div class="caliming-text">
        <TextAtom text="CALIMING"/>
    </div>
    <hr/>
    <div class="experiencia-text">
        <TextAtom text="Faça parte dessa experiência inesquecível."/>                      
    </div>  
    <div class="inputs-login">
        <div>
           <InputAtom class="nome" placeholder="Nome"/> 
        </div>
        <div>
            <InputAtom class="email" placeholder="Email"/>
        </div>
        <div>
            <InputAtom class="senha" placeholder="Senha"/>
        </div>
        
    </div>  
    <div class="button-cadastrar">
        <ButtonGreenAtom text="Cadastrar"/>
    </div>  
    <div class="possui-conta-text">
     <TextAtom text="Já possui conta?"></TextAtom>
     <TextAtom class="entrar" text="Entrar."/>
    </div>      
    <div class="prosseguir">
        <TextAtom text="Prosseguir sem conta ➜"></TextAtom>
    </div>
</div>
</template>

<style>

.experiencia-text {
    animation: TextAtom 2s ease-out forwards;
}

@keyframes TextAtom {
    from {
        transform: scaleX(0.2);
        filter: blur(20px);
        opacity: 0;
    }
   
}
.login-molecule{
    display: flex;
    width: 100%;
    justify-content: center;
    flex-direction: column;

}

.login-molecule {
    animation: login 1.5s ease-in-out;
}

@keyframes login {
    from {
        opacity: 0;
        transform: translateX(-10px); 
    }
    to {
        opacity: 1;
        transform: translateX(0);
    }
}



.possui-conta-text{
    display: flex;
    flex-direction:column;
    justify-content: center;
    width: 100%;
    margin-top: 15px;
    font-size: 0.8em ;
    color: #CAFE9D;


}

.caliming-text{
    font-size: 2.5em;
    margin: 15px;
}

.experiencia-text{
    color: #CAFE9D;
    margin-top: 20px;
}

.inputs-login div{
    margin: 15px 15px;
}

hr{
    margin: 1px 10px;
    color: #ffffff;
}

.button-cadastrar{
    margin: 5px 50px;
   
}

.entrar {
    color: #384031 ;
    cursor: pointer;
}

.entrar:hover{
    color: #3a4630c7;
}

.prosseguir {
    font-size: 12px;
    margin-top: 10px;
    cursor: pointer;
    transition: transform 0.5s ease-in-out;
}

.prosseguir:hover{
    transform: translateX(10px);

}



@media screen and (max-width: 1025px) {
    .login-molecule{
        margin: 10px;
    }
    
}

@media screen and (max-width: 769px) {
    .login-molecule{
        margin: 7px;
    }
    
}

@media screen and (max-width: 480px) {
    .login-molecule{
        margin: 5px;
    }
    
}
</style>