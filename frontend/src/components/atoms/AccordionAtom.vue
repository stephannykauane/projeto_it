<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
  title: string;
}>();

const isOpen = ref(false);
</script>

<template>
  <div class="accordion">
    <div class="text">
      <button class="accordion-header" @click="isOpen = !isOpen">
        {{ title }}
        <span :class="{ 'rotated': isOpen }">❯</span>
      </button>

      <Transition name="accordion">
        <div v-show="isOpen" class="accordion-content">
          <slot></slot>
        </div>
      </Transition>
    </div>
  </div>
</template>

<style scoped>
.accordion {
  width: 100%;
  border: 1px solid transparent;
  border-radius: 1.2em;
  overflow: hidden;
  background-color: #70885A;
  transition: 0.3s;
  position: absolute;
  top: 0;
  left: 0; 
  z-index: 10;
  transition: .3s;
}

.text{
  transition: .3s;
}

.accordion-header {
  width: 100%;
  background-color: #70885A;
  border-radius: 1.2em;
  color: white;
  text-align: left;
  padding: 0.1em 1.5em;
  border: none;
  font-family:'Konkhmer Sleokchher';
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  font-size: 1em;
  transition: .3s; 
}

.accordion-content {
  padding: 1em 0;
  display: flex;
  width: 100%;
  flex-direction: column;
  justify-content: center;
  background-color: #566844;
  border-radius: 0 0 1.1em 1.1em;
  transition: .3s;
}


.accordion-header span {
  font-size: 2em;
  color: rgba(179, 210, 151, 1);
  transition: transform 0.1s ease-in-out;
  transition: .3s;
}

.rotated {
  transform: rotate(90deg);
}


.accordion-enter-active, .accordion-leave-active {
  transition: max-height 0.3s ease-in-out, opacity 0.3s ease-in-out;
  overflow: hidden;
}

.accordion-enter-from, .accordion-leave-to {
  max-height: 0;
  opacity: 0;
}

.accordion-enter-to, .accordion-leave-from {
  max-height: 500px; 
  opacity: 1;
}


@media screen and (max-width: 900px) {
  .text {
    font-size: 0.9em;
  }
}

@media screen and (max-width: 769px) {
  .text {
    font-size: 0.7em;
  }

  .accordion{
    border-radius: 1em;
  }
}


</style>
