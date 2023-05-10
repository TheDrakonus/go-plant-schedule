<template>
  <div>
    <h1>Boxes</h1>
    <p>Rows:</p>
    <input class="" type="number" v-model="rows" />
    <p>Cols:</p>
    <input type="number" v-model="columns" />
  </div>
  <GardenBox :rows="rows" :columns="columns" />
</template>

<script setup lang="ts">
import { ref } from "vue";
import GardenBox from "../components/GardenBox.vue";

// props
interface Props {
  rows: number;
  columns: number;
  orientation: "portrait" | "landscape";
}

// const props = defineProps<Props>()
const rows = ref(3);
const columns = ref(3);
const selectedIndex = ref([0, 0]);

const selectRowItem = (row: number, column: number, evt: MouseEvent) => {
  selectedIndex.value = [row, column];
};

const isSelectedRowItem = (row: number, column: number) => {
  return selectedIndex.value[0] === row && selectedIndex.value[1] === column ? "__selected" : "";
};
</script>

<style lang="scss" scoped>
</style>
