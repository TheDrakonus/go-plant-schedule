<template>
    <div class="gardens">
        <h1>Gardens</h1>
        <div class="gardens__list">
        <div class="gardens__list__item" v-for="garden in gardens" :key="garden.id">
            <router-link to="">
            <div class="gardens__list__item__name">{{ garden.name }}</div>
            <div class="gardens__list__item__description">{{ garden.description }}</div>
            </router-link>
        </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const gardens = ref([
    {
        id: 1,
        name: 'Garden 1',
        description: 'Garden 1 description'
    },
    {
        id: 2,
        name: 'Garden 2',
        description: 'Garden 2 description'
    },
    {
        id: 3,
        name: 'Garden 3',
        description: 'Garden 3 description'
    }
])
</script>

<style scoped lang="scss">
.gardens {
    &__list {
        &__item {
            &__name {
                font-size: 1.5rem;
            }
            &__description {
                font-size: 1rem;
            }
        }
    }
}
</style>
