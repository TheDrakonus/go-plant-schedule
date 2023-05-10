<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container">
      <RouterLink class="navbar-brand" to="/"><img class="logo-img" :src="PlantImg" />Plant Schedule</RouterLink>
      <button class="navbar-toggler" type="button" @click="isCollapsed = !isCollapsed">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" :class="{ show: !isCollapsed }">
        <ul class="navbar-nav ms-auto">
          <li class="nav-item" v-for="(item, index) in navItems" :key="index">
            <RouterLink :to="item.to" class="nav-link" exact-active-class="active">{{ item.label }}</RouterLink>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { ref } from "vue";
import PlantImg from "@/assets/plant.png";

const isCollapsed = ref(true);
const navItems = [
  { label: "Plants", to: "/plants" },
  { label: "Boxes", to: "/boxes" },
  { label: "Gardens", to: "/gardens" },
  { label: "About", to: "/about" },
  { label: "Settings", to: "/settings" },
];

</script>

<style lang="scss" scoped>
.logo-img {
  width: 2rem;
  height: 2rem;
  margin-right: 0.5rem;
}
</style>
