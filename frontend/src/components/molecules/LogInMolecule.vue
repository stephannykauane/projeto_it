<script setup lang="ts">
import { useRouter } from 'vue-router';
import ButtonGreenAtom from '../atoms/ButtonGreenAtom.vue';
import InputAtom from '../atoms/InputAtom.vue';
import TextAtom from '../atoms/TextAtom.vue';
import { ref } from 'vue'
import calimingAPI from '../../services/CalimingAPIClient';


const email = ref('')
const senha = ref('')
const erro = ref('')
const router = useRouter()

const handleLogin = async () => {
    try {
        const success = await calimingAPI.login({
            email: email.value,
            senha: senha.value
        })
        if (success) {
            router.push('/calculator')
        } else {
            erro.value = 'Email ou senha inválidos'
        }
    } catch (err) {
        erro.value = 'Erro inesperado: ' + (err instanceof Error ? err.message : String(err))
    }
}

</script>

<template>
    <div class="login-molecule">
        <div class="caliming-text">
            <TextAtom text="CALIMING" />
        </div>
        <hr />
        <div class="bemvindo-text">
            <TextAtom text="Bem-vindo(a) de volta, sentimos sua falta!" />
        </div>
        <div class="inputs-login">
            <div>
                <InputAtom class="email" placeholder="Email" v-model="email" />
            </div>
            <div>
                <InputAtom class="senha" placeholder="Senha" v-model="senha" type="password" />
            </div>
        </div>
        <div v-if="erro"> {{ erro }}</div>
        <div class="button-entrar">
            <ButtonGreenAtom text="Entrar" @click="handleLogin" />
        </div>
        <div class="conta-text">
            <TextAtom text="Não possui conta?"></TextAtom>
            <router-link to="signup">
                <TextAtom class="cadastrar" text="Cadastrar." />
            </router-link>

        </div>
    </div>
</template>

<style scoped>


.caliming-text, .bemvindo-text, hr {
    animation: TextAtom 2s ease-out forwards;
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

@keyframes TextAtom {
    from {
        transform: scaleX(0.2);
        filter: blur(20px);
        opacity: 0;
    }

}

.login-molecule {
    display: flex;
    width: 100%;
    justify-content: center;
    flex-direction: column;

}

.conta-text {
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 100%;
    margin-top: 15px;
    font-size: 0.8em;
    color: #CAFE9D;
}

.caliming-text {
    font-size: 2.5em;
    margin: 15px;
}

.bemvindo-text {
    color: #CAFE9D;
    margin-top: 20px;

}

.inputs-login div {
    margin: 15px 15px;
}

hr {
    margin: 1px 10px;
    color: #ffffff;
}

.button-entrar {
    margin: 5px 60px;
}

.cadastrar {
    color: #384031;
    cursor: pointer;
}

.cadastrar:hover {
    color: #3a4630c7;
}

@media screen and (max-width: 1025px) {
    .login-molecule {
        margin: 10px;
    }

}

@media screen and (max-width: 769px) {
    .login-molecule {
        margin: 7px;
    }

}

@media screen and (max-width: 480px) {
    .login-molecule {
        margin: 5px;
    }

}
</style>