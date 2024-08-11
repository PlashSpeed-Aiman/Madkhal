<script setup lang="ts">
import img from '../assets/dashboard.svg'
import wifi from '../assets/wifi_24dp.svg'
import settingsIcon from '../assets/settings_24.svg'
import accountIcon from  '../assets/manage_accounts_24dp.svg'
import aboutIcon from '../assets/help_24dp.svg'
</script>

<template>
  <div
      class=" bg-[#455A64] shadow-xl text-white justify-center items-start flex flex-row p-2 h-fit md:lg:h-screen  md:lg:w-auto  md:lg:flex-col ">
    <nav class="flex flex-wrap gap-2 justify-center
      md:lg:flex-col sm:flex-row">
      <img class="hidden md:lg:block px-1 object-contain h-48 w-48" :src="img"/>
      <router-link to="/" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center" >
        <img class="mx-auto my-1" :src="wifi"/>
        Wifi
      </router-link>
      <router-link to="/imaalum" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center" >
        <img class="mx-auto my-1" :src="accountIcon"/>
        IMaalum</router-link>
      <router-link to="/settings" class="group  border  border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center"  >
        <img class="mx-auto my-1 hover:fill-white hover:stroke-white" :src="settingsIcon"/>
        Settings</router-link>
      <router-link to="/about" class="border border-[#455A64] bg-[#CFD8DC] text-[#212121] transition-colors duration-100 rounded p-2 hover:bg-[#607D8B] hover:text-white text-center"  >
        <img class="mx-auto my-1" :src="aboutIcon"/>
        About</router-link>

    </nav>


  </div>
</template>