<template>
    <div class="card-container">
        <div class="card" v-for="plant in plants" style="width: 18rem;">
            <div class="card-body">
                <h5 class="card-title">{{ plant.name }}</h5>
                <h6 class="card-subtitle mb-2 text-muted">Card subtitle</h6>
                <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                <a href="#" class="card-link">Card link</a>
                <a href="#" class="card-link">Another link</a>
            </div>
        </div>
    </div>

</template>

<script setup lang="ts">
    import { Ref, ref } from 'vue'
    interface Plant {
        id: number;
        name: string;
        description: string;
        location: string;
    }
    const plants: Ref<Plant[]>= ref([
        {
            id: 1,
            name: 'Plant 1',
            description: 'Plant 1 description',
            location: 'Plant 1 location',
        },
        {
            id: 2,
            name: 'Plant 2',
            description: 'Plant 2 description',
            location: 'Plant 2 location',
        },
        {
            id: 3,
            name: 'Plant 3',
            description: 'Plant 3 description',
            location: 'Plant 3 location',
        },
        {
            id: 4,
            name: 'Plant 4',
            description: 'Plant 4 description',
            location: 'Plant 4 location',
        },
        {
            id: 5,
            name: 'Plant 5',
            description: 'Plant 5 description',
            location: 'Plant 5 location',
        }
    ])
</script>

<style lang="scss" scoped>
.card-container {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    align-content: center;
    .card {
        margin: 0.65rem;
    }
}

.plants {
    &__card {
        &__item {
            &__name {
                font-size: 1.5rem;
            }
            &__description {
                font-size: 1rem;
            }
            &__location {
                font-size: 1rem;
            }
        }
    }
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