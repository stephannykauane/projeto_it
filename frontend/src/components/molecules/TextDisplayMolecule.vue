<script setup lang="ts">
defineProps({
  value: {
    type: [Number, String],
    required: true
  },
  label: {
    type: String,
    required: true
  }
});
</script>

<template>
  <div class="text-display-molecule">
    <div class="label-display">
      <h2>{{ label }}</h2>  
    </div>
    <div class="value-display">
      <p>{{ value }}</p>        
    </div>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Konkhmer Sleokchher');

.text-display-molecule {
  border: 2px solid;
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  width: 100%;
  padding: 0.5em;
  border-radius: 0.5em;
  border-color: #8aa07b;
  color: #ffffff;
  margin: 0.5em 0;
  transition: .3s;
}

.label-display {
  font-family: 'Konkhmer Sleokchher'; 
  font-size: 1.1em;
  display: flex;
  text-align: left;
  align-items: center;
  width: 100%;
}

.value-display {
  width: 100%;
  font-size: 1em;
  align-self: center;
  color: #333;

}

.value-display p {
  color: #ffffff;
  text-align: end;
  margin-right: 1em;
}

@media screen and (max-width: 769px) {
  .label-display,
  .value-display {
    font-size: 0.8em;
  }
}
</style>
