<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
  title: string;
}>();

const isOpen = ref(false);
</script>

<template>
  <div class="accordion-filter">
    <div class="text">
      <button class="accordion-header" @click="isOpen = !isOpen">
        {{ title }}
        <span :class="{ 'rotated': isOpen }"> ❯ </span>
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

.accordion-filter {
  width: 100%;
  border: 1px solid transparent;
  border-radius: 1em;
  overflow: hidden;
  background-color: #4f613f;

}

.text{
  transition: .3s;
}

.accordion-header {
  width: 100%;
  background-color: #4f613f;
  border-radius: 1em;
  color: white;
  text-align: left;
  padding: 0 1.4em; 
  border: none;
  font-family:'Konkhmer Sleokchher';
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
  font-size: 1em;
  transition: .3s; 
  gap: 6em;
}

.accordion-content {
  padding: 0;
  max-height: 200px;
  overflow-y: auto;
  display: flex;
  width: 100%;
  flex-direction: column;
  justify-content: space-between;
  background-color: #566844;
  border-radius: 0 0 1em 1em;
  transition: .3s;
}


.accordion-header span {
  font-size: 1.5em;
  color: rgba(179, 210, 151, 1);
  transition: transform 0.1s ease-in-out;
  transition: .3s;
  margin-right: left;
}


.rotated {
  transform: rotate(90deg);
}


.accordion-enter-active, .accordion-leave-active {
  transition: max-height 0.3s ease-in-out, opacity 0.3s ease-in-out;
  overflow: hidden;
}

.accordion-enter-from, .accordion-leave-to {
  max-height: 0px;
  opacity: 0;
}

.accordion-enter-to, .accordion-leave-from {
  max-height: 200px; 
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
