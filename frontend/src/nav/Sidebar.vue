<template>
    <div style="height: 90%;">
        <va-sidebar :minimized="minimized" minimized-width="64px">
            <template v-for="item in items" :key="item.title">
                <va-sidebar-item :href="item.url" :active="true">
                    <va-sidebar-item-content>
                        <va-icon :name="item.icon" />
                        <va-sidebar-item-title>
                            {{ item.title }}
                        </va-sidebar-item-title>
                    </va-sidebar-item-content>
                </va-sidebar-item>
            </template>
        </va-sidebar>
    </div>

    <va-button v-model="minimized" :onclick="toggleMinimized" class="mt-2" label="Minimized" />
</template>

<script setup lang="ts">
import { ref } from "vue";
const emit = defineEmits(["update:minimized"]);

const items = [
    { title: "Home", icon: "home", url: "/"},
    { title: "Plants", icon: "grass", url: "/plants" },
    { title: "Boxes", icon: "inventory", url: "/boxes" },
    { title: "Gardens", icon: "yard", url: "/gardens" },
    { title: "About", icon: "info", url: "/about" },
    { title: "Settings", icon: "settings", url: "/settings" },
];

const minimized = ref(false);

const toggleMinimized = () => {
    minimized.value = !minimized.value;
    emit("update:minimized", minimized.value);
};

</script>