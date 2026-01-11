<script setup lang="ts">
  export interface Category {
    id: number
    name: string
  }
  export interface Manufacturer {
    id: number
    name: string
  }
  export interface ProductData {
    id: number | string
    name: string
    category_id: number
    category?: Category
    manufacturer_id: number | string
    manufacturer: Manufacturer
    unit: string
    barcode: string
    default_price: number
    min_stock: number
    imageUrl?: string
  }

  const props = defineProps<{
    product: ProductData
  }>()

</script>
<template>
    <div v-if="product.id">
      <img
        :src="product.imageUrl"
        :alt="product.name"
      />

      <div class="info">
        <label>Имя: </label>
        <span>{{ product.name }}</span>
      </div>
      <div class="info">
        <label>Категория: </label>
        <span>{{ product.category?.name }}</span>
      </div>
      <div class="info">
        <label>Производитель: </label>
        <span>{{ product.manufacturer?.name }}</span>
      </div>
      <div class="info">
        <label>Штрихкод: </label>
        <span>{{ product.barcode }}</span></div>
      <div class="info">
        <label>Цена обычная: </label>
        <span>{{ product.default_price }}</span>
      </div>
      <div class="info">
        <label>Минимальное количество: </label>
        <span>{{ product.min_stock }} {{ product.unit }}</span>
      </div>

      <div class="bottom">
        <button class="left" @click="">Update</button>
        <button class="right" @click="">More</button>
      </div>
    </div>
    <div v-else class="placeholder">
        Cant selected product
    </div>
</template>

<style scoped>
  /* контейнер */
  div {
    font-family: system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
    color: #1f2937;
  }

  /* изображение продукта */
  img {
    width: 100%;
    max-height: 240px;
    object-fit: contain;
    border-radius: 8px;
    background: #f3f4f6;
    margin-bottom: 16px;
  }

  /* строка с полем */
  div .info {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    padding: 6px 0;
    border-bottom: 1px solid #e5e7eb;
  }

  /* подпись */
  label {
    font-size: 13px;
    color: #6b7280;
    margin-right: 12px;
    white-space: nowrap;
  }

  /* значение */
  span {
    font-size: 14px;
    font-weight: 500;
    color: #111827;
    text-align: right;
    word-break: break-word;
  }

  /* пустое состояние */
  .placeholder {
    color: #9ca3af;
    font-style: italic;
    padding: 24px;
    text-align: center;
  }

  .bottom {
    margin-top: auto;
    display: flex;
    justify-content: space-between;
    gap: 8px;
    padding-top: 16px;
    border-top: 1px solid #e5e7eb;
  }

  .bottom button {
    padding: 8px 14px;
    border-radius: 8px;
    border: none;
    cursor: pointer;
    font-size: 14px;
  }

  .bottom .left {
    background: #2563eb;
    color: white;
  }

  .bottom .right {
    background: #e5e7eb;
  }


</style>