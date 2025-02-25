<script setup lang="ts">
import CabecalhoMolecule from '../molecules/HeaderMolecule.vue';
import LeftHomeMolecule from '../molecules/LeftHomeMolecule.vue';
import RightHomeMolecule from '../molecules/RightHomeMolecule.vue';

</script>

<template>
  <div class="organism-homepage">
    <CabecalhoMolecule></CabecalhoMolecule>
        <div class="left-and-right">
            <div class="left">
              <LeftHomeMolecule></LeftHomeMolecule>
            </div>
            <div class="right">
              <RightHomeMolecule></RightHomeMolecule>
            </div>
        </div>
  </div>
     
     
</template>


<style>

.organism-homepage{
  display: flex;
  width: 100%;
  height: auto;
  flex-direction: column;
}
.left-and-right{
  display: flex;
  width: 100%;
  height: auto;
  margin-top: 1em;
  justify-content: stretch;
  flex-direction: row;

}

.left {
  width: 40%;
  margin: 0 5em;
  margin-top: 1em;
  padding-right: 7em;
}

@media screen and (max-width: 767px) {
  .left-and-right {
    flex-direction: column;
  }

  .left {
    width: 100%;
    margin: 0; 
    padding: 1em;
  }

  .right {
    width: 100%;
  }
}



</style>

