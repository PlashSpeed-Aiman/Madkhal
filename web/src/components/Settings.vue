<script setup lang="ts">
import {ref} from "vue";

const matricNo = ref('')
const password = ref('')

function savePrefs(){
  fetch('/setup', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*',
    },
    body: JSON.stringify({
      username: matricNo.value,
      password: password.value
    })
  }).then(res => res.json()).then(data => {
    alert(data.Message)
  })
}
</script>

<template>
  <div class="m-auto bg-slate-100 rounded-lg border p-6 flex flex-col items-center border-slate-300">
    <h1 class="text-3xl my-2">Settings</h1>
    <p >Setup your credentials and other settings</p>
    <form class="my-2 flex flex-col gap-2"  @submit.prevent="savePrefs">
      <p class="font-semibold">Matric Number</p>
      <input class="border rounded px-1 py-1 " v-model="matricNo"/>
      <p class="font-semibold">Password</p>
      <input class="border rounded px-1 py-1 " type="password" v-model="password"/>
      <input class="border bg-slate-500 text-white rounded p-2 hover:bg-slate-800  " type="submit">
    </form>
  </div>
</template>

<style scoped>

</style>