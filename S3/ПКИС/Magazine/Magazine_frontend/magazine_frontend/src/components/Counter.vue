<script setup lang="ts">
import { ref } from 'vue'

const count = ref(0)

function inc() {
  count.value++
}

function dec() {
  count.value--
}
</script>

<template>
  <div class="counter">
    <button @click="dec">-</button>
    <span>{{ count }}</span>
    <button @click="inc">+</button>
  </div>
</template>

<style scoped>
.counter {
  display: inline-flex;
  gap: 0.5rem;
  align-items: center;
}
</style>
