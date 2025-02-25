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
    <div class="bemvindo-text">
        <TextAtom text="Bem-vindo(a) de volta, sentimos sua falta!"/>                      
    </div>  
    <div class="inputs-login">
        <div>
           <InputAtom class="email" placeholder="Email"/> 
        </div>
        <div>
            <InputAtom class="senha" placeholder="Senha"/>
        </div>
        
    </div>  
    <div class="button-entrar">
        <ButtonGreenAtom text="Entrar"/>
    </div>  
    <div class="conta-text">
     <TextAtom text="Não possui conta?"></TextAtom>
     <TextAtom class="cadastrar" text="Cadastrar."/>
    </div>      
</div>
</template>

<style>

hr{
    animation: hr 2s ease-out forwards;
}

@keyframes hr{
    from {
        transform: scaleX(0.2);
        filter: blur(20px);
        opacity: 0;
    }
}

.bemvindo-text{
    animation: TextAtom 2s ease-out forwards;
}

@keyframes TextAtom {
    from {
        transform: scaleX(0.2);
        filter: blur(20px);
        opacity: 0;
    }
   
}
.caliming-text {
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

.conta-text{
    display: flex;
    flex-direction:column;
    justify-content: center;
    width: 100%;
    margin-top: 15px;
    font-size: 0.8em ;
    color: #CAFE9D;


}

.caliming-text{
    font-size: 2em;
    margin: 15px;
}

.bemvindo-text{
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

.button-entrar{
    margin: 5px 60px;
}

.cadastrar {
    color: #384031 ;
    cursor: pointer;
}

.cadastrar:hover{
    color: #3a4630c7;
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