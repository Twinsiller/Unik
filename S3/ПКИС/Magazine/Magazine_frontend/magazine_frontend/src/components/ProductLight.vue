<script setup lang="ts">
import { ref } from 'vue';


export interface ProductData {
  id: number | string
  name: string
  imageUrl?: string
}

const props = defineProps<{
  product: ProductData
}>()
</script>

<template>
  <article class="product-card">
    <div class="image">
      <img
        v-if="product.imageUrl"
        :src="product.imageUrl"
        :alt="product.name"
      />
      <div v-else class="placeholder" /> 
    </div>

    <div class="info">
      <span class="label">Name</span>
      <h3 class="name">{{ product.name }}</h3>

      <span class="label">Quantity</span>
      <!-- <p class="quantity">{{ product.quantity }}</p> -->
    </div>
  </article>
</template>

<style scoped>
.product-card {
  width: 160px;
  height: 160px;
  border-radius: 12px;
  background: #ffffff;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  padding: 10px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.image {
  width: 100%;
  height: 70px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.image img {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.placeholder {
  width: 100%;
  height: 100%;
  background: #e0e0e0;
  border-radius: 6px;
}

.info {
  font-size: 12px;
}

.label {
  font-size: 10px;
  color: #888;
}

.name {
  margin: 0;
  font-size: 14px;
}

.quantity {
  margin: 0;
}

.actions {
  display: flex;
  justify-content: space-between;
}

button {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  border: none;
  background: #f2f2f2;
  cursor: pointer;
}
</style>
