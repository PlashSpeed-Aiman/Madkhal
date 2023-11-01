<script setup lang="ts">
import SideBar from "./components/SideBar.vue";

</script>

<template>
  <div class="flex flex-col h-screen md:lg:flex-row md:lg:h-screen bg-green-500">
    <SideBar/>

    <router-view/>
 </div>

</template>

<style scoped>

</style>
