<script setup lang="ts">
import SideBar from "./components/SideBar.vue";

</script>

<template>
  <div class="flex flex-col h-screen md:lg:flex-row md:lg:h-screen bg-[#607D8B]">
    <SideBar/>
    <transition name="fade">
      <router-view />
    </transition>
 </div>

</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>
