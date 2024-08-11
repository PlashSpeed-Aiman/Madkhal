<script setup lang="ts">

const login = () => {
  fetch('http://localhost:8080/login', {
    method: 'GET',
    headers: {
      'Access-Control-Allow-Origin': '*',
    },
  }).then(res => res.json()).then(data => {
    alert(data.Message)
  }).catch(err => {
    alert(err)
  })
}
</script>

<template>
  <div class=" mx-auto my-auto bg-slate-100 rounded-lg shadow">
    <div class=" flex flex-col items-center border p-6 rounded-lg  ">
      <h1 class="text-3xl my-2">Wifi</h1>
      <p class="text-center items-center">Makes sure to setup your credentials in the settings page</p>
      <button v-on:click="login" class="my-2 md:lg:w-2/4 w-full border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center">Login</button>
      <button v-on:click="login" class="md:lg:w-2/4 w-full border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center">Logout</button>
    </div>
  </div>

</template>

<style scoped>

</style>