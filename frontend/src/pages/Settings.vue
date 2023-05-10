<template>
    <div class="settings">
        <h1>Settings</h1>
        <div class="settings__list">
        <div class="settings__list__item" v-for="setting in settings" :key="setting.id">
            <router-link to="">
            <div class="settings__list__item__name">{{ setting.name }}</div>
            <div class="settings__list__item__description">{{ setting.description }}</div>
            </router-link>
        </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const settings = ref([
    {
        id: 1,
        name: 'Setting 1',
        description: 'Setting 1 description'
    },
    {
        id: 2,
        name: 'Setting 2',
        description: 'Setting 2 description'
    },
    {
        id: 3,
        name: 'Setting 3',
        description: 'Setting 3 description'
    }
])
</script>

<style scoped lang="scss">
.settings {
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
