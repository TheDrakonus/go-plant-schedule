<template>
  <div class="top-appbar">
    <TopBar />
  </div>
  <div class="sidebar">
    <Sidebar @update:minimized="printMinimized" />
  </div>
  <div v-show="false" class="top-navigation">
    <Nav />
  </div>
  <div class="page-container">
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router';
import Nav from './nav/Nav.vue';
import Sidebar from './nav/Sidebar.vue';
import TopBar from './nav/TopBar.vue';

const printMinimized = (minimized: boolean) => {
  console.log(minimized);
};

</script>

<style scoped>
.top-appbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 100;
  height: 3rem;
}
.page-container {
  margin-top: 60px;
  padding: 20px;
  padding-left: 13%;
  padding-right: 13%;
}

.sidebar {
  position: fixed;
  top: 3rem;
  left: 0;
  height: 100%;
  width: 13%;
  z-index: 100;
}

@media (max-width: 768px) {
  .page-container {
    padding-left: 3%;
    padding-right: 3%;
  }
}
</style>
