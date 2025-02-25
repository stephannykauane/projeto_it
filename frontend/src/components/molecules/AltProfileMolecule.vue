<script setup lang="ts">
import ButtonAlterarAtom from '../atoms/ButtonAlterarAtom.vue';
import TextAtom from '../atoms/TextAtom.vue';

</script>

<template>
    <div class="container-molecule-alt" >
      <div class="alt">
        <div class="container-alt-nome">
          <TextAtom class="primeiro" text="SEU NOME"/>
          <TextAtom class="segundo" text="seu nome"/>
        </div>
        <div button-alterar>
            <ButtonAlterarAtom text="Alterar nome"/>
        </div>
      </div>
      <div class="alt">
         <div class="container-alt-email">
          <TextAtom class="primeiro" text="SEU EMAIL"/>
          <TextAtom class="segundo" text="seuemail@email.com"/>
        </div>
        <div button-alterar>
            <ButtonAlterarAtom text="Alterar nome"/>
        </div>
      </div>
      <div class="alt">
        <div class="container-alt-senha">
          <TextAtom class="primeiro" text="SUA SENHA"/>
          <TextAtom class="segundo" text="*********"/>
        </div>
        <div button-alterar>
            <ButtonAlterarAtom text="Alterar nome"/>
        </div>  
      </div>
    </div>
</template>

<style scoped>
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
  padding: 2em 1em;
  background-color: rgba(202, 254, 157, 0.5);
  border-radius: 1em;
  margin: 2em;

}

.primeiro {
  color: #384031;
  text-align: start;
}

.segundo {
  color: #ffffff;
  text-align: start;
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
}


</style>