<script setup>
import ButtonGreenAtom from '../atoms/ButtonGreenAtom.vue';
import InputAtom from '../atoms/InputAtom.vue';
import TextAtom from '../atoms/TextAtom.vue';
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import authService from '../../services/authService';

const email = ref('')
const senha = ref ('')
const confirmarSenha = ref ('')
const nome = ref('')
const erro = ref('')
const router = useRouter()  


const handleCadastro = async () => {
    if (senha.value !== confirmarSenha.value) {
        alert("As senhas não coincidem.")
        return 
    }

    try {
        await authService.signUp({
            email: email.value,
            senha: senha.value,
            nome: nome.value
        })
        router.push('/calculator')
    } catch (err) {
        erro.value = err.message
    }
}




</script>

<template>
<div class="login-molecule">
    <div class="caliming-text">
        <TextAtom text="CALIMING"/>
    </div>
    <hr/>
    <div class="experiencia-text">
        <TextAtom text="Faça parte dessa experiência"/>                      
    </div>  
    <div class="inputs-login">
        <div>
           <InputAtom class="nome" placeholder="Nome" v-model="nome"/> 
        </div>
        <div>
            <InputAtom class="email" placeholder="Email" v-model="email"/>
        </div>
        <div>
            <InputAtom class="senha" placeholder="Senha" type="password" v-model="senha"/>
        </div>
        <div>
            <InputAtom class="confirmar-senha" v-model="confirmarSenha" type="password" placeholder="Confirmar senha" />
        </div>
   
    </div>  
    <div class="button-cadastrar">
        <ButtonGreenAtom text="Cadastrar" @click="handleCadastro"/>
    </div>  
    <div class="possui-conta-text">
     <TextAtom text="Já possui conta?"></TextAtom>
     <router-link to="login">
        <TextAtom class="entrar" text="Entrar."/>
     </router-link>
    </div>      

</div>
</template>

<style scoped>

.experiencia-text {
    animation: TextAtom 2s ease-out forwards;
}

.caliming-text, hr {
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
    margin-top: 1em;
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
        margin: 0px;
    }
    
}


</style>