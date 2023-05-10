<template>
  <div class="cell-container">
    <div class="cell-column" v-for="c in columns">
      <div
        :class="`cell-row-item cell-row-item` + isSelectedRowItem(r, c)"
        @click="(evt) => selectRowItem(r, c, evt)"
        v-for="r in rows"
      >
        [{{ r }}, {{ c }}]
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";

// props
interface Props {
  rows: number;
  columns: number;
}

const props = defineProps<Props>()

const selectedIndex = ref([0, 0]);
const selectRowItem = (row: number, column: number, evt: MouseEvent) => {
  selectedIndex.value = [row, column];
};

const isSelectedRowItem = (row: number, column: number) => {
  return selectedIndex.value[0] === row && selectedIndex.value[1] === column ? "__selected" : "";
};
</script>

<style lang="scss" scoped>
.cell-container {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.cell-row-item {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  width: 4rem;
  height: 4rem;
  background-color: #27a465;

  &:hover {
    background-color: #25744c;
    cursor: pointer;
  }

  &__selected {
    background-color: #1f4b36;
  }
}

.cell-column {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>