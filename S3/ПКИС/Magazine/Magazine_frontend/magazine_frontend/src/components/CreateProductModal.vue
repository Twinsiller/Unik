<script setup lang="ts">
  import { onMounted, ref } from 'vue';
  import CreateCategoryModal from '@/components/CreateCategoryModal.vue'
  import CreateManufacturerModal from '@/components/CreateManufacturerModal.vue'

  const emit = defineEmits(['close'])

  type CreateProductPayload = {
    name: string
    category_id?: number
    manufacturer_id?: number
    unit: string
    barcode: string
    default_price: number
    min_stock: number
  }

  const form = ref<CreateProductPayload>({
    name: '',
    category_id: undefined,
    manufacturer_id: undefined,
    unit: '',
    barcode: '',
    default_price: 0,
    min_stock: 0,
  })

  const categories = ref<any[]>([])
  const manufacturers = ref<any[]>([])
  var temp = 0

  async function fetchCategories() {
    const res = await fetch(`http://localhost:8080/v1/categories`)
    const data = await res.json()
    categories.value = data.map((c: any) => ({
        ...c
    }))

    console.log(categories.value)
    temp = temp + 1
    console.log(temp)
  }
  async function fetchManufacturers() {
    const res = await fetch(`http://localhost:8080/v1/manufacturers`)
    const data = await res.json()
    manufacturers.value = data.map((m: any) => ({
        ...m
    }))

    console.log(manufacturers.value)
    temp = temp + 1
    console.log(temp)
  }



  onMounted(() => {
    fetchCategories()
    fetchManufacturers()
  })

  const imageFile = ref<File | null>(null)
  async function CreateProduct(){
    if(!form.value.name.trim()){
      alert('Название не может быть пустым')
      return
    }

    if(!form.value.category_id){
      alert('Выберите категорию')
      return
    }

    if(!form.value.manufacturer_id){
      alert('Выберите производителя продукта')
      return
    }

    if(!form.value.barcode.trim()){
      alert('Штрихкод не может быть пустым')
      return
    }

    if(!form.value.default_price){
      alert('Введите цену продукта')
      return
    }

    if(!form.value.min_stock){
      alert('Введите минимальное количество продуктов')
      return
    }

    form.value.default_price = Number(form.value.default_price)
    form.value.min_stock   = Number(form.value.min_stock)
    console.log(form.value)

    const res = await fetch('http://localhost:8080/v1/products', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(form.value)
    })

    if (!res.ok) {
      const text = await res.text()
      throw new Error(text)
    }

    const createdProduct = await res.json()
    const productId = createdProduct.id

    if (imageFile.value) {
      const fd = new FormData()
      fd.append('file', imageFile.value)

      await fetch(`http://localhost:8080/v1/products/${productId}/photo`, {
        method: 'POST',
        body: fd
      })
    }

    // если сервер возвращает созданную категорию
    console.log('created product', createdProduct)

    // закрываем модалку
    emit('close')
  }

  

  function onFileChange(e: Event) {
    const files = (e.target as HTMLInputElement).files
    imageFile.value = files?.[0] ?? null
  }



  const showCreateCategory = ref(false)
  const showCreateManufacturer = ref(false)
</script>

<template>
  <div class="overlay" @click.self="$emit('close')">
    <div class="modal">
      <h2>Add product</h2>

      <input id="input-name" placeholder="Введите название" v-model="form.name"/>
      <input
        type="file"
        accept="image/*"
        @change="onFileChange"
      />

      <label for="select-unit">Выберите единицу измерения</label>
      <select id="select-unit" v-model="form.unit">
        <option value="шт">шт</option>
        <option value="кг">кг</option>
      </select>

      <label for="select-category">Выберите категорию товара</label>
      <select id="select-category" v-model="form.category_id">
        <option v-for="c in categories"
          :key="c.id"
          :value="c.id"
          >{{ c.name }}</option>      
      </select>

      <label for="select-manufacturer">Выберите производителя товара</label>
      <select id="select-manufacturer" v-model="form.manufacturer_id">
        <option v-for="m in manufacturers"
          :key="m.id"
          :value="m.id"
          >{{ m.name }}</option>      
      </select>

      <button class="add-btn" @click="showCreateCategory = true">
          Добавить новую категорию
      </button>
      
      <button class="add-btn" @click="showCreateManufacturer = true">
          Добавить нового производителя
      </button>

      <label for="input-barcode">Введите штрихкод товара</label>
      <input id="input-barcode" placeholder="Введите штрихкод" v-model="form.barcode"/>
      
      <label for="input-default_price">Введите цену товара</label>
      <input id="input-default_price" placeholder="Введите цену" v-model="form.default_price"/>
      
      <label for="input-min_stock">Введите минимальное количество товара</label>
      <input id="input-min_stock" placeholder="Введите минимальное количество" v-model="form.min_stock"/>

      
      
      <button @click="$emit('close')">Close</button>
      <button @click="CreateProduct">Create</button>
    </div>
    
    <CreateCategoryModal
        v-if="showCreateCategory"
        @close="showCreateCategory = false"
        @update="fetchCategories"
    />
    <CreateManufacturerModal
        v-if="showCreateManufacturer"
        @close="showCreateManufacturer = false"
        @update="fetchManufacturers"
    />
  </div>
</template>

<style scoped>
  .overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.45);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  /* ===================== */
  /* Модалка */
  /* ===================== */

  .modal {
    background: #ffffff;
    padding: 20px;
    border-radius: 12px;
    width: 420px;
    max-height: 90vh;
    overflow-y: auto;
    box-shadow: 0 10px 30px rgba(0,0,0,0.15);
  }

  .modal h2 {
    margin-bottom: 16px;
    font-size: 18px;
    font-weight: 600;
  }

  /* ===================== */
  /* Поля формы */
  /* ===================== */

  .modal input,
  .modal select {
    width: 100%;
    padding: 8px 10px;
    margin-bottom: 10px;
    border-radius: 8px;
    border: 1px solid #d1d5db;
    font-size: 14px;
  }

  .modal input[type="file"] {
    padding: 4px;
  }

  .modal label {
    display: block;
    margin: 8px 0 4px;
    font-size: 13px;
    color: #4b5563;
  }

  /* фокус */
  .modal input:focus,
  .modal select:focus {
    outline: none;
    border-color: #2563eb;
    box-shadow: 0 0 0 1px #2563eb;
  }

  /* ===================== */
  /* Кнопки */
  /* ===================== */

  .add-btn {
    width: 100%;
    margin-top: 6px;
    padding: 8px;
    border-radius: 8px;
    border: 1px dashed #2563eb;
    background: #eff6ff;
    color: #2563eb;
    cursor: pointer;
    font-size: 14px;
  }

  .add-btn:hover {
    background: #dbeafe;
  }

  /* нижние кнопки */
  .modal button:last-of-type,
  .modal button:nth-last-of-type(2) {
    margin-top: 12px;
  }

  .modal button {
    padding: 8px 12px;
    border-radius: 8px;
    border: none;
    cursor: pointer;
    font-size: 14px;
  }

  .modal button:nth-last-of-type(2) {
    background: #e5e7eb;
  }

  .modal button:last-of-type {
    background: #2563eb;
    color: white;
  }

  .modal button:last-of-type:hover {
    background: #1d4ed8;
  }

</style>
