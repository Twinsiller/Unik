<script setup lang="ts">
  import { ref } from 'vue';

  const emit = defineEmits(['close','update'])

  type CreateCategoryPayload = {
    id: number
    name: string
  }

  const form = ref<CreateCategoryPayload>({
    id: -1,
    name: ''
  })

  async function CreateCategory() {
    if (!form.value.name.trim()) {
      alert('Название не может быть пустым')
      return
    }

    const res = await fetch('http://localhost:8080/v1/categories', {
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

    // если сервер возвращает созданную категорию
    const created = await res.json()
    console.log('created category', created)

    // закрываем модалку
    emit('update')
    emit('close')
  }

</script>

<template>
  <div class="overlay" @click.self="$emit('close')">
    <div class="modal">
      <h2>Add category</h2>

      <input id="input-name" placeholder="Введите название" v-model="form.name"/>

      <button @click="$emit('close')">Close</button>
      <button @click="CreateCategory">Create</button>
    </div>
  </div>
</template>

<style scoped>
.overlay {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.4);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  background: white;
  padding: 16px;
  border-radius: 12px;
  width: 400px;
}
</style>
