<template>
  <nav>
    <RouterLink to="/">Главная</RouterLink>
    |
    <RouterLink :to="{ name: '/product', params: {id: 1}}">Продукт</RouterLink>
    |
    <RouterLink to="/login">Вход</RouterLink>
    |
    <RouterLink to="/register">Регистрация</RouterLink>
  </nav>

  <RouterView />
</template>
