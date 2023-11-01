<script setup lang="ts">
import img from '../assets/dashboard.svg'

</script>

<template>
  <div
      class="bg-gray-800 text-white justify-center items-start flex flex-row p-2 h-fit md:lg:h-screen  md:lg:w-auto  md:lg:flex-col ">
    <nav class="flex flex-wrap gap-2 justify-center
      md:lg:flex-col sm:flex-row">
      <img class="hidden md:lg:block px-1 object-contain h-48 w-48" :src="img"/>
      <router-link to="/" class="border bg-slate-500 text-white rounded p-2 hover:bg-slate-800 text-center" >Wifi</router-link>
      <router-link to="/imaalum" class="border bg-slate-500 text-white rounded p-2 hover:bg-slate-800 text-center" >IMaalum</router-link>
      <router-link to="/settings" class="border bg-slate-500 text-white rounded p-2 hover:bg-slate-800 text-center" >Settings</router-link>
      <router-link to="/about" class="border bg-slate-500 text-white rounded p-2 hover:bg-slate-800 text-center" >About</router-link>

    </nav>


  </div>
</template>