<script setup lang="ts">
import { ref } from 'vue';
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue';
import TextAtom from '../atoms/TextAtom.vue';


const editandoNome = ref(false);
const editandoSenha = ref(false);

const alternarEdicaoNome = () => { editandoNome.value = !editandoNome.value; };
const alternarEdicaoSenha = () => { editandoSenha.value = !editandoSenha.value; };
</script>

<template>
  <div class="container-molecule-alt">

    <div class="alt">
      <div class="container-alt-email">
        <TextAtom class="primeiro" text="SEU EMAIL"/>
        <TextAtom class="segundo" text="seuemail@email.com"/>
      </div>
    </div>

    <div class="alt" :class="{ 'expandido': editandoNome }">
      <div class="container-alt-nome">
        <TextAtom class="primeiro" text="SEU NOME"/>
        <template v-if="!editandoNome">
          <TextAtom class="segundo" text="seu nome"/>
        </template>
        <template v-else>
          <TextAtom class="segundo" text="seu nome"/>
          <div class="input">
            <input type="text" class="input-alterar" placeholder="Digite seu novo nome">
          </div>          
        </template>
      </div>
      <div class="button-alterar">
        <ButtonAlterarAtom text="Alterar nome" @click="alternarEdicaoNome"/>
      </div>
    </div>

    <div class="alt" :class="{ 'expandido': editandoSenha }">
      <div class="container-alt-senha">
        <TextAtom class="primeiro" text="SUA SENHA"/>
        <template v-if="!editandoSenha">
          <TextAtom class="segundo" text="*********"/>
        </template>
        <template v-else>
          <TextAtom class="segundo" text="*********"/>
          <div class="input">
            <input type="password" class="input-alterar" placeholder="Digite sua nova senha">
          </div>
        </template>
      </div>
      <div class="button-alterar">
        <ButtonAlterarAtom text="Alterar senha" @click="alternarEdicaoSenha"/>
      </div>
    </div>
  </div>
</template>

<style scoped>

.alt.expandido {
  display: flex;
  transition: 0.5s;
  padding:  1em 0.5em;
  flex-direction: column;
}


.input{
  display: flex;
  justify-content: center;
  width: 100%;
  align-items: center;
}

.input-alterar {
  display: flex;
  justify-content: center;
  padding: 0.8em;
  margin: 0 5em;
  padding-left: 1em;
  margin-top: 1em;
  font-size: 0.9em;
  color: rgb(240, 248, 255);
  font-family:'Konkhmer Sleokchher' ;
  border-radius: 5em;
  border: 1px solid #ccc;
  width: 100%;
  background-color: rgba(112, 136, 90, 0.7);
  border: 2px solid transparent;
  cursor: pointer;
  outline: none;
  text-align: center;
}

.expandido .button-alterar {
  padding-top: 0em;
}
.container-text{
  justify-content: flex-start;
}
.container-molecule-alt{
  display: flex;
  justify-content: center;
  width: 100%;
  flex-direction: column;
}

.alt {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: 0.5em 0.7em;
  background-color: rgba(202, 254, 157, 0.5);
  border-radius: 1.3em;
  margin: 1em;
  transition: 0.5s;


}

.primeiro {
  color: #384031;
  text-align: start;
}

.segundo {
  color: #ffffff;
  text-align: start;
  margin-top:0.5em;
}

.primeiro, 
.segundo {
  text-align: left;
  align-self: flex-start; 
}

.container-alt-nome,
.container-alt-email,
.container-alt-senha {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding-bottom: 2em;
  padding-left: 0.5em;
}

.button-alterar {
  display: flex;
  justify-content: center;
  align-items: center;
  padding-top: 3em;
}

@media screen and (max-width: 769px) {
  .alt{
    flex-direction: column;
  }
  .container-alt-nome,
  .container-alt-email,
  .container-alt-senha {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0.5em;
}

.button-alterar {
   padding-top: 1em;
}

.input-alterar {
  font-size: 0.8em;
  padding: 0.7em 0;
  margin: 0 2em;

}

.expandido .segundo{
  margin-bottom: 1em;
}

}

@media screen and (max-width: 480px){
  .input-alterar {
    margin: 0;
  }
  
}


</style>