<!-- <script setup lang="ts">
import { ref, onMounted } from 'vue'
import Product from '@/components/Product.vue'

type Product = {
  id: number
  name: string
}

const products = ref<Product[]>([])

onMounted(async () => {
  const res = await fetch('http://localhost:8080/v1/products')
  const data = await res.json()

  products.value = data.map((p: any) => ({
    ...p,
    imageUrl: `http://localhost:8080/v1/products/${p.id}/photo`
  }))
})
</script>

<template>
  <h1>Products</h1>
  <Product
  v-for="p in products"
  :key="p.id"
  :product="p"
/>
</template> -->

<script setup lang="ts">
  import { ref, computed, onMounted } from 'vue'
  import { watch } from 'vue'
  import CreateProductModal from '@/components/CreateProductModal.vue'
  import ProductLight from '@/components/ProductLight.vue'
  import SelectedProduct from '@/components/SelectedProduct.vue'

  const products = ref<any[]>([])
  const page = ref(1)
  const perPage = ref(2)
  const selected = ref<any | null>(null)
  const showCreate = ref(false)
  const countProduct = ref(0)

  async function fetchProducts() {
    const res = await fetch(`http://localhost:8080/v1/products?page=${page.value}&limit=${perPage.value}`)
    const data = await res.json()
    
    products.value = data.items.map((p: any) => ({
      ...p,
      imageUrl: `http://localhost:8080/v1/products/${p.id}/photo`
    }))

    countProduct.value = data.count
  }

  onMounted(fetchProducts)

  watch(page, () => {
    fetchProducts()
  })

  watch(perPage, () => {
    fetchProducts()
    page.value = 1
  })

  const totalPages = computed(() =>
    Math.ceil(countProduct.value / perPage.value)
  )

</script>


<template>
  <div class="page">
    <!-- Левая часть -->
    <div class="list">
      <div class="topList">
        <button class="add-btn" @click="showCreate = true">
          + Add product
        </button> 

        <label>
          Фильтровать по:
        </label> 
        <select id="filter" v-model="perPage">
          <option :value=2>2</option>
          <option :value=5>5</option>
          <option :value=10>10</option>
          <option :value=25>15</option>
          <option :value=20>20</option>
        </select>
      </div>

      <CreateProductModal
        v-if="showCreate"
        @close="showCreate = false"
      />

      <div class="grid">
        <ProductLight v-if="products != null"
          v-for="p in products"
          :key="p.id"
          :product="p"
          @click="selected = p"
        />
        <h1 v-else>Не работает</h1>
      </div>

      <!-- Пагинация -->
      <div class="pagination">
        <button
          v-for="n in totalPages"
          :key="n"
          :class="{ active: page === n }"
          @click="page = n"
        >
          {{ n }}
        </button>
      </div>
    </div>

    <!-- Правая панель -->
    <aside class="details">
      <SelectedProduct
        v-if="selected"
        :key="selected.id"
        :product="selected"
      />
      <h1 v-else>Select Product</h1>
    </aside>

  </div>
</template>

<style>
  .page {
    display: grid;
    grid-template-columns: 1fr 380px;
    height: 100vh;
    background: #f9fafb;
    color: #111827;
    font-family: system-ui, -apple-system, BlinkMacSystemFont, sans-serif;
  }

  /* ===================== */
  /* Левая часть — список */
  /* ===================== */

  .list {
    display: flex;
    flex-direction: column;
    padding: 16px;
    overflow: hidden;
  }

  /* верхняя панель */
  .topList {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 16px;
  }

  .topList label {
    font-size: 14px;
    color: #4b5563;
  }

  .topList select {
    padding: 4px 8px;
    border-radius: 6px;
    border: 1px solid #d1d5db;
    background: white;
  }

  /* кнопка добавления */
  .add-btn {
    padding: 6px 12px;
    background: #2563eb;
    color: white;
    border: none;
    border-radius: 8px;
    cursor: pointer;
  }

  .add-btn:hover {
    background: #1d4ed8;
  }

  /* сетка товаров */
  .grid {
    flex: 1;
    overflow-y: auto;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
    gap: 12px;
    padding-right: 4px;
  }

  /* ===================== */
  /* Пагинация */
  /* ===================== */

  .pagination {
    display: flex;
    gap: 6px;
    justify-content: center;
    padding: 12px 0;
  }

  .pagination button {
    padding: 4px 8px;
    border-radius: 6px;
    border: 1px solid #d1d5db;
    background: white;
    cursor: pointer;
  }

  .pagination button.active {
    background: #2563eb;
    color: white;
    border-color: #2563eb;
  }

  /* ===================== */
  /* Правая панель — details */
  /* ===================== */

  .details {
    border-left: 1px solid #e5e7eb;
    background: white;
    padding: 16px;
    overflow-y: auto;
  }

  .details h1 {
    font-size: 18px;
    color: #9ca3af;
    text-align: center;
    margin-top: 40px;
  }

</style>